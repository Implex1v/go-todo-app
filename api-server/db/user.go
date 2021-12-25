package db

import (
	"api-server/types"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserDao interface {
	Get(id int64) (error, *types.User)
	GetAll() (error, *[]types.User)
}

type DefaultUserDao struct {
	db *gorm.DB
}

func (d DefaultUserDao) Get(id int64) (error, *types.User) {
	var user *types.User
	d.db.First(&user, id)
	if user == nil {
		return errors.New(fmt.Sprintf("UserDao: User with id '%d' not found", id)), nil
	} else {
		return nil, user
	}
}

func (d DefaultUserDao) GetAll() (error, *[]types.User) {
	var users *[]types.User
	d.db.Find(&users)
	if users == nil {
		return errors.New("UserDao: Error while fetching all users"), nil
	} else {
		return nil, users
	}
}

func NewUserDao(db *gorm.DB) UserDao {
	return &DefaultUserDao{db: db}
}
