package stream

import (
	"trainserver/zlgame_rpc"
	"trainserver/zlgame_rpc/grpc/zlgamegrpc"
)

type grpcStreamChanAcceptor struct {
	stream zlgamegrpc.ZLGameRPC_StreamServer
	local  string
	remote string
}

func (g *grpcStreamChanAcceptor) Local() string {
	return g.local
}

func (g *grpcStreamChanAcceptor) Remote() string {
	return g.remote
}

func (g *grpcStreamChanAcceptor) Recv(m *zlgame_rpc.ZLGameMessage) error {
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

func (g *grpcStreamChanAcceptor) Send(m *zlgame_rpc.ZLGameMessage) error {
	if m == nil {
		return nil
	}

	return g.stream.Send(&zlgamegrpc.ZLGameMessage{
		Header: m.Header,
		Body:   m.Body,
	})
}

func (g *grpcStreamChanAcceptor) Close() error {
	return nil
}
