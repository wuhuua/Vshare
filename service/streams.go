package service

import (
	"errors"
	"strconv"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/util"
)

const coverName = "vshare.png"
const videoServer = "192.168.3.9:8080"

func GetStreams() ([]model.Video, error) {
	VideoDao := repository.NewVideoDaoInstance()
	streams, err := VideoDao.GetVideos(30)
	if err == nil {
		if streams == nil {
			util.Logger.Error("no videos")
			return nil, errors.New("no videos")
		} else {
			userDao := repository.NewUserDaoInstance()
			var author *model.User
			for i, _ := range streams {
				author, _ = userDao.GetUserById(streams[i].UserId)
				streams[i].Author = *author
			}
			return streams, nil
		}
	} else {
		util.Logger.Error("find videos process error")
		return nil, errors.New("find videos process error")
	}
}

func GetStreamsById(userId int64) ([]model.Video, error) {
	VideoDao := repository.NewVideoDaoInstance()
	streams, err := VideoDao.GetVideos(30)
	if err == nil {
		if streams == nil {
			util.Logger.Error("no videos")
			return nil, errors.New("no videos")
		} else {
			userDao := repository.NewUserDaoInstance()
			favoriteDao := repository.NewFavoriteDaoInstance()
			var author *model.User
			for i, _ := range streams {
				author, _ = userDao.GetUserById(streams[i].UserId)
				streams[i].Author = *author
				streams[i].IsFavorite, _ = favoriteDao.IsFavorite(strconv.FormatInt(userId, 10), strconv.FormatInt(streams[i].Id, 10))
			}
			return streams, nil
		}
	} else {
		util.Logger.Error("find videos process error")
		return nil, errors.New("find videos process error")
	}
}

func PublishVideo(userId int64, videoName string, fileName string) {
	videoDao := repository.NewVideoDaoInstance()
	videoDao.InitVideo("http://"+videoServer+"/stream/"+fileName, "http://"+videoServer+"/stream/"+coverName, userId, videoName)
}

func GetStreamList(userId int64) ([]model.Video, error) {
	VideoDao := repository.NewVideoDaoInstance()
	streams, err := VideoDao.GetVideoList(userId)
	if err != nil {
		util.Logger.Error("find video list process error")
		return nil, errors.New("find video list process error")
	}
	if streams == nil {
		util.Logger.Error("lists are empty")
		return nil, errors.New("lists are empty")
	}
	userDao := repository.NewUserDaoInstance()
	author, _ := userDao.GetUserById(userId)
	for i := range streams {
		streams[i].Author = *author
	}
	return streams, nil
}
