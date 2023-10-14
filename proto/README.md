# 服务描述
```
公共Protobuf组件
```

区间分配：
```
BusinessServer: 1-500
PushServer: 501-700
```

文件说明：
```
common.proto 存放最外层的proto统一结构
message_code.proto 存放cmd命令号，命名格式"OP+消息名"。cmd消息用户推送消息，正常req-rsp类型Http消息不通过cmd来解析。
business.proto 存放BusinessServer模块使用的proto
push.proto 存放PushServer模块使用的proto
```

公共说明:
```
固定签名为Key: EE7a1c5bc548e542GBFc340c531657F4
所有Http请求的返回需要先使用 HTTPResponse 结构解析 然后判断 result 结果码
成功再使用业务结构解析 HTTPResponse 的 data 字段，
失败可以打印 HTTPResponse 结构中的 msg 显示错误描述
```

# 协议说明

## 策略相关

### 获得配置 

* 支持版本号:
 * 版本号 **1.0.0.1** => **a45b9b25e0a29f3477a39e0462c3475d**
 * 版本号 **1.0.0.2** => **540da50d8fa6bbf58182d65e8e3d6e7d**
* 请求地址模板: ***https://{0}/schedule/{1}***<br>
    {0} -- 固定的服务器域名地址或者文件服务器地址<br>
	{1} -- 版本号的md5<br>
* 返回业务结构 -- **Schedule**

## 验证相关

### 微信认证

**使用POST请求**
* 请求地址模板: ***https://{0}/{Auth}/auth/login?sign={1}***<br>
	{0} -- Nginx地址<br>
    {Auth} -- 通过模块ID获得的路径<br>
    {1} -- 签名 string.tolower(md5.sum({}))<br>
* 请求结构 -- **AuthReq**
* 返回没有数据结构，只有成功或失败

### 微信认证

**使用POST请求**
* 请求地址模板: ***https://{0}/{Auth}/auth/wxpubliclogin?sign={1}***<br>
	{0} -- Nginx地址<br>
    {Auth} -- 通过模块ID获得的路径<br>
    {1} -- 签名 string.tolower(md5.sum({}))<br>
* 请求结构 -- **AuthReq**
* 返回没有数据结构，只有成功或失败

### 玩吧认证

**使用POST请求**
* 请求地址模板: ***https://{0}/{Auth}/auth/wanbalogin?sign={1}***<br>
	{0} -- Nginx地址<br>
    {Auth} -- 通过模块ID获得的路径<br>
    {1} -- 签名 string.tolower(md5.sum({}))<br>
* 请求结构 -- **AuthReq**
* 返回没有数据结构，只有成功或失败

## 登录相关

**使用POST请求** 
* 请求地址模板: ***https://{0}/{Hall}/login?sign={1}***<br>
	{0} -- Nginx地址<br>
    {Hall} -- 通过模块ID获得的路径<br>
    {1} -- 签名 string.tolower(md5.sum({}))<br>
* 请求结构 -- **LoginReq**
* 返回结构 -- **LoginRsp**