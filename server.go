package main

import (
	"fmt"

	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/service"
	"github.com/Iscolito/Vshare/test"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
)

const HostIp = "192.168.3.9:8080"       //服务端ip
const mysqlIp = "127.0.0.1:3306"        //MySQL数据库ip
const mysqlPassword = "12345678"        //MySQL密码
const redisIp = "127.0.0.1:6379"        //Redis数据库ip
const redisPassword = ""                //Redis密码
var redisbase = []int{0, 1, 2, 3, 4, 5} //Redis数据库编号,0号缓存token,1号缓存关注列表,2号缓存被关注列表,3号缓存视频点赞,4号缓存用户喜欢,5存储聊天记录

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
	test.Test_regist(mysqlIp, mysqlPassword, redisIp, redisPassword)
}
