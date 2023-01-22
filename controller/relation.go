package controller

import (
	"context"
	"net/http"

	"github.com/Iscolito/Vshare/repository"
	"github.com/cloudwego/hertz/pkg/app"
)

type UserListResponse struct {
	repository.Response
	UserList []repository.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, repository.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: repository.Response{
			StatusCode: 0,
		},
		UserList: []repository.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: repository.Response{
			StatusCode: 0,
		},
		UserList: []repository.User{DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: repository.Response{
			StatusCode: 0,
		},
		UserList: []repository.User{DemoUser},
	})
}
