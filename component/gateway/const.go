package gateway

const (
	SUCCESS         = iota // 成功
	ErrNewPlayer           // 创建玩家失败
	ErrTokenEmpty          // token is empty
	ErrTokenDecrypt        // token decrypt failed
	ErrTokenFormat         // token format is invalid
	ErrTokenExpired        // token expired
	UnknownError           // 未知错误
	ErrParam               // 参数错误
	ErrParamNil            // 请求参数为空
	ErrParse               // 解析失败
	ErrDB                  // 数据库操作失败
	ErrRedis               // 缓存操作失败
	ErrConnect             // 连接失败
)
