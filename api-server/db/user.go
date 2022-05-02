package db

import (
	"api-server/types"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserDao interface {
	Create(user *types.User) (error, *types.User)
	Get(id int64) (error, *types.User)
	GetAll() (error, *[]types.User)
	Update(user *types.User) (error, *types.User)
}

type DefaultUserDao struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (d DefaultUserDao) Update(user *types.User) (error, *types.User) {
	d.db.Model(&user).Updates(user)
	e := d.db.Error
	if e != nil {
		d.logger.Warn(fmt.Sprintf("UserDao: cloud not update user: '%s'", e.Error()))
		return errors.New("could not update user"), nil
	} else {
		return nil, user
	}
}

func (d DefaultUserDao) Create(user *types.User) (error, *types.User) {
	e := d.db.Create(user).Error
	if e != nil {
		d.logger.Warn(fmt.Sprintf("UserDao: cloud not create user: '%s'", e.Error()))
		return errors.New("could not Create user"), nil
	} else {
		return nil, user
	}
}

func (d DefaultUserDao) Get(id int64) (error, *types.User) {
	var user *types.User
	var r = d.db.First(&user, id)
	if r.Error != nil {
		d.logger.Warn(fmt.Sprintf("UserDao: User with id '%d' not found: '%s'", id, r.Error.Error()))
		return errors.New("user not found"), nil
	} else {
		return nil, user
	}
}

func (d DefaultUserDao) GetAll() (error, *[]types.User) {
	var users *[]types.User
	var r = d.db.Find(&users)
	if r.Error != nil {
		d.logger.Warn(fmt.Sprintf("UserDao: Error while fetching all users: '%s'", r.Error.Error()))
		return errors.New("failed to fetch all users"), nil
	} else {
		return nil, users
	}
}

func NewUserDao(db *gorm.DB, l *zap.Logger) UserDao {
	return &DefaultUserDao{db: db, logger: l}
}
