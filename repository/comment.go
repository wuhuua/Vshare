package repository

import (
	"sync"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/util"
	"gorm.io/gorm"
)

type CommentDao struct {
}

var commentDao *CommentDao
var commentOnce sync.Once

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}

func (*CommentDao) InitCommentByContent(text string, date string, userId int64, videoId int64) (*model.Comment, error) {
	comment := &model.Comment{Content: text, CreateDate: date, UserId: userId, VideoId: videoId}
	db.Create(comment)
	db.Model(&model.Video{}).Where("id = ?", videoId).Update("commentcount", gorm.Expr("commentcount + ?", 1))
	return comment, nil
}

func (*CommentDao) GetCommentByVideoId(videoId int64) ([]model.Comment, error) {
	comments := make([]model.Comment, 0)
	err := db.Order("id desc").Where("videoid = ?", videoId).Find(&comments).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find comments err:" + err.Error())
		return nil, err
	}
	return comments, nil
}
