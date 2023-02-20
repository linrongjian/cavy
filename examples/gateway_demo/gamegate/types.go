package gamegate

type AuthData struct {
	Token   string
	UserId  string
	Channel string
}

type GetOnlineCountResult struct {
	Count int32
}

type KickUserParams struct {
	Id string
}

type BroadcastParams struct {
	Data []byte
}

type NotifyParams struct {
	Id   string
	Data []byte
}
