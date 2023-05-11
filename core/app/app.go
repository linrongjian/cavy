package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var (
	Opt = Options{}
)

type Server interface {
	Stop() error
}

type Option func(*Options)

type App struct {
	opts Options
}

func NewApp(opts ...Option) *App {
	options := Options{}
	app := &App{
		opts: options,
	}

	for _, o := range opts {
		o(&options)
	}
	return app
}

func (a *App) Init(opts ...Option) error {
	for _, o := range opts {
		o(&a.opts)
	}
	a.opts.Cli.Flags = append(a.opts.Cli.Flags, DefaultFlags...)
	a.opts.Cli.Run(os.Args)
	return nil
}

func Run(app Server) error {
	waitSignal(app)

	fmt.Println("cavy start...")
	sum := 0
	for {
		sum++
		time.Sleep(time.Second)
	}
}

func (a *App) Stop() error {
	log.Println("app an elegant exit")
	os.Exit(0)
	return nil
}

func (a *App) AddFlags(flags []cli.Flag) {
	a.opts.Cli.Flags = append(a.opts.Cli.Flags, flags...)
}

func (a *App) Options() Options {
	return a.opts
}
