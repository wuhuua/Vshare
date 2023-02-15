package main

import (
	"fmt"

	"github.com/Iscolito/Vshare/config"
	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/service"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
)

var HostIp = config.LoadConfig().HostIp
var mysqlIp = config.LoadConfig().MysqlIp
var mysqlPassword = config.LoadConfig().MysqlPassword
var redisIp = config.LoadConfig().RedisIp
var redisPassword = config.LoadConfig().RedisPassword

// Redis数据库编号,0号缓存token,1号缓存关注列表,2号缓存被关注列表,3号缓存视频点赞,4号缓存用户喜欢,5存储聊天记录,6存储聊天指针缓存
var redisbase = []int{0, 1, 2, 3, 4, 5, 6}

func RunServer() {
	var err error
	err = repository.InitMySQL(mysqlIp, mysqlPassword)
	if err != nil {
		fmt.Println("MySQL initialize error")
		return
	}
	for _, num := range redisbase {
		err = repository.InitRedis(redisIp, redisPassword, num)
		if err != nil {
			fmt.Printf("Redisbase%d initialize error", num)
			return
		}
	}

	go service.RunMessageServer()
	h := server.Default(
		server.WithHostPorts(HostIp),
		server.WithMaxRequestBodySize(20<<20),
		server.WithTransport(standard.NewTransporter),
	)

	initRouters(h)

	h.Spin()
}

func TestServer() {
	//test.Test_favorite(mysqlIp, mysqlPassword, redisIp, redisPassword)
	//test.Test_regist(mysqlIp, mysqlPassword, redisIp, redisPassword)
	//test.Test_friend(mysqlIp, mysqlPassword, redisIp, redisPassword, redisbase)
	//test.Test_message(mysqlIp, mysqlPassword, redisIp, redisPassword, redisbase)
}
