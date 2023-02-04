package repository

import (
	"sync"
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

func (*ChatDao) InitMessage(chatRoom string, text string) (int64, error) {
	return rdb[5].LPush(chatRoom, text).Result()
}

func (*ChatDao) GetMessage(chatRoom string, num int64) ([]string, error) {
	return rdb[5].LRange(chatRoom, 0, num).Result()
}

func (*ChatDao) InitChatRoom(chatRoom string, initMessage string) error {
	exist, err := rdb[5].Exists(chatRoom).Result()
	if err != nil {
		return err
	}
	if exist == 1 {
		return nil
	}
	rdb[5].LPush(chatRoom, initMessage).Result()
	return nil
}
