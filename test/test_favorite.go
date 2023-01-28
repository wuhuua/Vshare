package test

import (
	"fmt"

	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/service"
)

func Test_favorite(ip string, password string, redisip string, redispassword string) {
	repository.InitMySQL(ip, password)
	repository.InitRedis(redisip, redispassword, 3)
	videos, _ := service.GetLikeListById(20230002)
	fmt.Printf("%+v", videos)
}
