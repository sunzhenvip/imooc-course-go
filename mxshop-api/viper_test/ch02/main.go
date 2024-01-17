package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// 测试和生产环境配置隔离

type ServerConfig struct {
	ServiceName string `mapstructure:"name"`
	Port        int    `mapstructure:"port"`
}



func main() {
	v := viper.New()
	// 文件路径如何设置
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)
}
