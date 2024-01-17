package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 1、建立链接
	conn, _ := net.Dial("tcp", ":1234")
	var reply *string = new(string)
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err := client.Call("HelloService.Hello", "bobby", reply)
	// client.hello
	//
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(*reply)
}
