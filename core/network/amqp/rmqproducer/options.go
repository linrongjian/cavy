package rmqproducer

type Options struct {
	Config *Config
}

type Config struct {
	Uri          string
	Exchange     string
	ExchangeType string
	RoutingKey   string
	Reliable     bool
}

func WithConfig(config *Config) Option {
	return func(o *Options) {
		o.Config = config
	}
}
