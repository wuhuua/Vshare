package repository

import (
	"strconv"
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

func (*ChatDao) ChatCheckPoint(userId string, aimId string, gap int, checkType int) int64 {
	chatKey := userId + aimId
	exist, _ := rdb[6].Exists(chatKey).Result()
	var res string
	var msgNum int64
	/* 常规情况下处于轮询状态,时间戳始终更新,exist!=0
	* 出现exist==0时说明当前用户刚好打开聊天框,此时返回系统设计的最多返回条数
	 */
	if exist == 0 {
		if checkType == 0 {
			rdb[6].Set(chatKey, checkType, time.Duration(gap)*time.Second)
		}
		return -1
	} else {
		// 如果是轮询状态,查询对方是否更新消息,返回当前待消费消息数
		if checkType == 0 {
			res, _ = rdb[6].Get(chatKey).Result()
			msgNum, _ = strconv.ParseInt(res, 10, 64)
			rdb[6].Set(chatKey, checkType, time.Duration(gap)*time.Second)
			return msgNum
			// 如果是发送消息状态，则向对方的消息redis库中推送一条数据
		} else {
			rdb[6].Incr(chatKey)
			res, _ = rdb[6].Get(chatKey).Result()
			msgNum, _ = strconv.ParseInt(res, 10, 64)
			return msgNum
		}
	}
}
