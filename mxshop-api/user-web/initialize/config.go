package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/utils"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func GetConfigFileName() string {
	debug := GetEnvInfo(global.EnvMXSHOP)
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user-web/%s-pro.yaml", configFilePrefix)
	if !debug {
		configFileName = fmt.Sprintf("user-web/%s-debug.yaml", configFilePrefix)
	}
	return configFileName
}

func InitConfig() {
	InitSystemConfig()
	InitServerConfig()
}

func InitSystemConfig() {
	configFileName := GetConfigFileName()
	fmt.Println("配置文件名称打印：", configFileName)
	v := viper.New()
	// 文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 映射数据到全局变量
	if err := v.Unmarshal(global.SystemConfig); err != nil {
		panic(err)
	}
	// fmt.Println(global.ServerConfig)
	zap.S().Info("配置信息：&v", global.SystemConfig)

	// viper功能-动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		// fmt.Println("config file changed：", e.Name)
		zap.S().Infof("配置文件产生变化：%s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.SystemConfig)
		fmt.Printf("%#v", global.SystemConfig)

	})
}

func InitServerConfig() {
	configFileName := GetConfigFileName()
	fmt.Println("配置文件名称打印：", configFileName)
	// fmt.Println(configFileName)
	// fmt.Println(debug)

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
	// 线上模式走配置端口|线上随机端口
	if GetEnvInfo(global.EnvMXSHOP) {
		port, err := utils.GetFreePort()
		if err != nil {
			panic(err)
		}
		// 修改为随机获取端口
		global.ServerConfig.Port = port
	}
	// fmt.Println(global.ServerConfig)
	zap.S().Info("配置信息：&v", global.ServerConfig)
	fmt.Printf("%v", v.Get("user_srv.port")) // Viper可以通过key来直接获取值

	// viper功能-动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		// fmt.Println("config file changed：", e.Name)
		zap.S().Infof("配置文件产生变化：%s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
		fmt.Printf("%#v", global.ServerConfig)

	})
	// v.OnConfigChange()
	// 这个对象如何在其他文件中使用 - 全局变量

}
