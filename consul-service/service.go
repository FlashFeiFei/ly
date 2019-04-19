package consul_service

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
)

//获取服务的链接地址
func GetServiceConnect(serviceID string) string {
	client, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		log.Fatal("consul 客户端client出错: ", err)
	}

	var services map[string]*consulapi.AgentService

	services, err = client.Agent().Services()

	if err != nil {
		log.Fatal("获取服务列表错误: ", err)
	}

	if _, found := services[serviceID]; !found {
		log.Fatal("服务没有找到")
	}

	service := services[serviceID]

	return fmt.Sprintf("%s:%d", service.Address, service.Port)
}
