package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 是一个全局变量，首字母大写以便其他包访问
var DB *gorm.DB

func ConnectDatabase(dsn string) {
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("连接数据库失败:%v", err))
	}
	// 自动迁移
	database.AutoMigrate(&Todo{})
	// 赋值给全局变量
	DB = database
}
