package stream

import (
	"trainserver/zlgame_rpc"
	"trainserver/zlgame_rpc/grpc/zlgamegrpc"

	"google.golang.org/grpc"
)

type grpcStreamChanConnector struct {
	conn   *grpc.ClientConn
	stream zlgamegrpc.ZLGameRPC_StreamClient

	local  string
	remote string
}

func (g *grpcStreamChanConnector) Local() string {
	return g.local
}

func (g *grpcStreamChanConnector) Remote() string {
	return g.remote
}

func (g *grpcStreamChanConnector) Recv(m *zlgame_rpc.ZLGameMessage) error {
	if m == nil {
		return nil
	}

	msg, err := g.stream.Recv()
	if err != nil {
		return err
	}

	m.Header = msg.Header
	m.Body = msg.Body
	return nil
}

func (g *grpcStreamChanConnector) Send(m *zlgame_rpc.ZLGameMessage) error {
	if m == nil {
		return nil
	}

	return g.stream.Send(&zlgamegrpc.ZLGameMessage{
		Header: m.Header,
		Body:   m.Body,
	})
}

func (g *grpcStreamChanConnector) Close() error {
	return g.conn.Close()
}
