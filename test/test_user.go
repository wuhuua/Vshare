package test

import (
	"fmt"

	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/service"
)

func Test_user(ip string, password string, redisip string, redispassword string) {
	repository.InitMySQL(ip, password)
	repository.InitRedis(redisip, redispassword, 0)
	repository.InitRedis(redisip, redispassword, 1)
	user, err := service.GetUserInfo(20230002, "40475d1a0ade9f7c0ed800b623deed87")
	if err != nil {
		fmt.Println("Certain error")
	}
	fmt.Printf("%+v", user)
}
