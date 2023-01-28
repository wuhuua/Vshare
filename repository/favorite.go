package repository

import (
	"strconv"
	"sync"

	"github.com/Iscolito/Vshare/model"
)

type FavoriteDao struct {
}

var favoriteDao *FavoriteDao
var favoriteOnce sync.Once

func NewFavoriteDaoInstance() *FavoriteDao {
	favoriteOnce.Do(
		func() {
			favoriteDao = &FavoriteDao{}
		})
	return favoriteDao
}

func (*FavoriteDao) FavoriteAction(userId string, videoId string) error {
	var err error
	_, err = rdb[3].SAdd(userId, videoId).Result()
	if err != nil {
		return err
	}
	_, err = rdb[4].SAdd(videoId, userId).Result()
	if err != nil {
		return err
	}
	likeCount, _ := rdb[4].SCard(userId).Result()
	err = db.Model(&model.Video{}).Where("id = ?", videoId).Update("favoritecount", likeCount).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FavoriteDao) IsFavorite(userId string, videoId string) (bool, error) {
	return rdb[3].SIsMember(userId, videoId).Result()
}

func (*FavoriteDao) UnLikeAction(userId string, videoId string) error {
	var err error
	_, err = rdb[3].SRem(userId, videoId).Result()
	if err != nil {
		return err
	}
	_, err = rdb[4].SRem(videoId, userId).Result()
	if err != nil {
		return err
	}
	likeCount, _ := rdb[4].SCard(userId).Result()
	err = db.Model(&model.Video{}).Where("id = ?", videoId).Update("favoritecount", likeCount).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FavoriteDao) GetLikeList(userId string) ([]int64, error) {
	strLikeList, _ := rdb[3].SMembers(userId).Result()
	if strLikeList == nil {
		return nil, nil
	}
	likeList := make([]int64, len(strLikeList))
	var id int64
	for i, strId := range strLikeList {
		id, _ = strconv.ParseInt(strId, 10, 64)
		likeList[i] = id
	}
	return likeList, nil
}
