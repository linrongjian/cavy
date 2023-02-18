package fastgrpc

import (
	"crypto/tls"
	"fastgameserver/fastgameserver/network/gamerpc"
	"fastgameserver/fastgameserver/util"
	"net"
)

func NewZLGameRpc(opts ...gamerpc.Option) gamerpc.GameRpc {
	var options gamerpc.Options
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
