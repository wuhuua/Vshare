package test

import (
	"fmt"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/repository"
)

func Test_gorm(ip string, password string) {
	repository.InitMySQL(ip, password)
	var user *model.User
	var id int64
	var name string
	id = 20230001
	name = "Iori"
	user, _ = repository.NewUserDaoInstance().GetUserById(id)
	fmt.Println("按照id查找到用户名:" + user.Name)
	user, _ = repository.NewUserDaoInstance().GetUserByName(name)
	fmt.Println("按照用户名查找到id")
	fmt.Println(user.Id)
	id, _ = repository.NewUserDaoInstance().InitUserByName("wuhuua", "12345678")
	fmt.Println("创建新用户,id为:")
	fmt.Println(id)
	user, _ = repository.NewUserDaoInstance().GetUserById(id)
	fmt.Println("新用户用户名为:" + user.Name)
}
