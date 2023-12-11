package launch

import "cavy/core/module"

type LaunchBase struct {
	*module.Module
}

func (launch *LaunchBase) Init() error {
	return nil
}
