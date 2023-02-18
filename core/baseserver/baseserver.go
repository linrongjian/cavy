package baseserver

import (
	"fastgameserver/core/logger"
	"os"
	"os/signal"
)

var (
	Instance *CServer = nil
)

type CServer interface {
	IServer

	Init(...Option) error

	Options() Options
}

type IServer interface {
	Run() error

	Stop() error
}

type Option func(*Options)

type game struct {
	opts Options
}

func (g *game) Run() error {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	logger.Info("Leaf closing down (signal: %v)", sig)

	return nil
}

func (g *game) Stop() error {
	return nil
}

func (g *game) Options() Options {
	return g.opts
}

func (g *game) Init(opts ...Option) error {
	for _, o := range opts {
		o(&g.opts)
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

func NewGame(opts ...Option) IServer {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	return &game{
		opts: options,
	}
}
