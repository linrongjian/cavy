# ServerEngin

基于go语言的高性能服务器框架


## Core

### Network

- Client

  通信客户端
  

- Server

  通信服务端
  

- Transport

  通信接口基类
  

- Protocols

  支持协议类型
  

	- grpcwrap
	- httpwrap
	- mqwrap
	- tcpwrap
	- udpwrap

### Router

通信路由


### Store

数据存储


### Baseserver

服务器基类


- Run()
- Stop()

## Service

服务组件


### Gateway

网关组件


### Business

业务逻辑组件


### Center

中控组件


### Backend

后台管理组件


### Auth

账号组件


