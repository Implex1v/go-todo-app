package db

import (
	"api-server/types"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func (s *Suite) Test_UserDao_GetOne() {
	expected := &types.User{
		Password: "fisch!",
		Username: "longcat",
		Email:    "meow@long.cat",
	}

	query := "SELECT * FROM \"users\" WHERE \"users\".\"id\" = $1 AND \"users\".\"deleted_at\" IS NULL ORDER BY \"users\".\"id\" LIMIT 1"
	s.mockResult(query, []types.User{*expected})

	err, actual := s.dao.Get(1)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), expected, actual)
}

func (s *Suite) Test_UserDao_GetNone() {
	query := "SELECT * FROM \"users\" WHERE \"users\".\"id\" = $1 AND \"users\".\"deleted_at\" IS NULL ORDER BY \"users\".\"id\" LIMIT 1"
	s.mockResult(query, []types.User{})

	err, actual := s.dao.Get(1)
	if err == nil {
		s.T().Errorf("Expected error but got nil")
	}
	if actual != nil {
		s.T().Errorf("Expected nil but got %+v\n", actual)
	}
}

func (s *Suite) mockResult(query string, users []types.User) {
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

	s.mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(1).
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

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	d := postgres.Dialector{
		Config: &postgres.Config{
			Conn: db,
		},
	}

	s.DB, err = gorm.Open(d, &gorm.Config{})
	require.NoError(s.T(), err)

	s.dao = NewUserDao(s.DB)
}
