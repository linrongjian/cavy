package api

import (
	"net"
	"strconv"
	"strings"
)

// Service interface contains Start and Stop methods which are called
// when the service is started and stopped. The Init method is called
// before the service is started.
//
// The Start and Init methods must be non-blocking.
//
// Implement this interface and pass it to the Run function to start your program.
type Service interface {

	// Init is called before the program/service is started
	// This method must be non-blocking.
	Init(s *Conf) error

	// Start is called after Init. This method must be non-blocking.
	Start() error

	// Stop is called in response to syscall.SIGINT, syscall.SIGTERM, or when a
	// Service is stopped.
	Stop() error
}

// ServerType 服务器类型
type ServerType string

// MQServerInfo Mq服务器信息
type MQServerInfo struct {
	Addr     string
	Account  string
	Password string
}

// DBServerInfo 数据库信息
type DBServerInfo struct {
	Addr     string
	Account  string
	Password string
	DBName   string
}

// RedisServerInfo Redis服务器信息
type RedisServerInfo struct {
	Addr     string
	DataBase int
}

// StConfig 配置文件
type StConfig struct {
	Name         string          `json:"Name"`
	Listen       string          `json:"Listen"`
	RPCPort      string          `json:"RPCPort"`
	CertFile     string          `json:"CertFile"`
	KeyFile      string          `json:"KeyFile"`
	RedisServer  RedisServerInfo `json:"RedisServer"`
	MQServer     MQServerInfo    `json:"MQServer"`
	DBServer     DBServerInfo    `json:"DBServer"`
	ConsulAddr   string          `json:"ConsulAddr"`
	AuthAddr     string          `json:"AuthAddr"`
	RechargeAddr string          `json:"RechargeAddr"`
	Language 	 string          `json:"Language"`
}

// Conf 服务器配置信息
type Conf struct {
	Addr          string          //服务器监听地址
	Name          string          //服务器名称
	Type          ServerType      //服务器类型
	ConnRedis     bool            //连接redis		//默认不连接
	RedisServer   RedisServerInfo //redis服务器信息
	ConnMQ        bool            //连接MQ // 默认不连接
	MQServer      MQServerInfo    //MQ服务器信息
	ConnDB        bool            //连接数据库		//默认不连接
	DBServer      DBServerInfo    //数据库信息
	OrderTestMode bool            //支付测试模式开关 打开情况下只支付最小额度
	ConsulAddr    string          //服务注册地址
	CertFile      string
	Keyfile       string
	Version       string
	StdConf       StConfig
}

var (
	//AuthServer 认证服
	AuthServer ServerType = "auth"
	//RechargeServer 支付服务器
	RechargeServer ServerType = "recharge_xzy"
	//PushServer 推送服务器
	PushServer ServerType = "push"
)

// 服务器类型描述
var (
	typeMap = map[ServerType]string{
		AuthServer:     "auth",
		RechargeServer: "recharge_xzy",
		PushServer:     "push",
	}
)

// String 获得服务器类型名称
func (t ServerType) String() string {
	if _, ok := typeMap[t]; ok {
		return typeMap[t]
	}
	return "invalid"
}

// ConvertLocalAddr 转换成本地IP 格式 IP:PORT
func ConvertLocalAddr(addr string) string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	listenAddr := addr
	protIdx := strings.LastIndex(listenAddr, ":")
	return localAddr[0:idx] + listenAddr[protIdx:]
}

// GetIPAddr 获取服务器ip地址
func (c *Conf) GetIPAddr() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	host, _, _ := net.SplitHostPort(localAddr)
	return host
}

// GetIPPort 获取服务器端口
func (c *Conf) GetIPPort() int {
	_, port, _ := net.SplitHostPort(c.Addr)
	portInt, _ := strconv.Atoi(port)
	return portInt
}
