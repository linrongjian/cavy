package grpcwrap

import (
	"cavy/core/network/transport"
	"cavy/core/util"
	"crypto/tls"
	"net"
)

func NewGameRpc(opts ...transport.Option) transport.Transport {
	var options transport.Options
	for _, o := range opts {
		o(&options)
	}
	return &grpcZLGameRpc{opts: options}
}

func init() {
	// cmd.DefaultZLGameRpc["grpc"] = NewGrpcStream
}

func getTLSConfig(addr string) (*tls.Config, error) {
	hosts := []string{addr}

	// check if its a valid host:port
	if host, _, err := net.SplitHostPort(addr); err == nil {
		if len(host) == 0 {
			hosts = util.IPs()
		} else {
			hosts = []string{host}
		}
	}

	// generate a certificate
	cert, err := util.Certificate(hosts...)
	if err != nil {
		return nil, err
	}

	return &tls.Config{Certificates: []tls.Certificate{cert}}, nil
}
