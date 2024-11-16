package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var Conf Config

func Load() {
	// yamlFile, err := os.ReadFile("./config/settings.yaml")
	// if err != nil {
	// 	fmt.Println("[config init error] os.ReadFile 配置文件读取失败 " + err.Error())
	// }
	// err = yaml.Unmarshal(yamlFile, &Conf)
	// if err != nil {
	// 	fmt.Println("[config init error] yaml.Unmarshal 配置文件解析失败 " + err.Error())
	// 	return
	// }
	// fmt.Println("[Success] 配置文件读取成功！！！")

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local" // 默认环境
	}

	viper.SetConfigName(env)      // 设置配置文件名称
	viper.AddConfigPath("config") // 设置配置文件路径
	viper.SetConfigType("yaml")   // 设置配置文件类型

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[ERROR] Reading config file, %s", err)
	}

	// 将配置文件内容反序列化到结构体
	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatalf("[ERROR] Unable to decode into struct, %v", err)
	}

	log.Println("[INFO] 配置文件读取成功！！！")
}
