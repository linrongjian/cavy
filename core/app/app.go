package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/urfave/cli/v2"
)

type Server interface {
	Run() error
	Stop() error
}

type Option func(*Options)

type App struct {
	opts Options
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

func (a *App) Init(opts ...Option) error {
	for _, o := range opts {
		o(&a.opts)
	}
	a.opts.Cli.Run(os.Args)
	return nil
}

func (a *App) Run() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT,
		syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("Program Exit...", s)
				a.Stop()
			case syscall.SIGUSR1:
				fmt.Println("usr1 signal", s)
			case syscall.SIGUSR2:
				fmt.Println("usr2 signal", s)
			default:
				fmt.Println("other signal", s)
			}
		}
	}()

	fmt.Println("Program Start...")
	sum := 0
	for {
		sum++
		time.Sleep(time.Second)
	}
}

func (a *App) Stop() error {
	os.Exit(0)
	return fmt.Errorf("app stop")
}

func (a *App) AddFlags(flags []cli.Flag) {
	a.opts.Cli.Flags = append(a.opts.Cli.Flags, flags...)
}

func (a *App) Options() Options {
	return a.opts
}
