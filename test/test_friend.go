package test

import (
	"fmt"

	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/service"
)

func Test_friend(mysqlIp string, mysqlPassword string, redisIp string, redisPassword string, redisBase []int) {
	for _, num := range redisBase {
		err := repository.InitRedis(redisIp, redisPassword, num)
		if err != nil {
			fmt.Printf("Redisbase%d initialize error", num)
			return
		}
	}
	repository.InitMySQL(mysqlIp, mysqlPassword)
	friends, _ := service.GetFriendsById(20230002)
	fmt.Printf("%+v", friends)
}
