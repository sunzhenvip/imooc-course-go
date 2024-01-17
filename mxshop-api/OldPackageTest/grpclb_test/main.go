package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	__proto "mxshop-api/user-web/proto"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important

	"google.golang.org/grpc"
)

// 负载均衡测试
func main() {
	conn, err := grpc.Dial(
		"consul://127.0.0.1:8500/user_srv?wait=14s&tag=srv",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), // 写法固定
	)
	if err != nil {
		log.Fatal(err)
	}
	userSrvClient := __proto.NewUserClient(conn)
	for i := 0; i < 10; i++ {
		rsp, err := userSrvClient.GetUserList(context.Background(), &__proto.PageInfo{
			Pn:    1,
			PSize: 2,
		})
		fmt.Println(rsp)
		fmt.Println(err)
	}
	defer conn.Close()

}
