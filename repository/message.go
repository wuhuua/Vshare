package repository

import (
	"sync"
	"time"
)

type ChatDao struct {
}

var chatDao *ChatDao
var chatOnce sync.Once

func NewChatDaoInstance() *ChatDao {
	chatOnce.Do(
		func() {
			chatDao = &ChatDao{}
		})
	return chatDao
}

func (*ChatDao) InitMessage(chatRoom string, text string, hour int) (int64, error) {
	return rdb[5].LPush(chatRoom, text, time.Duration(hour)*time.Hour).Result()
}

func (*ChatDao) GetMessage(chatRoom string) ([]string, error) {
	return rdb[5].LRange(chatRoom, 0, 30).Result()
}
