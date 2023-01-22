package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/service"
	"github.com/cloudwego/hertz/pkg/app"
)

type FeedResponse struct {
	repository.Response
	VideoList []repository.Video `json:"video_list,omitempty"`
	NextTime  int64              `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(ctx context.Context, c *app.RequestContext) {
	videos, _ := service.GetStreams()
	c.JSON(http.StatusOK, FeedResponse{
		Response:  repository.Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
