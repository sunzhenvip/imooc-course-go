package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := GetFreePort()
	if err != nil {
		panic(err)
	}
	fmt.Println(addr)
}

func GetFreePort() (addr int, err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
