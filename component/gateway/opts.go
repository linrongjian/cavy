package gateway

import (
	"context"

	"cavy/core/store/redis"
)

type Opts struct {
	// Server baseserver.Server
	Rds *redis.Store
	//mysql
	//mq
	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error

	Context context.Context

	Signal bool
}

func newOptions(opts ...Option) Opts {
	opt := Opts{
		Context: context.Background(),
		Signal:  true,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func Context1(ctx context.Context) Option {
	return func(o *Opts) {
		o.Context = ctx
	}
}

func HandleSignal(b bool) Option {
	return func(o *Opts) {
		o.Signal = b
	}
}

func BeforeStart(fn func() error) Option {
	return func(o *Opts) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}

func BeforeStop(fn func() error) Option {
	return func(o *Opts) {
		o.BeforeStop = append(o.BeforeStop, fn)
	}
}

func AfterStart(fn func() error) Option {
	return func(o *Opts) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}

func AfterStop(fn func() error) Option {
	return func(o *Opts) {
		o.AfterStop = append(o.AfterStop, fn)
	}
}
