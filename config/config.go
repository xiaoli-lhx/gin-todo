package config

import (
	"log"

	"github.com/spf13/viper"
)

// 定义结构体来映射 config.yaml 的内容
type Config struct {
	Database struct {
		Dsn string `mapstructure:"dsn"`
	} `mapstructure:"database"`
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
}

// 全局变量，用来存放读取到的配置
var AppConfig *Config

// 初始化配置
func InitConfig() {
	viper.SetConfigName("config") // 配置文件名 (不带后缀)
	viper.SetConfigType("yaml")   // 文件类型
	viper.AddConfigPath(".")      // 搜索路径 (当前目录)
	// 2. 搜索上一级目录 (适用于进入 main 目录后运行 go run main.go)
	viper.AddConfigPath("../")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 把读取到的配置解析到结构体变量中
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
}
