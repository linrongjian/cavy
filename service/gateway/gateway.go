package gateway

import (
	"servergo/core/app"
	"servergo/core/logger"
	"servergo/core/network/protocols/httpwrap"
	"servergo/core/network/protocols/mqwrap"
	"os"
	"os/signal"
)

type GateServer interface {
	app.Server
	Init(...Option) error
	Options() Options
}

type Option func(*Options)

type gateServer struct {
	*app.App
	opts Options
}

func (s *gateServer) Run() error {

	s.App.Run()

	// Ctx = s.Options().Context
	httpwrap.CreateHTTPServer()
	// ClearOnline()

	mqwrap.Startup()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	logger.Info("Leaf closing down (signal: %v)", sig)

	return nil
}

func (s *gateServer) Stop() error {
	return nil
}

func (s *gateServer) Options() Options {
	return s.opts
}

func (s *gateServer) Init(opts ...Option) error {
	for _, o := range opts {
		o(&s.opts)
	}

	httpwrap.RegisterGetHandleNoUserID("/", onConnectHandle) //获取入口信息

	s.AddFlags(mqwrap.Flags)

	// cmd.AddFlags(defaultFlags)
	// cmd.AddFlags(redis.Flags)
	// cmd.AddFlags(mq.Flags)
	// cmd.AddFlags(mysql.Flags)
	// if err := s.opts.Cmd.Init(); err != nil {
	// logger.Fatal(err)
	// }
	//mq.Startup()
	// mysql.Startup()

	// err := redis.Connect()
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	// s.opts.Rds = redis.S()
	return nil
}

func (s *gateServer) getOnlineCount() int {
	return 1
}

func (s *gateServer) kickUser(userId string) {
	// return this.wsGateway.kick(userId);
}

func (s *gateServer) broadcast(data interface{}) {
	// return this.wsGateway.broadcast(data);
}

func (s *gateServer) notify(userId string, data interface{}) {
	// return this.wsGateway.notify(userId, data)
}

func NewGateServer(opts ...Option) GateServer {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	app := app.NewApp()
	app.Init()
	return &gateServer{
		App:  app,
		opts: options,
	}
}
