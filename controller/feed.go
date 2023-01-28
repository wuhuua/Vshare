package controller

import (
	"context"
	"net/http"
	"path"
	"time"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/service"
	"github.com/cloudwego/hertz/pkg/app"
)

type FeedResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

func Feed(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	validId, _ := service.GetTokenId(token)
	var videos []model.Video
	if token == "" || validId == 0 {
		videos, _ = service.GetStreams()
	} else {
		videos, _ = service.GetStreamsById(validId)
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}

func VideoFile(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")
	filename := path.Join("./public/static/", name)
	c.File(filename)
}
