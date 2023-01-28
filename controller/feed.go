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

// Feed same demo video list for every request
func Feed(ctx context.Context, c *app.RequestContext) {
	videos, _ := service.GetStreams()
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
