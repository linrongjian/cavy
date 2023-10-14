/*
 * @Author: zheng
 * @Date: 2019-08-14 14:25:24
 * @Description:
 */

package gerrors

// Error 错误
type Error int32

const (
	/*
		通用模块错误 0~499
	*/

	// Success 成功
	Success = Error(iota)
	// Failed 失败
	Failed = Error(1)
	// ParamNil 请求参数为空
	ParamNil = Error(2)
	// SignErr 签名错误
	SignErr = Error(3)
	// CookieErr Cookie错误
	CookieErr = Error(4)
	// LoginFailed 登录失败
	LoginFailed = Error(5)
	// TokenErr Token错误
	TokenErr = Error(6)
	// TokenExpire Token过期
	TokenExpire = Error(7)
	// ParseErr 解析失败
	ParseErr = Error(8)
	// RedisErr 缓存操作失败
	RedisErr = Error(9)
	// DbErr 缓存操作失败
	DbErr = Error(10)
	// RequestMethodErr 请求模式错误
	RequestMethodErr = Error(11)
	// RequestBodyErr 请求Body错误
	RequestBodyErr = Error(12)
	// RequestParamErr 请求参数错误
	RequestParamErr = Error(13)
	// ParseRequestBodyErr 解析请求包错误
	ParseRequestBodyErr = Error(14)
	// VerifySignFailed 验证签名失败
	VerifySignFailed = Error(15)
	// ConfigErr 配置错误
	ConfigErr = Error(16)
	// InBlackList 请求太频繁，已被加入黑名单，请稍后再登
	InBlackList = Error(17)
	// SaveDBFailed 保存DB失败
	SaveDBFailed = Error(18)
	// ErrClose 功能暂未开放
	ErrClose = Error(19)
	// 未成年认证
	Err18Validation = Error(20)
	// ErrWxAccount 非实名用户账号不可发放
	/*
		登录模块错误 500~999
	*/

	// NotFindSearchPlayer 没找到对应玩家
	NotFindSearchPlayer = Error(500)
	// NotExistsPlayer 用户不存在
	NotExistsPlayer = Error(501)
	// WBVerifyErr 玩吧验证失败
	WBVerifyErr = Error(503)
	// NotSupportTestLogin 不支持账号密码登录
	NotSupportTestLogin = Error(504)
	//OpenIDIsNull openID为NULL
	OpenIDIsNull = Error(505)
	// WXVerifyFailed 微信验证失败
	WXVerifyFailed = Error(506)
	// UnionIDIsNull unionID为NULL
	UnionIDIsNull = Error(507)
	// ErrIdentifyPassword 密码错误
	ErrIdentifyPassword = Error(508)
	// ErrPasswordNotEmpty 密码不为空
	ErrPasswordNotEmpty = Error(509)
	// ErrWechatHasExist 该微信已绑定玩家
	ErrWechatHasExist = Error(510)
	//ErrWechatExpire 微信授权已过期
	ErrWechatExpire = Error(511)
	// ErrPlatformHasBind 该平台用户已绑定
	ErrPlatformHasBind = Error(512)
	// ErrPlatformHasBind 绑定平台失败
	ErrPlatformBindErr = Error(513)
	// ErrWxAccount 非实名用户账号不可发放
	ErrWxAccountErr = Error(514)
	// ErrQQMemberCDKErr QQ会员CDK错误
	ErrQQMemberCDKErr = Error(515)
	/*
		支付模块错误 1000~1499
	*/

	// MidasGetBalanceErr 米大师拉取钻石失败
	MidasGetBalanceErr = Error(1000)
	// MidasPayErr 米大师消耗钻石失败
	MidasPayErr = Error(1001)
	// CreateOrderErr 下单失败
	CreateOrderErr = Error(1002)
	// WXPayErr 微信支付失败
	WXPayErr = Error(1003)
	// YiJieVerifyFailed 易接验证失败
	YiJieVerifyFailed = Error(1004)
	// ErrWxAccessToken 微信accessToken错误
	ErrWxAccessToken = Error(1005)
	// ErrBilibiliPay bilibili支付错误
	ErrBilibiliPay = Error(1006)

	/*
		绑定手机模块错误 1500~1600
	*/

	// ErrPhoneSendCode 发送验证码失败
	ErrPhoneSendCode = Error(1500)
	// ErrBindGetCodeFast 获取验证码速度过快，请稍后再试
	ErrBindGetCodeFast = Error(1501)
	// ErrPhoneNo 手机号码不合法
	ErrPhoneNo = Error(1502)
	// ErrPhoneModifyOldCodeErr  您输入的原手机验证码错误
	ErrPhoneModifyOldCodeErr = Error(1503)
	// ErrPhoneModifyNewCodeErr  您输入的新手机验证码错误
	ErrPhoneModifyNewCodeErr = Error(1504)
	// ErrPhoneAuth 手机注册失败
	ErrPhoneAuth = Error(1505)
	// ErrPhoneAuthExist 手机已经注册过
	ErrPhoneAuthExist = Error(1506)
	// ErrPhoneCodeErr 验证不通过
	ErrPhoneCodeErr = Error(1507)
	// ErrPhoneBindFail 手机号绑定失败
	ErrPhoneBindFail = Error(1508)
	// ErrPhoneHasBind 手机号已绑定
	ErrPhoneHasBind = Error(1509)
	// ErrPassword 密码不合法
	ErrPassword = Error(1510)
	// ErrUserNameExist 用户名已经注册过
	ErrUserNameExist = Error(1511)
	// ErrUserName 用户名不合法
	ErrUserName = Error(1512)
)

