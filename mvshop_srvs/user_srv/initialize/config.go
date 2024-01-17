package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mvshop_srvs/user_srv/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user_srv/%s-pro.yaml", configFilePrefix)
	if !debug {
		// /Users/sunzhen/go_dev/goproject/src/imooc/mvshop_srvs/user_srv/config-debug.yaml
		// user-srv/%s-debug.yaml
		configFileName = fmt.Sprintf("user_srv/%s-debug.yaml", configFilePrefix)
	}
	// fmt.Println("\r\n孙震")
	// fmt.Println(configFileName)

	v := viper.New()
	// 文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 这个对象如何在其他文件中使用
	// serverConfig := config.ServerConfig{}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}

	// fmt.Println(global.ServerConfig)
	zap.S().Info("配置信息：&v", global.ServerConfig)
	// fmt.Printf("%v", v.Get("name"))

	// viper都功能-动态监控变化
	// v.WatchConfig()
	// v.OnConfigChange(func(e fsnotify.Event) {
	// 	// fmt.Println("config file changed：", e.Name)
	// 	zap.S().Infof("配置文件产生变化：%s", e.Name)
	// 	_ = v.ReadInConfig()
	// 	_ = v.Unmarshal(global.ServerConfig)
	// 	fmt.Println(global.ServerConfig)
	// })
	// 这个对象如何在其他文件中使用 - 全局变量
}
