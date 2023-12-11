package server

import "cavy/core/module"

type SvrBase struct {
	*module.ModuleContainer
}

func (svr *SvrBase) Init() error {
	return nil
}

func (svr *SvrBase) Start() error {
	return nil
}

func (svr *SvrBase) Stop() error {
	return nil
}

func (svr *SvrBase) Destroy() error {
	return nil
}

func (svr *SvrBase) Load() error {
	return nil
}

func (svr *SvrBase) Save() error {
	return nil
}
