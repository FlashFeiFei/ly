# 学习如何做微服务，rpc调用

## 技术栈
- rpc框架grpc
- 服务发现consul


## 系统概述

主要有两个服务

- 用户服务
- 财务服务


用户服务：注册一个用户，并且默认给新用户充值100元（其实是1元，显示出来的100其实是单位为分）

财务服务：可以查询注册的用户，还有多少余额


## 运行流程

### 开启服务发现consul

```
consul agent -dev
```

### 运行用户服务

```
go run user-service/main.go
```

### 运行财务服务

```
go run finance-service/main.go
```


### 调用用户服务，创建一个用户

```
go run client/createuser.go -name 小明 -sex 男
```

### 调用财务服务，查询余额

```
go run client/queryfinance.go -name 小明
```


