package controller

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/service"
	"github.com/cloudwego/hertz/pkg/app"
)

type VideoListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(ctx context.Context, c *app.RequestContext) {
	token := c.PostForm("token")
	userId, _ := service.GetTokenId(token)
	if userId == 0 {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user, _ := service.GetUserInfo(userId, token)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/static/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	service.PublishVideo(userId, filename, finalName)
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

func PublishList(ctx context.Context, c *app.RequestContext) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	validId, _ := service.GetTokenId(token)
	if validId == 0 {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "No access permission"})
	} else {
		videos, _ := service.GetStreamList(userId)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			VideoList: videos,
		})
	}
}
