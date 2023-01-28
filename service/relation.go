package service

import (
	"errors"
	"strconv"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/util"
)

func FlolowById(userId int64, followId int64) error {
	relationDao := repository.NewRelationDaoInstance()
	return relationDao.FollowAction(strconv.FormatInt(userId, 10), strconv.FormatInt(followId, 10))
}

func UnFlolowById(userId int64, followId int64) error {
	relationDao := repository.NewRelationDaoInstance()
	return relationDao.UnFollowAction(strconv.FormatInt(userId, 10), strconv.FormatInt(followId, 10))
}

func GetFollowListById(userId int64) ([]model.User, error) {
	relationDao := repository.NewRelationDaoInstance()
	idList, _ := relationDao.GetFollowList(strconv.FormatInt(userId, 10))
	if idList == nil {
		util.Logger.Error("get followlist error")
		return nil, errors.New("get followlist error")
	}
	userList := make([]model.User, len(idList))
	userDao := repository.NewUserDaoInstance()
	var user *model.User
	for i, id := range idList {
		user, _ = userDao.GetUserById(id)
		userList[i] = *user
		userList[i].IsFollow = true
	}
	return userList, nil
}

func GetFollowerListById(userId int64) ([]model.User, error) {
	relationDao := repository.NewRelationDaoInstance()
	idList, _ := relationDao.GetFollowerList(strconv.FormatInt(userId, 10))
	if idList == nil {
		util.Logger.Error("get followerlist error")
		return nil, errors.New("get followerlist error")
	}
	userList := make([]model.User, len(idList))
	userDao := repository.NewUserDaoInstance()
	var user *model.User
	for i, id := range idList {
		user, _ = userDao.GetUserById(id)
		userList[i] = *user
		userList[i].IsFollow, _ = relationDao.IsFollowed(strconv.FormatInt(userId, 10), strconv.FormatInt(id, 10))
	}
	return userList, nil
}

func GetFriendsById(userId int64) ([]model.User, error) {
	relationDao := repository.NewRelationDaoInstance()
	idList, _ := relationDao.GetFollowerList(strconv.FormatInt(userId, 10))
	if idList == nil {
		util.Logger.Error("get followerlist error")
		return nil, errors.New("get followerlist error")
	}
	userList := make([]model.User, 0)
	userDao := repository.NewUserDaoInstance()
	var user *model.User
	var isFollow bool
	for _, id := range idList {
		user, _ = userDao.GetUserById(id)
		isFollow, _ = relationDao.IsFollowed(strconv.FormatInt(userId, 10), strconv.FormatInt(id, 10))
		if isFollow {
			user.IsFollow = true
			userList = append(userList, *user)
		}
	}
	return userList, nil
}
