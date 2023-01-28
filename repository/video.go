package repository

import (
	"sync"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/util"
	"gorm.io/gorm"
)

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return videoDao
}

func (*VideoDao) GetVideoById(id int64) (*model.Video, error) {
	streams := model.Video{}
	err := db.Where("id = ?", id).Find(&streams).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find video by id err:" + err.Error())
		return nil, err
	}
	return &streams, nil
}

func (*VideoDao) GetVideos(num int) ([]model.Video, error) {
	streams := make([]model.Video, 0)
	err := db.Order("id desc").Limit(num).Find(&streams).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find videos err:" + err.Error())
		return nil, err
	}
	return streams, nil
}

func (*VideoDao) InitVideo(videoPath string, coverPath string, userId int64, videoName string) error {
	video := &model.Video{UserId: userId, PlayUrl: videoPath, CoverUrl: coverPath, VideoName: videoName}
	db.Create(video)
	return nil
}

func (*VideoDao) GetVideoList(userId int64) ([]model.Video, error) {
	streams := make([]model.Video, 0)
	err := db.Order("id desc").Where("userid = ?", userId).Find(&streams).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find lists err:" + err.Error())
		return nil, err
	}
	return streams, nil
}
