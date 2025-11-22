package main

import (
	"fmt"
	"gin-todo/config"
	"gin-todo/models"
	"gin-todo/routers"
	"os"
)

func main() {
	// 👇 加上这几行调试代码
	dir, _ := os.Getwd()
	fmt.Println("------------------------------------------------")
	fmt.Println("程序运行的当前目录是:", dir)
	fmt.Println("------------------------------------------------")
	// 1. 加载配置 (最先执行)
	config.InitConfig()
	// 2. 初始化数据库
	models.ConnectDatabase(config.AppConfig.Database.Dsn)

	// 3. 初始化路由
	r := routers.SetupRouter()

	// 4. 启动服务
	r.Run()
}
