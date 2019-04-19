package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	ly_consul "ly/consul-service"
	"ly/finance-service/proto/finance"
	"time"
)

//查询用户的余额grpc调用测试

var CMD_Q_Name string

func init() {
	flag.StringVar(&CMD_Q_Name, "name", "", "查询余额的姓名")
	flag.Parse()
}
func main() {

	if CMD_Q_Name == "" {
		log.Fatalln("查询名为空")
	}

	//从consul服务发现中获取userServerNode的地址
	conn_address := ly_consul.GetServiceConnect("financeServerNode")
	// 创建连接，没有经过tls加密
	conn, err := grpc.Dial(conn_address, grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	finance_client := finance.NewFinanceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resutle, err := finance_client.QueryMoney(ctx, &finance.User{
		Name: CMD_Q_Name,
	})

	if err != nil {
		log.Println(err)
	}

	log.Println("余额结果")
	log.Println(resutle.User.Name)
	log.Println(resutle.Money)
}
