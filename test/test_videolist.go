package test

import (
	"fmt"

	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/service"
)

func Test_videolist(ip string, password string) {
	repository.InitMySQL(ip, password)
	streams, _ := service.GetStreamList(20230001)
	fmt.Printf("%+v\n", streams)
}
