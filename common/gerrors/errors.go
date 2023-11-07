/*
 * @Author: zheng
 * @Date: 2019-08-14 14:25:24
 * @Description:
 */

package gerrors

// Error 错误
// type Error int32

const (
	/*
		通用模块错误 0~499
	*/

	// Success 成功
	Success = iota
	// Failed 失败
	Failed = 1
	// ParamNil 请求参数为空
	ParamNil = 2
	// SignErr 签名错误
	SignErr = 3
	// CookieErr Cookie错误
	CookieErr = 4
	// LoginFailed 登录失败
	LoginFailed = 5
	// TokenErr Token错误
	TokenErr = 6
	// TokenExpire Token过期
	TokenExpire = 7
	// ParseErr 解析失败
	ParseErr = 8
	// RedisErr 缓存操作失败
	RedisErr = 9
	// DbErr 缓存操作失败
	DbErr = 10
	// RequestMethodErr 请求模式错误
	RequestMethodErr = 11
	// RequestBodyErr 请求Body错误
	RequestBodyErr = 12
	// RequestParamErr 请求参数错误
	RequestParamErr = 13
	// ParseRequestBodyErr 解析请求包错误
	ParseRequestBodyErr = 14
	// VerifySignFailed 验证签名失败
	VerifySignFailed = 15
	// ConfigErr 配置错误
	ConfigErr = 16
	// InBlackList 请求太频繁，已被加入黑名单，请稍后再登
	InBlackList = 17
	// SaveDBFailed 保存DB失败
	SaveDBFailed = 18
	// ErrClose 功能暂未开放
	ErrClose = 19
	// 未成年认证
	Err18Validation = 20
	// ErrWxAccount 非实名用户账号不可发放
	/*
		登录模块错误 500~999
	*/

	// NotFindSearchPlayer 没找到对应玩家
	NotFindSearchPlayer = 500
	// NotExistsPlayer 用户不存在
	NotExistsPlayer = 501
	// WBVerifyErr 玩吧验证失败
	WBVerifyErr = 503
	// NotSupportTestLogin 不支持账号密码登录
	NotSupportTestLogin = 504
	//OpenIDIsNull openID为NULL
	OpenIDIsNull = 505
	// WXVerifyFailed 微信验证失败
	WXVerifyFailed = 506
	// UnionIDIsNull unionID为NULL
	UnionIDIsNull = 507
	// ErrIdentifyPassword 密码错误
	ErrIdentifyPassword = 508
	// ErrPasswordNotEmpty 密码不为空
	ErrPasswordNotEmpty = 509
	// ErrWechatHasExist 该微信已绑定玩家
	ErrWechatHasExist = 510
	//ErrWechatExpire 微信授权已过期
	ErrWechatExpire = 511
	// ErrPlatformHasBind 该平台用户已绑定
	ErrPlatformHasBind = 512
	// ErrPlatformHasBind 绑定平台失败
	ErrPlatformBindErr = 513
	// ErrWxAccount 非实名用户账号不可发放
	ErrWxAccountErr = 514
	// ErrQQMemberCDKErr QQ会员CDK错误
	ErrQQMemberCDKErr = 515
	/*
		支付模块错误 1000~1499
	*/

	// MidasGetBalanceErr 米大师拉取钻石失败
	MidasGetBalanceErr = 1000
	// MidasPayErr 米大师消耗钻石失败
	MidasPayErr = 1001
	// CreateOrderErr 下单失败
	CreateOrderErr = 1002
	// WXPayErr 微信支付失败
	WXPayErr = 1003
	// YiJieVerifyFailed 易接验证失败
	YiJieVerifyFailed = 1004
	// ErrWxAccessToken 微信accessToken错误
	ErrWxAccessToken = 1005
	// ErrBilibiliPay bilibili支付错误
	ErrBilibiliPay = 1006

	/*
		绑定手机模块错误 1500~1600
	*/

	// ErrPhoneSendCode 发送验证码失败
	ErrPhoneSendCode = 1500
	// ErrBindGetCodeFast 获取验证码速度过快，请稍后再试
	ErrBindGetCodeFast = 1501
	// ErrPhoneNo 手机号码不合法
	ErrPhoneNo = 1502
	// ErrPhoneModifyOldCodeErr  您输入的原手机验证码错误
	ErrPhoneModifyOldCodeErr = 1503
	// ErrPhoneModifyNewCodeErr  您输入的新手机验证码错误
	ErrPhoneModifyNewCodeErr = 1504
	// ErrPhoneAuth 手机注册失败
	ErrPhoneAuth = 1505
	// ErrPhoneAuthExist 手机已经注册过
	ErrPhoneAuthExist = 1506
	// ErrPhoneCodeErr 验证不通过
	ErrPhoneCodeErr = 1507
	// ErrPhoneBindFail 手机号绑定失败
	ErrPhoneBindFail = 1508
	// ErrPhoneHasBind 手机号已绑定
	ErrPhoneHasBind = 1509
	// ErrPassword 密码不合法
	ErrPassword = 1510
	// ErrUserNameExist 用户名已经注册过
	ErrUserNameExist = 1511
	// ErrUserName 用户名不合法
	ErrUserName = 1512
)

