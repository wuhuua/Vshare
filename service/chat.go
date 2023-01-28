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

func formatMessage(text string) string {
	return fmt.Sprintf("%s%s", util.GetTime(), text)
}

func SendMessageById(userId int64, aimId int64, text string) {
	chatDao := repository.NewChatDaoInstance()
	chatDao.InitMessage(getChatRoom(userId, aimId), formatMessage(text), 24)
}

func GetMessageList(userId int64, aimId int64) ([]model.Message, error) {
	chatDao := repository.NewChatDaoInstance()
	messages, _ := chatDao.GetMessage(getChatRoom(userId, aimId))
	lenHeader := len(util.GetTime())
	chatList := make([]model.Message, len(messages))
	for i, message := range messages {
		chatList[i].Id = int64(i)
		chatList[i].Content = message[lenHeader:]
		chatList[i].CreateTime = message[:lenHeader]
	}
	return chatList, nil
}
