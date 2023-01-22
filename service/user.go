package service

import (
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/util"
)

func tokenize(rawStr string) string {
	data := []byte(rawStr)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func Register(name string, password string) (int64, string, error) {
	userDao := repository.NewUserDaoInstance()
	if user, _ := userDao.GetUserByName(name); user == nil {
		id, _ := userDao.InitUserByName(name, password)
		return id, tokenize(name + password), nil
	} else {
		util.Logger.Error("name exists")
		return 0, "", errors.New("name exists")
	}
}

func Login(name string, password string) (int64, string, error) {
	userDao := repository.NewUserDaoInstance()
	user, _ := userDao.GetUserByName(name)
	if user == nil {
		util.Logger.Error("no such name")
		return 0, "", errors.New("no such name")
	} else {
		return user.Id, tokenize(name + password), nil
	}
}

func GetUserInfo(id int64, token string) (*repository.User, error) {
	userDao := repository.NewUserDaoInstance()
	user, _ := userDao.GetUserById(id)
	if user == nil {
		util.Logger.Error("no userinfo")
		return nil, errors.New("no userinfo")
	} else if token != tokenize(user.Name+user.Password) {
		util.Logger.Error("no permission")
		return nil, errors.New("no permission")
	} else {
		return user, nil
	}
}
