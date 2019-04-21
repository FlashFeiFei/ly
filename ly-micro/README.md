# 在windows中交叉编译linux可执行文件，在centos中搭建运行环境

# 一些特殊要求
因为我比较懒和菜，服务器的ip我是在代码中写死的，并没有写个配置文件去配置他，所以，要想测试，需要改变代码中的连接ip

## 技术栈
- docker-compose

## docker镜像

consul 镜像是谷歌的，看不懂这个

golang 的镜像是docker.hub 

## 目录结构

>ly-micro
>>consul 服务发现容器
>>>docker-compose.yml 服务发现consul启动配置

>>finance-service 财务服务
>>>bin 财务服务golang的二进制文件

>>>docker-compose.yml 财务服务goalng运行容器

>>user-service 用户服务
>>>bin 用户服务golang的二进制文件

>>>docker-compose.yml 用户服务golang运行容器



## 在windows中构建交叉编译Linux的代码

编译充值服务

```
/finance-service/go-build.bat

手动将编译好的finance-service重命名为micro  ##呵呵，因为我不懂写命令，所以暴力的做法鼠标右键
```

编译用户服务

```
/user-service/go-build.bat

手动将编译好的user-service重命名为micro  ##同理可证，我还是不懂写命令
```

### linux上启动服务

# 上传代码（呵呵，我没有配置gitlab，暴力做法）
 
将刚刚编译好的finance-service和user-service都分别放到linux上的finance-service/bin和user-service/bin目录。**注意文件名最后都是叫micro**因为我在docker-compose写死了，高级操作不会，超出了我的范围

## 运行项目

1. 创建一个项目公用的网络

```
docker network create consul_ly
```

2. 先运行服务发现

```
进入 ly-micro/consul

docker-compose up -d
```

3. 运行用户服务

```
进入 ly-micro/user-service

docker-compose up -d
```


4. 运行财务服务

```
进入 ly-micro/finance-service

docker-compose up -d
```

### 测试

然后就可以运行，client/createuser.go 和 client/queryfinance.go 了



### docker-compose的一些命令
```
docker-compose up -d             # 启动容器
docker-compose down              # 关闭容器
docker-compose run 服务名字 bash #运行容器，并进入容器里面
docker-compose logs 服务名字     #查看容器的日志，自己的理解，是可以查看容器里面的程序的输出流输出的东西 
```