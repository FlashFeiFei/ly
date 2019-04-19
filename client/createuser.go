package main

//创建用户的grpc调用测试
import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	ly_consul "ly/consul-service"
	"ly/user-service/proto/user"
	"time"
)

var CMD_C_Name string
var CMD_C_Sex string

func init() {
	//第一个参数是绑定到的变量中，第二个参数是命令行的参数，第三个参数是默认这，第四个参数是解释这个命令的作用
	flag.StringVar(&CMD_C_Name, "name", "", "创建用户姓名")
	flag.StringVar(&CMD_C_Sex, "sex", "", "创建用户性别")
	flag.Parse()
}

func main() {

	if CMD_C_Name == "" || CMD_C_Sex == "" {
		log.Fatalln("性别或者性别为空")
	}

	//从consul服务发现中获取userServerNode的地址
	conn_address := ly_consul.GetServiceConnect("userServerNode")
	// 创建连接，没有经过tls加密
	conn, err := grpc.Dial(conn_address, grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	user_client := user.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resutle, err := user_client.CreateUser(ctx, &user.UserInfo{
		Name: CMD_C_Name,
		Sex:  CMD_C_Sex,
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("创建用户结果")
	log.Println(resutle.Result)
	log.Println(resutle.UserInfo.Name)
	log.Println(resutle.UserInfo.Sex)
}
