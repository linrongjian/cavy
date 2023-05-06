package logserver

import (
	"eventgo/component/logserver/modules/journal"
	"eventgo/component/logserver/modules/logconsumer"
	"eventgo/core/app"
	"eventgo/core/network/protocols/mqwrap"
	"fmt"
	"log"
)

var (
	logReport *journal.LogReport
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

func NewLogServer(opts ...Option) (LogServer, error) {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}

	s := logServer{
		App:  app.NewApp(),
		opts: options,
	}

	err := s.Init()
	if err != nil {
		return nil, err
	}
	return &s, err
}

func (s *logServer) Init(opts ...Option) error {
	var err error
	for _, o := range opts {
		o(&s.opts)
	}
	s.AddFlags(mqwrap.Flags)

	s.App.Init()

	s.opts.logConsumer, err = logconsumer.NewLogConsumer()
	if err != nil {
		return fmt.Errorf("log consumer: %s", err)
	}

	mqwrap.Startup()

	// logReport = journal.NewLogReport()
	// logReport.Init()

	// s.App.InitComplete()
	return nil
}

// func (s *logServer) Run() error {

// 	s.App.Run()
// 	return nil
// }

func (s *logServer) Stop() error {
	log.Printf("log server is stopping")
	return nil
}

func (s *logServer) Options() Options {
	return s.opts
}
