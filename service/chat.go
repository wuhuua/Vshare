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
	chatDao.ChatCheckPoint(strconv.FormatInt(aimId, 10), strconv.FormatInt(userId, 10), 2, 1)
}

func GetMessageList(userId int64, aimId int64) ([]model.Message, error) {
	chatDao := repository.NewChatDaoInstance()
	chatRoom := getChatRoom(userId, aimId)
	strUserId := strconv.FormatInt(userId, 10)
	strAimId := strconv.FormatInt(aimId, 10)
	msgNum := chatDao.ChatCheckPoint(strUserId, strAimId, 2, 0)
	// 此处作为返回消息数量的最大值,留出修改的接口
	if msgNum == -1 {
		msgNum = 30
	} else if msgNum == 0 {
		var emptyChatList []model.Message
		return emptyChatList, nil
	}
	messages, _ := chatDao.GetMessage(chatRoom, msgNum)
	lenHeader := len(util.GetTime())
	lenUserId := len(strUserId)
	lenAimId := len(strAimId)
	lenId := lenUserId + lenAimId
	chatList := make([]model.Message, len(messages))
	var strTime string
	for i, message := range messages {
		num := len(messages) - 1 - i
		chatList[num].Id = int64(20230001 + num)
		chatList[num].Content = message[lenHeader+lenId:]
		/* 前端demo暂时要求返回int64类型的数据
		chatList[i].CreateTime = message[:lenHeader]
		*/
		strTime = message[:lenHeader]
		chatList[num].CreateTime, _ = strconv.ParseInt(strTime[11:13]+strTime[14:16], 10, 64)
		chatList[num].FromUserId, _ = strconv.ParseInt(message[lenHeader:lenHeader+lenUserId], 10, 64)
		chatList[num].ToUserId, _ = strconv.ParseInt(message[lenHeader+lenUserId:lenHeader+lenId], 10, 64)
	}
	return chatList, nil
}
