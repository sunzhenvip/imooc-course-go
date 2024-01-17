package initialize

import (
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop-api/user-web/global"
	__proto "mxshop-api/user-web/proto"
)

func InitSrvConn() {
	SrvConnRoundRobin()
}

// SrvConnDefault 默认模式
func SrvConnDefault() {
	cfg := consulApi.DefaultConfig()
	consulInfo := fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	cfg.Address = consulInfo
	userSrvHost := ""
	userSrvPort := 0
	client, err := consulApi.NewClient(cfg)
	// 查找数据
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}
	// 目前只获取一个
	for _, val := range data {
		userSrvHost = val.Address
		userSrvPort = val.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("用户服务连接失败")
	}
	// 拨号连接
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GteUserList]连接 用户服务失败", "msg", err.Error())
	}
	// 1、可能出现的问题 后续用户服务下线了， 修改端口了 3、改IP了之类的问题  负载均衡来做
	// 2、已经事先创立好了链接，这样后续不用进行tcp的三次握手
	// 3、一个连接多个groutine公用 性能。-连接池
	global.UserSrvClient = __proto.NewUserClient(userConn)
}

// SrvConnRoundRobin 负载均衡模式
func SrvConnRoundRobin() {
	conn, err := grpc.Dial(
		"consul://127.0.0.1:8500/user_srv?wait=14s&tag=srv",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), // 写法固定
	)
	if err != nil {
		panic(err.Error())
	}
	global.UserSrvClient = __proto.NewUserClient(conn)
}
