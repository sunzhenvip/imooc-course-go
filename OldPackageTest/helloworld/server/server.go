package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	// 返回值是通过修改reply的值

	// return nil
	*reply = "hello," + request
	return nil
}

func main() {
	// rpc 三步走
	// 1、实例化一个server
	// 2、注册处理逻辑handler
	listener, _ := net.Listen("tcp", ":1234")

	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 3、启动服务
	conn, _ := listener.Accept() // 当一个链接进来的时候
	rpc.ServeConn(conn)
	// 一连串的代码大部分都是net包好像和rpc没有关系
	// rpc 调用需要解决几个问题 1、call id 2、序列化和反序列化
}
