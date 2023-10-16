package httpsvr

type HttpReply struct {
	Err  int    `json:"err"`
	Msg  string `json:"msg"`
	Data []byte `json:"data"`
}