// var errMsg = map[Error]string{
// 	/*
// 		通用模块错误 0~499
// 	*/
// 	Success:             "成功",
// 	Failed:              "失败",
// 	ParamNil:            "请求参数为空",
// 	SignErr:             "请求签名错误",
// 	CookieErr:           "请求cookie错误",
// 	LoginFailed:         "登录失败",
// 	TokenErr:            "Token错误",
// 	TokenExpire:         "Token过期",
// 	ParseErr:            "数据解析失败",
// 	RedisErr:            "缓存操作失败",
// 	DbErr:               "数据库操作失败",
// 	NotFindSearchPlayer: "没找到对应玩家",
// 	RequestMethodErr:    "请求模式错误",
// 	RequestBodyErr:      "请求Body错误",
// 	RequestParamErr:     "请求参数错误",
// 	ParseRequestBodyErr: "解析请求包错误",
// 	VerifySignFailed:    "验证签名失败",
// 	SaveDBFailed:        "保存DB失败",
// 	ErrWxAccessToken:    "授权登录失败",
// 	ErrClose:            "功能暂未开放",
// 	Err18Validation:     "未成年认证",

// 	/*
// 		登录模块错误 500~999
// 	*/
// 	NotExistsPlayer:     "用户不存在",
// 	ConfigErr:           "配置错误",
// 	InBlackList:         "请求太频繁，已被加入黑名单，请稍后再登",
// 	WBVerifyErr:         "玩吧验证失败",
// 	NotSupportTestLogin: "不支持账号密码登录",
// 	OpenIDIsNull:        "OpenID为空",
// 	WXVerifyFailed:      "微信验证失败",
// 	UnionIDIsNull:       "unionID为NULL",
// 	ErrIdentifyPassword: "密码错误",
// 	ErrPasswordNotEmpty: "密码不为空",
// 	ErrPassword:         "密码格式错误",
// 	ErrUserName:         "用户名格式错误",
// 	ErrWechatHasExist:   "微信号已绑定玩家",
// 	ErrWechatExpire:     "微信授权已过期",
// 	ErrPlatformHasBind:  "该平台已绑定",
// 	ErrPlatformBindErr:  "绑定平台失败",
// 	ErrWxAccountErr:     "非实名用户账号不可发放",

// 	/*
// 		支付模块错误 1000~1499
// 	*/
// 	CreateOrderErr:     "下单失败",
// 	MidasPayErr:        "米大师消耗钻石失败",
// 	MidasGetBalanceErr: "米大师拉取钻石失败",
// 	YiJieVerifyFailed:  "易接验证失败",

// 	/*
// 		绑定手机模块错误 1500~1600
// 	*/
// 	ErrPhoneSendCode:         "发送短信验证码失败",
// 	ErrBindGetCodeFast:       "获取验证码速度过快，请稍后再试",
// 	ErrPhoneNo:               "手机号码不合法",
// 	ErrPhoneModifyOldCodeErr: "您输入的原手机验证码错误",
// 	ErrPhoneModifyNewCodeErr: "您输入的新手机验证码错误",
// 	ErrPhoneAuth:             "手机注册失败",
// 	ErrPhoneAuthExist:        "手机号已绑定，请联系客服!",
// 	ErrPhoneCodeErr:          "验证码错误",
// 	ErrPhoneBindFail:         "手机号绑定失败",
// 	ErrPhoneHasBind:          "手机号已绑定",
// 	ErrUserNameExist:         "该账号已存在，请换一个试试",
// }

// // String 获得错误码描述信息
// func (e Error String( string {
// 	v, ok := errMsg[e]
// 	if !ok {
// 		return "未定义错误描述"
// 	}

// 	return v
// }