var errMsg = map[Error]string{
	/*
		通用模块错误 0~499
	*/
	Success:             "成功",
	Failed:              "失败",
	ParamNil:            "请求参数为空",
	SignErr:             "请求签名错误",
	CookieErr:           "请求cookie错误",
	LoginFailed:         "登录失败",
	TokenErr:            "Token错误",
	TokenExpire:         "Token过期",
	ParseErr:            "数据解析失败",
	RedisErr:            "缓存操作失败",
	DbErr:               "数据库操作失败",
	NotFindSearchPlayer: "没找到对应玩家",
	RequestMethodErr:    "请求模式错误",
	RequestBodyErr:      "请求Body错误",
	RequestParamErr:     "请求参数错误",
	ParseRequestBodyErr: "解析请求包错误",
	VerifySignFailed:    "验证签名失败",
	SaveDBFailed:        "保存DB失败",
	ErrWxAccessToken:    "授权登录失败",
	ErrClose:            "功能暂未开放",
	Err18Validation:     "未成年认证",

	/*
		登录模块错误 500~999
	*/
	NotExistsPlayer:     "用户不存在",
	ConfigErr:           "配置错误",
	InBlackList:         "请求太频繁，已被加入黑名单，请稍后再登",
	WBVerifyErr:         "玩吧验证失败",
	NotSupportTestLogin: "不支持账号密码登录",
	OpenIDIsNull:        "OpenID为空",
	WXVerifyFailed:      "微信验证失败",
	UnionIDIsNull:       "unionID为NULL",
	ErrIdentifyPassword: "密码错误",
	ErrPasswordNotEmpty: "密码不为空",
	ErrPassword:         "密码格式错误",
	ErrUserName:         "用户名格式错误",
	ErrWechatHasExist:   "微信号已绑定玩家",
	ErrWechatExpire:     "微信授权已过期",
	ErrPlatformHasBind:  "该平台已绑定",
	ErrPlatformBindErr:  "绑定平台失败",
	ErrWxAccountErr:     "非实名用户账号不可发放",

	/*
		支付模块错误 1000~1499
	*/
	CreateOrderErr:     "下单失败",
	MidasPayErr:        "米大师消耗钻石失败",
	MidasGetBalanceErr: "米大师拉取钻石失败",
	YiJieVerifyFailed:  "易接验证失败",

	/*
		绑定手机模块错误 1500~1600
	*/
	ErrPhoneSendCode:         "发送短信验证码失败",
	ErrBindGetCodeFast:       "获取验证码速度过快，请稍后再试",
	ErrPhoneNo:               "手机号码不合法",
	ErrPhoneModifyOldCodeErr: "您输入的原手机验证码错误",
	ErrPhoneModifyNewCodeErr: "您输入的新手机验证码错误",
	ErrPhoneAuth:             "手机注册失败",
	ErrPhoneAuthExist:        "手机号已绑定，请联系客服!",
	ErrPhoneCodeErr:          "验证码错误",
	ErrPhoneBindFail:         "手机号绑定失败",
	ErrPhoneHasBind:          "手机号已绑定",
	ErrUserNameExist:         "该账号已存在，请换一个试试",
}

// String 获得错误码描述信息
func (e Error) String() string {
	v, ok := errMsg[e]
	if !ok {
		return "未定义错误描述"
	}

	return v
}
