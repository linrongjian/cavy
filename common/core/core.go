package core

import (
	"cavy/common/api"
)

type rechargeCore struct {
}

// NewService 创建服务
func NewService() api.Service {
	return new(rechargeCore)
}

func (c *rechargeCore) Init(s *api.Conf) error {

	return nil
}

func (c *rechargeCore) Start() error {
	return nil
}

func (c *rechargeCore) Stop() error {
	return nil
}
