package repository

import (
	"sync"

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

func (*VideoDao) GetVideoById(id int64) (*Video, error) {
	streams := Video{}
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

func (*VideoDao) GetVideos(num int) ([]Video, error) {
	streams := make([]Video, 0)
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
