package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/service"
	"github.com/cloudwego/hertz/pkg/app"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]model.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	model.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	model.Response
	User model.User `json:"user"`
}

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	id, token, _ := service.Register(username, password)
	if id == 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 0},
			UserId:   id,
			Token:    token,
		})
	}
}

func Login(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	id, token, _ := service.Login(username, password)

	if id != 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 0},
			UserId:   id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist or password incorrect"},
		})
	}
}

func UserInfo(ctx context.Context, c *app.RequestContext) {
	id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	user, err := service.GetUserInfo(id, token)
	if err == nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 0},
			User:     *user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist or no permission"},
		})
	}
}
