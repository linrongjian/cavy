package app

import (
	"fastserver/core/logger"
	"os"
	"os/signal"

	"github.com/urfave/cli/v2"
)

// type IApp interface {
// 	Init(...Option) error
// 	Options() Options
// }

type Server interface {
	Run() error
	Stop() error
}

type Option func(*Options)

type App struct {
	opts Options
	cli  *cli.App
}

func (a *App) Run() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	logger.Info("Leaf closing down (signal: %v)", sig)
	return nil
}

func (a *App) InitComplete() error {
	a.cli.Run(os.Args)
	return nil
}

func (a *App) Stop() error {
	return nil
}

func (a *App) Options() Options {
	return a.opts
}

func (a *App) Init(opts ...Option) error {
	for _, o := range opts {
		o(&a.opts)
	}
	a.cli = cli.NewApp()
	return nil
}

func (a *App) AddFlags(flags []cli.Flag) {
	a.cli.Flags = append(a.cli.Flags, flags...)
}

func NewApp(opts ...Option) *App {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	return &App{
		opts: options,
	}
}
