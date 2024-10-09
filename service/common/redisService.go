package common

/**
 * @Author elastic·H
 * @Date 2024-10-09
 * @File: redisService.go
 * @Description:
 */

import (
	"gin-init/database"
	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	Client *redis.Client
}

func NewRedisService() *RedisService {
	return &RedisService{
		Client: database.RDB,
	}
}
