package fastgrpc

import (
	"context"
	"crypto/tls"
	"fastgameserver/fastgameserver/network/gamerpc"
	"fastgameserver/fastgameserver/protocol/pbgrpc"
	"fastgameserver/fastgameserver/util"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcZLGameRpc struct {
	opts gamerpc.Options
}

func (t *grpcZLGameRpc) Dial(addr string, opts ...gamerpc.ConnectorOption) (gamerpc.Connector, error) {
	dopts := gamerpc.ConnectorOptions{
		Timeout: gamerpc.DefaultDialTimeout,
	}

	for _, opt := range opts {
		opt(&dopts)
	}

	options := []grpc.DialOption{
		grpc.WithTimeout(dopts.Timeout),
	}

	var creds credentials.TransportCredentials
	if t.opts.Secure || t.opts.TLSConfig != nil {
		config := t.opts.TLSConfig
		if config == nil {
			config = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
		creds = credentials.NewTLS(config)
	} else {
		creds = insecure.NewCredentials()
	}
	options = append(options, grpc.WithTransportCredentials(creds))

	// dial the server
	conn, err := grpc.Dial(addr, options...)
	if err != nil {
		return nil, err
	}

	// create stream
	stream, err := pbgrpc.NewPbGameRPCClient(conn).Stream(context.Background())
	if err != nil {
		return nil, err
	}

	return &grpcStreamChanConnector{
		conn:   conn,
		stream: stream,
		local:  "localhost",
		remote: addr,
	}, nil
}

func (t *grpcZLGameRpc) Listen(addr string, opts ...gamerpc.AcceptorOption) (gamerpc.Acceptor, error) {
	var options gamerpc.AcceptorOptions
	for _, o := range opts {
		o(&options)
	}

	ln, err := util.Listen(addr, func(addr string) (net.Listener, error) {
		return net.Listen("tcp", addr)
	})
	if err != nil {
		return nil, err
	}

	return &grpcZlGameRpcAcceptor{
		listener: ln,
		tls:      t.opts.TLSConfig,
		secure:   t.opts.Secure,
	}, nil
}

func (t *grpcZLGameRpc) Init(opts ...gamerpc.Option) error {
	for _, o := range opts {
		o(&t.opts)
	}
	return nil
}

func (t *grpcZLGameRpc) Options() gamerpc.Options {
	return t.opts
}

func (t *grpcZLGameRpc) String() string {
	return "grpc"
}
