package db

import (
	"api-server/types"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserDao interface {
	Create(user *types.User) (error, *types.User)
	Get(id int64) (error, *types.User)
	GetAll() (error, *[]types.User)
}

type DefaultUserDao struct {
	db *gorm.DB
}

func (d DefaultUserDao) Create(user *types.User) (error, *types.User) {
	e := d.db.Create(user).Error
	if e != nil {
		return errors.New(fmt.Sprintf("UserDao: Cloud not create User: '%s'", e.Error())), nil
	} else {
		return nil, user
	}
}

func (d DefaultUserDao) Get(id int64) (error, *types.User) {
	var user *types.User
	var r = d.db.First(&user, id)
	if r.Error != nil {
		return errors.New(fmt.Sprintf("UserDao: User with id '%d' not found: '%s'", id, r.Error.Error())), nil
	} else {
		return nil, user
	}
}

func (d DefaultUserDao) GetAll() (error, *[]types.User) {
	var users *[]types.User
	var r = d.db.Find(&users)
	if r.Error != nil {
		return errors.New(fmt.Sprintf("UserDao: Error while fetching all users: '%s'", r.Error.Error())), nil
	} else {
		return nil, users
	}
}

func NewUserDao(db *gorm.DB) UserDao {
	return &DefaultUserDao{db: db}
}
