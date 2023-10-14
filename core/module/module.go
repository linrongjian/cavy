package module

type IModule interface {
	Init() error
	Start() error
	Stop() error
	Destroy() error
	Load() error
	Save() error
}

type Module struct {
}

func (svr *Module) Init() error {
	return nil
}

func (svr *Module) Start() error {
	return nil
}

func (svr *Module) Stop() error {
	return nil
}

func (svr *Module) Destroy() error {
	return nil
}

func (svr *Module) Load() error {
	return nil
}

func (svr *Module) Save() error {
	return nil
}
