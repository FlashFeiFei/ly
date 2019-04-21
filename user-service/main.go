package main

import (
	"context"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	ly_consul "ly/consul-service"
	"ly/finance-service/proto/finance"
	pb "ly/user-service/proto/user"
	"net"
	"time"
)

const (
	//consul心跳检测的端口
	consul_check_port = 8080
	//grpc服务的端口
	grpc_service_port = 50051
)

type Server struct {
	//用户记录
	User map[string]bool
}

//创建服务
func (s *Server) CreateUser(ctx context.Context, in *pb.UserInfo) (*pb.Result, error) {

	user_name := in.Name
	result := new(pb.Result)

	user_info := new(pb.UserInfo)
	user_info.Name = user_name
	user_info.Sex = in.Sex

	//用户信息
	result.UserInfo = user_info

	if _, ok := s.User[user_name]; !ok {
		//不存在的时候录取数据,默认充值1元
		s.User[user_name] = true
		_, err := s.financeService(user_name)

		if err != nil {
			//默认充值失败
			return nil, err
		}

	}

	return result, nil
}

//默认充值1元
func (s *Server) financeService(user_name string) (*finance.Result, error) {
	//获取financeServerNode服务的链接地址
	conn_address := ly_consul.GetServiceConnect("financeServerNode")
	// 发送请求，没有经过tls加密
	conn, err := grpc.Dial(conn_address, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	finance_client := finance.NewFinanceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//默认充值1元
	result, err := finance_client.Recharge(ctx, &finance.RechargeInfo{
		User: &finance.User{
			Name: user_name,
		},
		Money: 100,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

//服务发现
func registerServer() {
	//服务注册
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "userServerNode"
	registration.Name = "userServer"
	registration.Port = grpc_service_port
	registration.Tags = []string{"serverNode"}
	registration.Address = "118.25.8.93"
	//自己封装了一下定时检查的脚本
	ly_consul.RegisterServer(consul_check_port, registration)
}

func main() {

	//注册服务到服务发现consul中
	go registerServer()

	//grpc服务开启
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpc_service_port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user_server := &Server{
		User: make(map[string]bool),
	}
	pb.RegisterUserServer(s, user_server)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
