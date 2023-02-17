package stream

import (
	"context"
	"crypto/tls"
	"net"
	mnet "trainserver/util/net"
	zlgame_rpc2 "trainserver/zlgame_rpc"
	"trainserver/zlgame_rpc/grpc/zlgamegrpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcZLGameRpc struct {
	opts zlgame_rpc2.Options
}

func (t *grpcZLGameRpc) Dial(addr string, opts ...zlgame_rpc2.ConnectorOption) (zlgame_rpc2.Connector, error) {
	dopts := zlgame_rpc2.ConnectorOptions{
		Timeout: zlgame_rpc2.DefaultDialTimeout,
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
	stream, err := zlgamegrpc.NewZLGameRPCClient(conn).Stream(context.Background())
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

func (t *grpcZLGameRpc) Listen(addr string, opts ...zlgame_rpc2.AcceptorOption) (zlgame_rpc2.Acceptor, error) {
	var options zlgame_rpc2.AcceptorOptions
	for _, o := range opts {
		o(&options)
	}

	ln, err := mnet.Listen(addr, func(addr string) (net.Listener, error) {
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

func (t *grpcZLGameRpc) Init(opts ...zlgame_rpc2.Option) error {
	for _, o := range opts {
		o(&t.opts)
	}
	return nil
}

func (t *grpcZLGameRpc) Options() zlgame_rpc2.Options {
	return t.opts
}

func (t *grpcZLGameRpc) String() string {
	return "grpc"
}
