package service

import (
	"gin-init/database"
	"gorm.io/gorm"
)

var db *gorm.DB

// var rdb *redis.Client

// var ctx context.Context

func init() {
	db = database.DB
	// rdb = database.RDB
	// ctx = database.Ctx
}
