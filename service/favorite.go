package service

import (
	"errors"
	"strconv"

	"github.com/Iscolito/Vshare/model"
	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/util"
)

func LikeById(userId int64, videoId int64) error {
	favoriteDao := repository.NewFavoriteDaoInstance()
	return favoriteDao.FavoriteAction(strconv.FormatInt(userId, 10), strconv.FormatInt(videoId, 10))
}

func UnLikeById(userId int64, videoId int64) error {
	favoriteDao := repository.NewFavoriteDaoInstance()
	return favoriteDao.UnLikeAction(strconv.FormatInt(userId, 10), strconv.FormatInt(videoId, 10))
}

func GetLikeListById(userId int64) ([]model.Video, error) {
	favoriteDao := repository.NewFavoriteDaoInstance()
	idList, _ := favoriteDao.GetLikeList(strconv.FormatInt(userId, 10))
	if idList == nil {
		util.Logger.Error("get likelist error")
		return nil, errors.New("get likelist error")
	}
	videoList := make([]model.Video, len(idList))
	videoDao := repository.NewVideoDaoInstance()
	userDao := repository.NewUserDaoInstance()
	user, _ := userDao.GetUserById(userId)
	var video *model.Video
	for i, id := range idList {
		video, _ = videoDao.GetVideoById(id)
		videoList[i] = *video
		videoList[i].IsFavorite = true
		videoList[i].Author = *user
	}
	return videoList, nil
}
