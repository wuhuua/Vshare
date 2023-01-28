package test

import (
	"fmt"

	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/service"
	"github.com/Iscolito/Vshare/util"
)

func Test_redis(ip string, password string, baseNum int) {
	repository.InitRedis(ip, password, baseNum)
	token := service.Tokenize("Iori" + "12345678" + util.GetDate())
	service.SendToken(token, 20230001)
	id, _ := service.GetTokenId(token)
	fmt.Println(id)
	//DelToken("abcd")
}
