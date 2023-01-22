package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/service"
	"github.com/cloudwego/hertz/pkg/app"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]repository.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	repository.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	repository.Response
	User repository.User `json:"user"`
}

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	id, token, _ := service.Register(username, password)
	if id == 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: repository.Response{StatusCode: 0},
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
			Response: repository.Response{StatusCode: 0},
			UserId:   id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(ctx context.Context, c *app.RequestContext) {
	id, _ := strconv.Atoi(c.Query("id"))
	token := c.Query("token")
	user, err := service.GetUserInfo(int64(id), token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: repository.Response{StatusCode: 0},
			User:     *user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist or no permission"},
		})
	}
}
