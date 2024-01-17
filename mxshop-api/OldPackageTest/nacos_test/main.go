package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"mxshop-api/user-web/config"
	"time"
)

func main() {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         "83c8fd49-81bd-4bb1-a289-1283c23a3fa4", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		// RotateTime:          "1h",
		// MaxAge:              3,
		LogLevel: "debug",
	}

	// 创建服务发现客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		constant.KEY_CLIENT_CONFIG:  clientConfig,
		constant.KEY_SERVER_CONFIGS: serverConfigs,
	})

	//  获取Config
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev",
	})
	if err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	json.Unmarshal([]byte(content), &serverConfig)
	configClient.ListenConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			// fmt.Println(data)
			json.Unmarshal([]byte(data), &serverConfig)
			fmt.Sprintf("%v", serverConfig)
		},
	})

	fmt.Println(content)
	time.Sleep(time.Hour)
}
