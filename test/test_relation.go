package test

import (
	"fmt"

	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/service"
)

func follow_test(user string, follow string) {
	relationDao := repository.NewRelationDaoInstance()
	err := relationDao.FollowAction(user, follow)
	if err != nil {
		fmt.Println("error for redis set")
		return
	}
}

func isfollow_test(user string, follow string) bool {
	relationDao := repository.NewRelationDaoInstance()
	flag, err := relationDao.IsFollowed(user, follow)
	if err != nil {
		fmt.Println("read set error")
	}
	return flag
}

func unfollow_test(user string, follow string) {
	relationDao := repository.NewRelationDaoInstance()
	err := relationDao.UnFollowAction(user, follow)
	if err != nil {
		fmt.Println("error for redis unfollow")
		return
	}
}

func Test_relation(mysqlIp string, mysqlPassword string, redisIp string, redisPassword string) {
	repository.InitRedis(redisIp, redisPassword, 1)
	repository.InitRedis(redisIp, redisPassword, 2)
	repository.InitMySQL(mysqlIp, mysqlPassword)
	var flag bool
	follow_test("20230001", "20230002")
	follow_test("20230001", "20230002")
	fmt.Println("Successfully inserted")
	flag = isfollow_test("20250005", "20230002")
	fmt.Printf("no found result is:%+v\n", flag)
	flag = isfollow_test("20230001", "20230002")
	fmt.Printf("found result is:%+v\n", flag)
	unfollow_test("20230001", "20230002")
	unfollow_test("20230001", "20230002")
	fmt.Println("successfully unfollowed")
	followlist, _ := service.GetFollowListById(20230001)
	fmt.Printf("%+v", followlist)
}
