package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1、建立链接
	client, _ := rpc.Dial("tcp", ":1234")
	var reply *string = new(string)
	err := client.Call("HelloService.Hello", "bobby", reply)

	// client.hello
	//
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(*reply)
}
