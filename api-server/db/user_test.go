package db

import (
	"api-server/types"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"regexp"
	"testing"
)

func (s *Suite) Test_UserDao_Get_One() {
	expected := &types.User{
		Password: "fisch!",
		Username: "longcat",
		Email:    "meow@long.cat",
	}

	query := "SELECT * FROM \"users\" WHERE \"users\".\"id\" = $1 AND \"users\".\"deleted_at\" IS NULL ORDER BY \"users\".\"id\" LIMIT 1"
	args := func(e *sqlmock.ExpectedQuery) *sqlmock.ExpectedQuery {
		return e.WithArgs(1)
	}

	s.mockResult(query, []types.User{*expected}, args)

	err, actual := s.dao.Get(1)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), expected, actual)
}

func (s *Suite) Test_UserDao_Get_None() {
	query := "SELECT * FROM \"users\" WHERE \"users\".\"id\" = $1 AND \"users\".\"deleted_at\" IS NULL ORDER BY \"users\".\"id\" LIMIT 1"
	args := func(e *sqlmock.ExpectedQuery) *sqlmock.ExpectedQuery {
		return e.WithArgs(1)
	}
	s.mockResult(query, []types.User{}, args)

	err, actual := s.dao.Get(1)
	if err == nil {
		s.T().Errorf("Expected error but got nil")
	}
	if actual != nil {
		s.T().Errorf("Expected nil but got %+v\n", actual)
	}
}

func (s *Suite) Test_UserDao_GetAll_None() {
	query := ""
	s.mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(s.mock.NewRows([]string{
			"id",
			"created_at",
			"updated_at",
			"deleted_at",
			"username",
			"email",
			"password",
		}))

	err, actual := s.dao.GetAll()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 0, len(*actual))
}

func (s *Suite) Test_UserDao_GetAll_Multiple() {
	query := ""
	expected := []types.User{
		*&types.User{
			Password: "nuss!",
			Username: "nuss",
			Email:    "nuss@liebe.com",
		}, *&types.User{
			Password: "fisch!",
			Username: "longcat",
			Email:    "meow@long.cat",
		},
	}

	args := func(e *sqlmock.ExpectedQuery) *sqlmock.ExpectedQuery {
		return e
	}

	s.mockResult(query, expected, args)

	err, actual := s.dao.GetAll()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(*actual))
	assert.Equal(s.T(), expected, *actual)
}

func (s *Suite) mockResult(
	query string,
	users []types.User,
	args func(e *sqlmock.ExpectedQuery) *sqlmock.ExpectedQuery,
) {
	if len(users) == 0 {
		s.mock.
			ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(1).
			WillReturnError(errors.New("Empty"))

		return
	}

	rows := s.mock.NewRows([]string{
		"id",
		"created_at",
		"updated_at",
		"deleted_at",
		"username",
		"email",
		"password",
	})

	for _, u := range users {
		rows.AddRow(
			"0",
			u.CreatedAt,
			u.UpdatedAt,
			u.DeletedAt,
			u.Username,
			u.Email,
			u.Password,
		)
	}

	args(s.mock.ExpectQuery(regexp.QuoteMeta(query))).
		WillReturnRows(rows)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	dao  UserDao
	user *types.User
}

func (s *Suite) SetupSuite() {
	var db *sql.DB
	var err error
	nopLogger := logger.Default.LogMode(logger.Silent)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	d := postgres.Dialector{
		Config: &postgres.Config{
			Conn: db,
		},
	}

	s.DB, err = gorm.Open(d, &gorm.Config{
		Logger: nopLogger,
	})
	require.NoError(s.T(), err)

	s.dao = NewUserDao(s.DB, zap.NewNop())
}
