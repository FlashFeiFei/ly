package consul_service

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"net/http"
)

//把服务，注册到服务发现中
func RegisterServer(consul_check_port int, registration *consulapi.AgentServiceRegistration) {
	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)

	if err != nil {
		log.Fatal("consul client error :", err)
	}
	//定时检查
	registration.Check = &consulapi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, consul_check_port, "/check"),
		Timeout:                        "3s",  //连接超时3秒
		Interval:                       "5s",  //5秒心跳检测一次
		DeregisterCriticalServiceAfter: "30s", //check失败后30秒删除服务
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("注册服务发生错误:", err)
	}

	//心跳检测接口检测
	http.HandleFunc("/check", consulCheck)
	//http服务器开始
	http.ListenAndServe(fmt.Sprintf(":%d", consul_check_port), nil)
}

//consul心跳检测函数
func consulCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "consulCheck")
}
