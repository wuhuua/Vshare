package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/service"
	"github.com/cloudwego/hertz/pkg/app"
)

type UserListResponse struct {
	model.Response
	UserList []model.User `json:"user_list"`
}

type FriendListResponse struct {
	model.Response
	FriendList []model.Friend `json:"user_list"`
}

func RelationAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	action, _ := strconv.Atoi(c.Query("action_type"))
	followId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	userId, _ := service.GetTokenId(token)
	if userId == 0 {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "no permission to follow"})
	} else {
		if action == 1 {
			service.FlolowById(userId, followId)
			c.JSON(http.StatusOK, model.Response{StatusCode: 0})
		} else {
			service.UnFlolowById(userId, followId)
			c.JSON(http.StatusOK, model.Response{StatusCode: 0})
		}
	}
}

func FollowList(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	validId, _ := service.GetTokenId(token)
	if validId != userId {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "No access permission"})
	} else {
		followList, _ := service.GetFollowListById(userId)
		c.JSON(http.StatusOK, UserListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			UserList: followList,
		})
	}
}

func FollowerList(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	validId, _ := service.GetTokenId(token)
	if validId != userId {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "No access permission"})
	} else {
		followerList, _ := service.GetFollowerListById(userId)
		c.JSON(http.StatusOK, UserListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			UserList: followerList,
		})
	}
}

func FriendList(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	validId, _ := service.GetTokenId(token)
	if validId != userId {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "No access permission"})
	} else {
		friends, _ := service.GetFriendsById(userId)
		c.JSON(http.StatusOK, FriendListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			FriendList: friends,
		})
	}
}
