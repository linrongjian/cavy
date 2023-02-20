package grpcwrap

import (
	"crypto/tls"
	"fastserver/core/network/transport"
	"fastserver/core/protocol/pbgrpc"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type grpcZlGameRpcAcceptor struct {
	listener net.Listener
	secure   bool
	tls      *tls.Config
}

func (t *grpcZlGameRpcAcceptor) Addr() string {
	return t.listener.Addr().String()
}

func (t *grpcZlGameRpcAcceptor) Close() error {
	return t.listener.Close()
}

func (t *grpcZlGameRpcAcceptor) Accept(fn func(transport.Channel)) error {
	var opts []grpc.ServerOption

	// setup tls if specified
	if t.secure || t.tls != nil {
		config := t.tls
		if config == nil {
			var err error
			addr := t.listener.Addr().String()
			config, err = getTLSConfig(addr)
			if err != nil {
				return err
			}
		}

		creds := credentials.NewTLS(config)
		opts = append(opts, grpc.Creds(creds))
	}

	// new service
	srv := grpc.NewServer(opts...)

	// register service
	pbgrpc.RegisterPbGameRPCServer(srv, &grpcStreamDispatcher{addr: t.listener.Addr().String(), fn: fn})

	// start serving
	return srv.Serve(t.listener)
}
