package transport

import (
	"context"
	"crypto/tls"
	"time"
)

type Options struct {
	Addrs     []string
	Secure    bool
	TLSConfig *tls.Config
	Timeout   time.Duration
	Context   context.Context
}

type ConnectorOptions struct {
	Stream  bool
	Timeout time.Duration
	Context context.Context
}

type AcceptorOptions struct {
	Context context.Context
}

func Addrs(addrs ...string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}

func Secure(b bool) Option {
	return func(o *Options) {
		o.Secure = b
	}
}

func TLSConfig(t *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = t
	}
}

func WithStream() ConnectorOption {
	return func(o *ConnectorOptions) {
		o.Stream = true
	}
}

func WithTimeout(d time.Duration) ConnectorOption {
	return func(o *ConnectorOptions) {
		o.Timeout = d
	}
}
