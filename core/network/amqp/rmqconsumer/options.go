package rmqconsumer

type Options struct {
	Handle Handle
	Config *Config
}

type Config struct {
	Uri          string
	Exchange     string
	ExchangeType string
	Queue        string
	Tag          string
	RoutingKey   string
}

func WithConfig(config *Config) Option {
	return func(o *Options) {
		o.Config = config
	}
}

func WithHandle(h Handle) Option {
	return func(o *Options) {
		o.Handle = h
	}
}
