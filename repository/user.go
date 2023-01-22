package repository

import (
	"sync"

	"github.com/Iscolito/Vshare/util"
	"gorm.io/gorm"
)

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) GetUserById(id int64) (*User, error) {
	users := User{}
	err := db.Where("id = ?", id).Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return &users, nil
}

func (*UserDao) GetUserByName(name string) (*User, error) {
	users := User{}
	err := db.Where("name = ?", name).Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find user by name err:" + err.Error())
		return nil, err
	}
	return &users, nil
}

func (*UserDao) InitUserByName(name string, password string) (int64, error) {
	user := &User{Name: name, Password: password}
	db.Create(user)
	return user.Id, nil
}
