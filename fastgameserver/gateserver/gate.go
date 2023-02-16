package server

import (
	"fastgameserver/fastgameserver/store/mysql"
	"fastgameserver/fastgameserver/store/redis"
	"os"
	"os/signal"
	"sync"
	"trainserver/cmd"
	"trainserver/module"
	"trainserver/util/logger"
	"trainserver/zlgame_rpc/mq"
)

type Option func(*Options)

type service struct {
	opts Options
	sync.Mutex
	once sync.Once
}

func newService(opts ...Option) ZLGame {
	service := new(service)
	options := newOptions(opts...)

	service.opts = options

	return service
}

func (s *service) Name() string {
	return "zlgame"
}

func (s *service) Init(opts ...Option) {
	for _, o := range opts {
		o(&s.opts)
	}

	cmd.AddFlags(defaultFlags)
	cmd.AddFlags(redis.Flags)
	cmd.AddFlags(mq.Flags)
	cmd.AddFlags(mysql.Flags)
	if err := s.opts.Cmd.Init(); err != nil {
		logger.Fatal(err)
	}

	//mq.Startup()
	mysql.Startup()

	err := redis.Connect()
	if err != nil {
		logger.Fatal(err)
	}
	s.opts.Rds = redis.S()
}

func (s *service) Options() Options {
	return s.opts
}

func (s *service) String() string {
	return "zlgame"
}

func (s *service) Run(mods ...module.Module) error {

	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}

	Ctx = s.Options().Context
	CreateHTTPServer()
	ClearOnline()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	logger.Info("Leaf closing down (signal: %v)", sig)
	module.Destroy()

	return nil
}

//// 取消服务
//func Cancel(name string, id ...string) error {
//	in := &proto.Cancel{Name: name}
//	if len(id) > 0 {
//		in.NodeId = id[0]
//	}
//	return Pub("cancel", in)
//}

//app := &cli.App{
//Flags: redis.Flags,
//}
//
//err := app.Run(os.Args)
//if err != nil {
//logger.Fatal(err)
//}
//
//logger.Info(redis.Opts.RedisUrl)
