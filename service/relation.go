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
	strFollowId := strconv.FormatInt(followId, 10)
	strUserId := strconv.FormatInt(userId, 10)
	isFriends, _ := relationDao.IsFollowed(strFollowId, strUserId)
	if isFriends {
		chatDao := repository.NewChatDaoInstance()
		chatRoom := getChatRoom(userId, followId)
		err := chatDao.InitChatRoom(chatRoom, formatMessage("我们现在是朋友啦", userId, followId))
		if err != nil {
			return err
		}
	}
	return relationDao.FollowAction(strUserId, strFollowId)
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

func GetFriendsById(userId int64) ([]model.Friend, error) {
	relationDao := repository.NewRelationDaoInstance()
	idList, _ := relationDao.GetFollowerList(strconv.FormatInt(userId, 10))
	if idList == nil {
		util.Logger.Error("get followerlist error")
		return nil, errors.New("get followerlist error")
	}
	friendList := make([]model.Friend, 0)
	userDao := repository.NewUserDaoInstance()
	messageDao := repository.NewChatDaoInstance()
	var user *model.User
	var isFollow bool
	var messages []string
	var message string
	var msgType int64
	lenHeader := len(util.GetTime())
	lenUserId := len(strconv.FormatInt(userId, 10))
	var lenAimId int
	var lenId int
	for _, id := range idList {
		user, _ = userDao.GetUserById(id)
		isFollow, _ = relationDao.IsFollowed(strconv.FormatInt(userId, 10), strconv.FormatInt(id, 10))
		if isFollow {
			user.IsFollow = true
			messages, _ = messageDao.GetMessage(getChatRoom(userId, id), 1)
			if strconv.FormatInt(userId, 10) == messages[0][lenHeader:lenHeader+lenUserId] {
				msgType = 1
			} else {
				msgType = 0
			}
			lenAimId = len(strconv.FormatInt(id, 10))
			lenId = lenUserId + lenAimId
			message = messages[0][lenHeader+lenId:]
			friendList = append(friendList, model.Friend{User: *user, Message: message, MsgType: msgType})
		}
	}
	return friendList, nil
}
