package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/service"
	"github.com/cloudwego/hertz/pkg/app"
)

type CommentListResponse struct {
	model.Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	model.Response
	Comment model.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	userId, _ := service.GetTokenId(token)
	if userId != 0 {
		if actionType == "1" {
			user, err := service.GetUserInfo(userId, token)
			if err != nil {
				c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
			}
			text := c.Query("comment_text")
			videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
			id, date, _ := service.PublishComment(text, userId, videoId)
			c.JSON(http.StatusOK, CommentActionResponse{Response: model.Response{StatusCode: 0},
				Comment: model.Comment{
					Id:         id,
					User:       *user,
					Content:    text,
					CreateDate: date,
				}})
			return
		}
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func CommentList(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	userId, _ := service.GetTokenId(token)
	if userId == 0 {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "No permission"})
		return
	}
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	comments, err := service.GetCommentList(videoId)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "Find comments error"})
		return
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    model.Response{StatusCode: 0},
		CommentList: comments,
	})
}
