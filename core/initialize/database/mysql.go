package database

import (
	"gin-init/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitMysql() {
	// 连接 Mysql, 获得 DB 实例
	conn, err := gorm.Open(mysql.Open(config.Conf.Mysql.GetDsn()), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Println("[ERROR] 连接 Mysql 数据库失败, error=" + err.Error())
		return
	}

	// 连接成功
	log.Println("[INFO] Mysql数据库连接成功！！！")

	// 迁移
	// TODO UnitInfo 迁移有问题
	// err = conn.AutoMigrate(&entity.UserModel{}, &entity.UnitInfo{})
	// if err != nil {
	// 	fmt.Println("[database mysql] 创建表失败！")
	// }

	//
	db, err := conn.DB()
	if err != nil {
		log.Println("[ERROR] 获取sql实例失败！")
	}
	db.SetMaxIdleConns(config.Conf.Mysql.MaxConn)
	db.SetMaxOpenConns(config.Conf.Mysql.MaxOpen)

	//
	DB = conn
}
