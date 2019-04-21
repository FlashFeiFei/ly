package main

import (
	"context"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	ly_consul "ly/consul-service"
	pb "ly/finance-service/proto/finance"
	"net"
)

const (
	//consul心跳检测的端口
	consul_check_port = 8081
	//grpc服务的端口
	grpc_service_port = 50052
)

type Server struct {
	//账本
	AccountBook map[string]int64
}

//充值
func (s *Server) Recharge(ctx context.Context, in *pb.RechargeInfo) (*pb.Result, error) {
	user_name := in.User.Name
	money := in.Money
	if _, ok := s.AccountBook[user_name]; ok {
		//金额叠加
		s.AccountBook[user_name] = s.AccountBook[user_name] + money
	} else {
		//第一次
		s.AccountBook[user_name] = money
	}

	result := new(pb.Result)
	result.Result = "充值成功"
	result.RechargeInfo = &pb.RechargeInfo{
		Money: in.Money,
		User: &pb.User{
			Name: user_name,
		},
	}

	return result, nil
}

//查询余额
func (s *Server) QueryMoney(ctx context.Context, in *pb.User) (*pb.Money, error) {
	user_name := in.Name

	result_money := new(pb.Money)

	User := &pb.User{
		Name: user_name,
	}

	result_money.User = User

	if _, ok := s.AccountBook[user_name]; ok {
		//返回余额
		result_money.Money = s.AccountBook[user_name]

	} else {
		result_money.Money = 0
	}

	return result_money, nil
}

//服务发现

func registerServer() {

	//服务注册
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "financeServerNode"
	registration.Name = "financeServer"
	registration.Port = grpc_service_port
	registration.Tags = []string{"serverNode"}
	registration.Address = "118.25.8.93"
	//自己封装了一下定时检查的脚本
	ly_consul.RegisterServer(consul_check_port, registration)
}

func main() {
	//注册服务发现
	go registerServer()

	//grpc服务开启
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpc_service_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//初始化服务
	finance_server := &Server{
		AccountBook: make(map[string]int64),
	}

	//grpc调用注册
	pb.RegisterFinanceServer(s, finance_server)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
