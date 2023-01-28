package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/service"
	"github.com/cloudwego/hertz/pkg/app"
)

func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	strVideoId := c.Query("video_id")
	videoId, _ := strconv.ParseInt(strVideoId, 10, 64)
	validId, _ := service.GetTokenId(token)
	if validId == 0 {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	} else {
		service.LikeById(validId, videoId)
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	}
}

func FavoriteList(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	strId := c.Query("user_id")
	userId, _ := strconv.ParseInt(strId, 10, 64)
	validId, _ := service.GetTokenId(token)
	if validId == 0 {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "No permission to look up"})
	} else {
		videos, _ := service.GetLikeListById(userId)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			VideoList: videos,
		})
	}

}
