package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"time"
)

const conSulHost = "192.168.31.90:8500"

func Register(address string, port int, name string, tags []string, id string) {
	cfg := api.DefaultConfig()
	cfg.Address = conSulHost

	client, err := api.NewClient(cfg)

	if err != nil {
		panic(err)
	}
	// 生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP: fmt.Sprintf("http://%s:%d", "192.168.31.90", port), // 需要写实际IP HTTP模式监控检查
		// GRPC: address, // grpc 模式健康检查
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}

func AllServices() {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}

func FilterService() {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter(`Service == "user-web"`)
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}

func main() {
	Register("192.168.31.90", 9000, "user-web", []string{"mxshop", "bobby"}, "user-web")
	time.Sleep(time.Hour)
	// AllServices()
	// FilterService()
}
