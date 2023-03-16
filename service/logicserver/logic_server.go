package logicserver

import (
	"fastserver/core/app"
	"fastserver/core/logger"
	"os"
	"os/signal"
)

type LogicServer interface {
	app.Server
	Init(...Option) error
	Options() Options
}

type Option func(*Options)

type logicServer struct {
	*app.App
	opts Options
}

func (s *logicServer) Run() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	logger.Info("Leaf closing down (signal: %v)", sig)
	return nil
}

func (s *logicServer) Stop() error {
	return nil
}

func (s *logicServer) Options() Options {
	return s.opts
}

func (s *logicServer) Init(opts ...Option) error {
	for _, o := range opts {
		o(&s.opts)
	}

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

func NewLogicServer(opts ...Option) LogicServer {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}

	app := app.NewApp()
	app.Init()
	return &logicServer{
		App:  app,
		opts: options,
	}
}
