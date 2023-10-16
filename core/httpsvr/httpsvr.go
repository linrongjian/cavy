package httpsvr

type HttpResponse struct {
	Errcode int32  // 结果码 - 0:成功 other:失败
	Msg     string // 错误信息 - 失败时读取
	Date    []byte // 消息包体 - 成功时读取
}
