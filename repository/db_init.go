package repository

import (
	"fmt"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var rdb [7]*redis.Client

func InitMySQL(ip string, password string) error {
	var err error
	dsn := fmt.Sprintf("root:%s@tcp(%s)/vshare?charset=utf8mb4&parseTime=True&loc=Local", password, ip)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

func InitRedis(ip string, password string, baseNum int) error {
	var err error
	rdb[baseNum] = redis.NewClient(&redis.Options{Addr: ip, Password: password, DB: baseNum})
	_, err = rdb[baseNum].Ping().Result()
	return err
}
