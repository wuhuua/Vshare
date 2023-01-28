package test

import (
	"fmt"

	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/service"
)

func Test_stream(ip string, password string) {
	repository.InitMySQL(ip, password)
	streams, _ := service.GetStreams()
	fmt.Printf("%+v\n", streams)
}
