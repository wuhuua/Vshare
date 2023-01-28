package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/util"
)

func Register(name string, password string) (int64, string, error) {
	userDao := repository.NewUserDaoInstance()
	if user, _ := userDao.GetUserByName(name); user.Id != 0 {
		fmt.Printf("error is %+v", *user)
		util.Logger.Error("name exists")
		return 0, "", errors.New("name exists")
	}
	id, _ := userDao.InitUserByName(name, password)
	token := Tokenize(name + password + util.GetDate())
	err := SendToken(token, id)
	if err != nil {
		util.Logger.Error("add token error")
		return 0, "", errors.New("add token error")
	}
	return id, token, nil
}

func Login(name string, password string) (int64, string, error) {
	userDao := repository.NewUserDaoInstance()
	user, _ := userDao.GetUserByName(name)
	if user == nil {
		util.Logger.Error("no such name")
		return 0, "", errors.New("no such name")
	}
	if user.Password != password {
		util.Logger.Error("password error")
		return 0, "", errors.New("password error")
	}
	token := Tokenize(name + password + util.GetDate())
	id := user.Id
	err := SendToken(token, id)
	if err != nil {
		util.Logger.Error("add token error")
		return 0, "", errors.New("add token error")
	}
	return id, token, nil
}

func GetUserInfo(id int64, token string) (*model.User, error) {
	userDao := repository.NewUserDaoInstance()
	relationDao := repository.NewRelationDaoInstance()
	user, _ := userDao.GetUserById(id)
	if user == nil {
		util.Logger.Error("no userinfo")
		return nil, errors.New("no userinfo")
	}
	validId, err := GetTokenId(token)
	if err != nil {
		util.Logger.Error("get token error")
		return nil, errors.New("get token error")
	}
	if validId == 0 {
		util.Logger.Error("no permission")
		return nil, errors.New("no permission")
	}
	flag, _ := relationDao.IsFollowed(strconv.FormatInt(validId, 10), strconv.FormatInt(id, 10))
	user.IsFollow = flag
	return user, nil
}
