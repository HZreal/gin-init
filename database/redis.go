package database

import (
	"fmt"
	"gin-init/config"
	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	// 连接 Redis, 获得 DB 实例
	client := redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.GetAddr(),
		Password: config.Conf.Redis.Password, // no password set
		DB:       config.Conf.Redis.DB,       // use default DB
	})
	fmt.Println("[Success] Redis数据库连接成功！！！")

	//
	RDB = client
}
