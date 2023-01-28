package repository

import (
	"strconv"
	"sync"

	"github.com/Iscolito/Vshare/model"
)

type RelationDao struct {
}

var relationDao *RelationDao
var relationOnce sync.Once

func NewRelationDaoInstance() *RelationDao {
	relationOnce.Do(
		func() {
			relationDao = &RelationDao{}
		})
	return relationDao
}

func (*RelationDao) FollowAction(userId string, followId string) error {
	var err error
	_, err = rdb[1].SAdd(userId, followId).Result()
	if err != nil {
		return err
	}
	_, err = rdb[2].SAdd(followId, userId).Result()
	if err != nil {
		return err
	}
	followCount, _ := rdb[1].SCard(userId).Result()
	err = db.Model(&model.User{}).Where("id = ?", userId).Update("followcount", followCount).Error
	if err != nil {
		return err
	}
	followerCount, _ := rdb[2].SCard(followId).Result()
	err = db.Model(&model.User{}).Where("id = ?", followId).Update("followercount", followerCount).Error
	if err != nil {
		return err
	}
	return nil
}

func (*RelationDao) IsFollowed(userId string, aimId string) (bool, error) {
	return rdb[1].SIsMember(userId, aimId).Result()
}

func (*RelationDao) UnFollowAction(userId string, followId string) error {
	var err error
	_, err = rdb[1].SRem(userId, followId).Result()
	if err != nil {
		return err
	}
	_, err = rdb[2].SRem(followId, userId).Result()
	if err != nil {
		return err
	}
	followCount, _ := rdb[1].SCard(userId).Result()
	err = db.Model(&model.User{}).Where("id = ?", userId).Update("followcount", followCount).Error
	if err != nil {
		return err
	}
	followerCount, _ := rdb[2].SCard(followId).Result()
	err = db.Model(&model.User{}).Where("id = ?", followId).Update("followercount", followerCount).Error
	if err != nil {
		return err
	}
	return nil
}

func (*RelationDao) GetFollowList(userId string) ([]int64, error) {
	strFollowList, _ := rdb[1].SMembers(userId).Result()
	if strFollowList == nil {
		return nil, nil
	}
	followList := make([]int64, len(strFollowList))
	var id int64
	for i, strId := range strFollowList {
		id, _ = strconv.ParseInt(strId, 10, 64)
		followList[i] = id
	}
	return followList, nil
}

func (*RelationDao) GetFollowerList(userId string) ([]int64, error) {
	strFollowerList, _ := rdb[2].SMembers(userId).Result()
	if strFollowerList == nil {
		return nil, nil
	}
	followerList := make([]int64, len(strFollowerList))
	var id int64
	for i, strId := range strFollowerList {
		id, _ = strconv.ParseInt(strId, 10, 64)
		followerList[i] = id
	}
	return followerList, nil
}
