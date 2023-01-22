package controller

import (
	"context"
	"net/http"

	"github.com/Iscolito/Vshare/repository"
	"github.com/cloudwego/hertz/pkg/app"
)

type CommentListResponse struct {
	repository.Response
	CommentList []repository.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	repository.Response
	Comment repository.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			c.JSON(http.StatusOK, CommentActionResponse{Response: repository.Response{StatusCode: 0},
				Comment: repository.Comment{
					Id:         1,
					User:       user,
					Content:    text,
					CreateDate: "05-01",
				}})
			return
		}
		c.JSON(http.StatusOK, repository.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    repository.Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
