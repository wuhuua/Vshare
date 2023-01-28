package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/service"
	"github.com/cloudwego/hertz/pkg/app"
)

var tempChat = map[string][]model.Message{}

type ChatResponse struct {
	model.Response
	MessageList []model.Message `json:"message_list"`
}

func MessageAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	strAimId := c.Query("to_user_id")
	aimId, _ := strconv.ParseInt(strAimId, 10, 64)
	content := c.Query("content")
	validId, _ := service.GetTokenId(token)
	if validId == 0 {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	} else {
		service.SendMessageById(validId, aimId, content)
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	}
}

func MessageChat(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	strAimId := c.Query("to_user_id")
	aimId, _ := strconv.ParseInt(strAimId, 10, 64)
	validId, _ := service.GetTokenId(token)
	if validId == 0 {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	} else {
		messages, _ := service.GetMessageList(validId, aimId)
		c.JSON(http.StatusOK, ChatResponse{Response: model.Response{StatusCode: 0}, MessageList: messages})
	}
}
