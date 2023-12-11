package consul

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	////"battlePlatform/wjrgit.qianz.com/common/api"	//"battlePlatform/wjrgit.qianz.com/common/hook"

	"github.com/linrongjian/cavy/common/api"
	"github.com/linrongjian/cavy/common/hook"

	"github.com/linrongjian/cavy/common/mlog"

	consul "github.com/hashicorp/consul/api"
)

var client *consul.Client

var (
	// ErrRegisterFailed 服务注册失败
	ErrRegisterFailed = errors.New("consul: register server failed")
)

func init() {
	log := mlog.NewLogger(mlog.Fields{})

	hook.AddHook(func(s *api.Conf) {
		if s.ConsulAddr == "" {
			log.Warn("Not find consul Addr")
			return
		}
		config := consul.DefaultConfig()
		config.Address = s.ConsulAddr
		c, err := NewClient(config)
		if err != nil {
			panic(err)
		}
		client = c

		log.Info("Connect Consul")
	})
	initConf()
}

// NewClient 新建consul客户端
func NewClient(config *consul.Config) (*consul.Client, error) {
	if config == nil {
		config = consul.DefaultConfig()
	}
	c, err := consul.NewClient(config)
	return c, err
}

// GetClient 获取cosul客户端
func GetClient() *consul.Client {
	return client
}

func getVersion(entry map[string]string) string {
	return entry["Version"]
}

// RegisterToConsul 服务注册到consul
func RegisterToConsul(s *api.Conf) error {
	if client == nil {
		return nil
	}

	startTime := time.Now()

	log := mlog.NewLogger(mlog.Fields{
		"ID":      fmt.Sprintf("%s-%s:%d", s.Name, s.GetIPAddr(), s.GetIPPort()),
		"Name":    s.Name,
		"Port":    s.GetIPPort(),
		"Address": s.GetIPAddr(),
	})

	// 创建一个新服务
	registration := &consul.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s:%d", s.Name, s.GetIPAddr(), s.GetIPPort()),
		Name:    s.Name,
		Tags:    []string{fmt.Sprintf("startTime:%d", startTime.UnixNano())},
		Port:    s.GetIPPort(),
		Address: s.GetIPAddr(),
		//Meta:    map[string]string{"Versopn": s.Version},
	}

	// 增加check
	ck := &consul.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("https://%s%s", fmt.Sprintf("%s:%d", s.GetIPAddr(), s.GetIPPort()), "/health"),
		Interval:                       "3s",   // 检查间隔
		Timeout:                        "5s",   // 响应超时时间
		DeregisterCriticalServiceAfter: "300s", // 注销节点超时时间
		TLSSkipVerify:                  true,
	}
	// 注册check服务
	registration.Check = ck

	log.WithFields(mlog.Fields{"Check": ck.HTTP})

	// 向consul注册服务
	if err := client.Agent().ServiceRegister(registration); err != nil {
		log.Error(err)
		return ErrRegisterFailed
	}

	log.Infof("register %s to consul success", s.Name)
	return nil
}

// GetServiceAddress 获取服务地址(随机)
func GetServiceAddress(name string) string {
	log := mlog.NewLogger(mlog.Fields{
		"serviceName": name,
	})

	services, err := client.Agent().Services()
	if err != nil {
		log.Errorf("get consul services err:%v", err)
		return ""
	}
	if len(services) == 0 {
		log.Errorf("services count is zero")
		return ""
	}
	srvs := make([]*consul.AgentService, 0, len(services))
	for id, item := range services {
		if !strings.HasPrefix(id, name) {
			continue
		}
		srvs = append(srvs, item)
	}
	if len(srvs) > 0 {
		srv := srvs[rand.Intn(len(srvs))]
		return fmt.Sprintf("%s:%d", srv.Address, srv.Port)
	}
	return ""
}
