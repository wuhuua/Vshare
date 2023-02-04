package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// Configuration 项目配置
type Configuration struct {
	// 主机地址
	HostIp string `json:"hostIp"`
	// MySQL Ip地址
	MysqlIp string `json:"mysqlIp"`
	// MySQL密码
	MysqlPassword string `json:"mysqlPassword"`
	// Redis Ip地址
	RedisIp string `json:"redisIp"`
	// Redis密码
	RedisPassword string `json:"redisPassword"`
	// 静态文件(如视频)存储Ip
	StaticIp string `json:"staticIp"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 给配置赋默认值
		config = &Configuration{
			HostIp:        "127.0.0.1:8080",
			MysqlIp:       "127.0.0.1:3306",
			MysqlPassword: "12345678",
			RedisIp:       "127.0.0.1:6379",
			RedisPassword: "",
			StaticIp:      "127.0.0.1:8080",
		}

		// 判断配置文件是否存在，存在直接JSON读取
		_, err := os.Stat("config.json")
		if err == nil {
			f, err := os.Open("config.json")
			if err != nil {
				log.Fatalf("open config err: %v", err)
				return
			}
			defer f.Close()
			encoder := json.NewDecoder(f)
			err = encoder.Decode(config)
			if err != nil {
				log.Fatalf("decode config err: %v", err)
				return
			}
		}

	})
	return config
}
