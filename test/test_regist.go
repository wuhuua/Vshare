package test

import (
	"fmt"

	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/service"
)

func Test_regist(ip string, password string, redisip string, redispassword string) {
	repository.InitMySQL(ip, password)
	repository.InitRedis(redisip, redispassword, 0)
	id, token, _ := service.Register("wuhuua", "12345678")
	fmt.Printf("%+v_%+v", id, token)
}
