package logserver

import (
	"fastserver/core/app"
	"fastserver/core/network/protocols/mqwrap"
	"fastserver/service/logserver/report"
)

var (
	logReport *report.LogReport
)

type LogServer interface {
	app.Server
	Init(...Option) error
	Options() Options
}

type Option func(*Options)

type logServer struct {
	*app.App
	opts Options
}

func NewLogServer(opts ...Option) LogServer {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	app := app.NewApp()
	app.Init()
	return &logServer{
		App:  app,
		opts: options,
	}
}

func (s *logServer) Init(opts ...Option) error {
	for _, o := range opts {
		o(&s.opts)
	}
	s.AddFlags(mqwrap.Flags)
	s.App.InitComplete()
	return nil
}

func (s *logServer) Run() error {
	mqwrap.Startup()
	logReport = report.NewLogReport()
	logReport.Complete()
	s.App.Run()
	return nil
}

func (s *logServer) Stop() error {
	return nil
}

func (s *logServer) Options() Options {
	return s.opts
}
