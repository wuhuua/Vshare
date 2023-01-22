package main

import (
	"github.com/Iscolito/Vshare/service"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
)

func main() {
	go service.RunMessageServer()

	h := server.Default(
		server.WithHostPorts("127.0.0.1:8080"),
		server.WithMaxRequestBodySize(20<<20),
		server.WithTransport(standard.NewTransporter),
	)

	initRouters(h)

	h.Spin()
}
