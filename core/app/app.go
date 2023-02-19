package app

import (
	"fastgameserver/core/logger"
	"os"
	"os/signal"
)

var ()

type App interface {
	IApp

	Init(...Option) error

	Options() Options
}

type IApp interface {
	Run() error

	Stop() error
}

type Option func(*Options)

type app struct {
	opts Options
}

func (g *app) Run() error {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	logger.Info("Leaf closing down (signal: %v)", sig)

	return nil
}

func (g *app) Stop() error {
	return nil
}

func (g *app) Options() Options {
	return g.opts
}

func (g *app) Init(opts ...Option) error {
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

func NewGame(opts ...Option) IApp {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	return &app{
		opts: options,
	}
}
