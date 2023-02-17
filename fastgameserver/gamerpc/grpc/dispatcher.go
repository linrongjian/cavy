package stream

import (
	"trainserver/zlgame_rpc/grpc/zlgamegrpc"

	"google.golang.org/grpc/peer"
)

type grpcStreamDispatcher struct {
	addr string
	fn   func(gamerpc.GameChan)
}

func (m *grpcStreamDispatcher) Stream(s zlgamegrpc.ZLGameRPC_StreamServer) (err error) {

	sock := &grpcStreamChanAcceptor{
		stream: s,
		local:  m.addr,
	}

	p, ok := peer.FromContext(s.Context())
	if ok {
		sock.remote = p.Addr.String()
	}

	defer func() {
		if r := recover(); r != nil {
			// logger.Error(r, string(debug.Stack()))
			sock.Close()
			// err = errors.InternalServerError("go.micro.transport", "panic recovered: %v", r)
		}
	}()

	m.fn(sock)

	return err
}
