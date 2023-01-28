package service

import (
	"errors"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/util"
)

func PublishComment(text string, userId int64, videoId int64) (int64, string, error) {
	commentDao := repository.NewCommentDaoInstance()
	comment, _ := commentDao.InitCommentByContent(text, util.GetDate(), userId, videoId)
	return comment.Id, comment.CreateDate, nil
}

func GetCommentList(videoId int64) ([]model.Comment, error) {
	commentDao := repository.NewCommentDaoInstance()
	comments, err := commentDao.GetCommentByVideoId(videoId)
	if err != nil {
		util.Logger.Error("find videos process error")
		return nil, errors.New("find videos process error")
	}
	userDao := repository.NewUserDaoInstance()
	var user *model.User
	for i, _ := range comments {
		user, _ = userDao.GetUserById(comments[i].UserId)
		comments[i].User = *user
	}
	return comments, nil
}
