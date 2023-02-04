package service

import (
	"fmt"
	"strconv"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/util"
)

func getChatRoom(userId int64, aimId int64) string {
	if userId < aimId {
		return fmt.Sprintf("%s_%s", strconv.FormatInt(userId, 10), strconv.FormatInt(aimId, 10))
	} else {
		return fmt.Sprintf("%s_%s", strconv.FormatInt(aimId, 10), strconv.FormatInt(userId, 10))
	}
}

func formatMessage(text string, userId int64, aimId int64) string {
	strUserId := strconv.FormatInt(userId, 10)
	strAimId := strconv.FormatInt(aimId, 10)
	return fmt.Sprintf("%s%s%s%s", util.GetTime(), strUserId, strAimId, text)
}

func SendMessageById(userId int64, aimId int64, text string) {
	chatDao := repository.NewChatDaoInstance()
	chatDao.InitMessage(getChatRoom(userId, aimId), formatMessage(text, userId, aimId))
}

func GetMessageList(userId int64, aimId int64) ([]model.Message, error) {
	chatDao := repository.NewChatDaoInstance()
	messages, _ := chatDao.GetMessage(getChatRoom(userId, aimId), 30)
	lenHeader := len(util.GetTime())
	lenUserId := len(strconv.FormatInt(userId, 10))
	lenAimId := len(strconv.FormatInt(aimId, 10))
	lenId := lenUserId + lenAimId
	chatList := make([]model.Message, len(messages))
	for i, message := range messages {
		chatList[i].Id = int64(20230001 + i)
		chatList[i].Content = message[lenHeader+lenId:]
		chatList[i].CreateTime = message[:lenHeader]
		chatList[i].FromUserId, _ = strconv.ParseInt(message[lenHeader:lenHeader+lenUserId], 10, 64)
		chatList[i].ToUserId, _ = strconv.ParseInt(message[lenHeader+lenUserId:lenHeader+lenId], 10, 64)
	}
	return chatList, nil
}
