# FastGameServer(高性能游戏服务器框架)

## 游戏逻辑服:LogicServer

### 独立数据业务服

### 公共数据业务服

- 读写锁
- 防死锁

### 继承GameServer

## 游戏登录服:LoginServer

### 分配游戏服

### Token

### 账号

- 生成账号

### 验证session

### 继承GameServer

## 游戏网关服:GateServer

### session管理

- 接受消息
- 踢下线
- 创建用户
- 消息推送
- 查找用户
- 删除用户

### 游戏服管理

- 消息路由

### main()

### 消息处理:Handler()

### 路由对象:Router

### 继承GameServer

## 路由模块:Router

### 所有服务器:servers:Map

### 路由策略:Select

### 网关服与游戏服务绑定

## Nginx

### 网关服路由

### 登录服路由

## Consul

### 网关进程监控

### 登陆服进程监控

### 游戏服进程监控

## 数据:Model

### 内存缓存

- 内存数据结构

### redis

- 主从

### 存储

- 定时回写磁盘
- 接口回写磁盘
- 接口回写redis

### mysql

- 分库分表

## 通信服务端:Server

### 服务器配置:serverConf

### 所有客户端连接:conns:sync.Map

### 传输处理器:Transport

### 停止服务:Shutdown()

### 启动服务:Start()

- 监听客户端连接:Transport.Listen()

## 通信客户端:Client

### 客户端配置:clientConf

### 传输处理器:Transport

### 连接服务器:Connect()

- 连接服务端:Transport.connection

## 通信:Network

### websocket模块

### grpc模块

### rabbitMQ模块

### http模块

- gin
- 原生

### tcp模块

### udp模块

### 传输处理器:Transport

- 监听:Listen()
- 连接:Connect()
- 通道对象:Channel

### 通道:Channel

- 接受数据:Recv()

	- 接受消息缓存队列:recvQueue

- 发送数据:Send()

	- 发送消息队列:sendQueue
	- 发送失败队列:sendFailQueue

- 关闭:Close()

## 游戏服:GameServer

### redisConn

### mysqlConn

### logger

### run()

- 服务:Server
- 客户端:Clien

