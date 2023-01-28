package test

import (
	"fmt"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/service"
)

func Test_comment(ip string, password string) {
	repository.InitMySQL(ip, password)
	var comments []model.Comment
	comments, _ = service.GetCommentList(20230001)
	fmt.Printf("%+v", comments)
	id, date, _ := service.PublishComment("test comments", 20230002, 20230001)
	fmt.Printf("新发布评论id为:%v,发布时间为:%v", id, date)
	comments, _ = service.GetCommentList(20230001)
	fmt.Printf("%+v", comments)
}
