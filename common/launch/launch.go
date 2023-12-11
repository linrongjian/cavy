package launch

import "github.com/linrongjian/cavy/core/module"

type LaunchBase struct {
	*module.Module
}

func (launch *LaunchBase) Init() error {
	return nil
}
