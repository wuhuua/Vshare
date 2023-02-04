package test

import (
	"fmt"

	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/service"
)

func Test_message(mysqlIp string, mysqlPassword string, redisIp string, redisPassword string, redisBase []int) {
	for _, num := range redisBase {
		err := repository.InitRedis(redisIp, redisPassword, num)
		if err != nil {
			fmt.Printf("Redisbase%d initialize error", num)
			return
		}
	}
	repository.InitMySQL(mysqlIp, mysqlPassword)
	messages, _ := service.GetMessageList(20230001, 20230002)
	fmt.Printf("%+v", messages)
	service.SendMessageById(20230001, 20230002, "test")
	messages, _ = service.GetMessageList(20230001, 20230002)
	fmt.Printf("%+v", messages)
}
