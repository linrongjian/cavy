package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// Game System Error Code (1 ~ 99)
const (
	CodeOk        int32 = 0
	CodeUniversal       = 1
	CodeUnknown         = 2
	CodeInvalid         = 3
	CodeExists          = 4

	CodeStarted  = 90
	CodeFinished = 91
	CodeOther    = 99

	// HTTP Status Code (100 ~ 600)
	CodeUnauthorized = 401
	CodeForbidden    = 403
	CodeNotFound     = 404
	CodeTimeout      = 408

	CodeServerError        = 500
	CodeServiceUnavailable = 503

	ErrNewPlayer    // 创建玩家失败
	ErrTokenEmpty   // token is empty
	ErrTokenDecrypt // token decrypt failed
	ErrTokenFormat  // token format is invalid
	ErrTokenExpired // token expired
	UnknownError    // 未知错误
	ErrParam        // 参数错误
	ErrParamNil     // 请求参数为空
	ErrParse        // 解析失败
	ErrDB           // 数据库操作失败
	ErrRedis        // 缓存操作失败
	ErrConnect      // 连接失败
)

func init() {
	AddStatusText(defaultLang, CodeUniversal, "universal")
	AddStatusText(defaultLang, CodeUnknown, "unknown")
	AddStatusText(defaultLang, CodeInvalid, "invalid")
	AddStatusText(defaultLang, CodeExists, "exists")
	AddStatusText(defaultLang, CodeStarted, "started")
	AddStatusText(defaultLang, CodeFinished, "finished")
	AddStatusText(defaultLang, CodeOther, "other")

	AddStatusText(defaultLang, CodeUnauthorized, http.StatusText(CodeUnauthorized))
	AddStatusText(defaultLang, CodeForbidden, http.StatusText(CodeForbidden))
	AddStatusText(defaultLang, CodeNotFound, http.StatusText(CodeNotFound))
	AddStatusText(defaultLang, CodeTimeout, http.StatusText(CodeTimeout))
	AddStatusText(defaultLang, CodeServerError, http.StatusText(CodeServerError))
	AddStatusText(defaultLang, CodeServiceUnavailable, http.StatusText(CodeServiceUnavailable))

	AddStatusText("zh-cn", ErrNewPlayer, "")
	AddStatusText("zh-cn", ErrTokenEmpty, "")
	AddStatusText("zh-cn", ErrTokenDecrypt, "")
	AddStatusText("zh-cn", ErrTokenFormat, "")
	AddStatusText("zh-cn", ErrTokenExpired, "")
	AddStatusText("zh-cn", UnknownError, "")
	AddStatusText("zh-cn", ErrParam, "")
	AddStatusText("zh-cn", ErrParamNil, "")
	AddStatusText("zh-cn", ErrParse, "")
	AddStatusText("zh-cn", ErrDB, "")
	AddStatusText("zh-cn", ErrRedis, "")
	AddStatusText("zh-cn", ErrConnect, "")
}

var isStack bool

func init() {
	isStack, _ = strconv.ParseBool(os.Getenv("GAME_ERROR_STACK"))
}

func IsStack() bool {
	return isStack
}

func message(args ...interface{}) string {
	var msg string
	switch len(args) {
	case 0:
		msg = ""
	case 1:
		msg = ParseStr(args[0])
	default:
		msg = fmt.Sprintf(ParseStr(args[0]), args[1:]...)
	}
	return msg
}

// New returns an error with the supplied proto.
// New also records the stack trace at the point it was called.
func New(code int32, a ...interface{}) *Error {
	e := &Error{
		Code:   code,
		Detail: message(a...),
		Status: StatusText(code),
	}
	if isStack {
		e.stack = callers()
	}
	return e
}

type Error struct {
	Code   int32  `json:"code"`
	Status string `json:"status"`
	Detail string `json:"detail"`
	*stack `json:"-"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.Error())
			if e.stack != nil {
				e.stack.Format(s, verb)
			}
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func Wrap(err error, args ...interface{}) error {
	if err == nil {
		return nil
	}
	e := &wrapError{
		detail: message(args...),
		error:  err,
	}
	if isStack {
		e.stack = callers()
	}
	return e
}

type wrapError struct {
	detail string
	error
	*stack
}

func (w *wrapError) Error() string { return fmt.Sprintf("%s : %s", w.detail, w.error) }

func (w *wrapError) Cause() error { return w.error }

func (w *wrapError) Unwrap() error { return w.error }

func (w *wrapError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.Cause())
			if w.stack != nil {
				w.stack.Format(s, verb)
			}
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}

type causer interface {
	Cause() error
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

func Parse(err error) *Error {
	if err == nil {
		return nil
	}

	switch ee := err.(type) {
	case causer:
		return Parse(ee.Cause())
	case *Error:
		return ee
	// case *errors.Error:
	// 	e := &Error{
	// 		Code:   ee.Code,
	// 		Status: ee.Status,
	// 		Detail: ee.Detail,
	// 	}
	// 	if isStack {
	// 		e.stack = callers()
	// 	}
	// 	return e
	default:
		es := ee.Error()
		e := new(Error)
		if er := json.Unmarshal([]byte(es), e); er != nil {
			e = Unknown(es)
		}
		return e
	}
}

// func MicroError(id string, err error) *errors.Error {
// 	if err == nil {
// 		return nil
// 	}

// 	switch ee := Cause(err).(type) {
// 	case *errors.Error:
// 		return ee
// 	case *Error:
// 		return &errors.Error{
// 			Id:     id,
// 			Code:   ee.Code,
// 			Detail: ee.Detail,
// 			Status: ee.Status,
// 		}
// 	default:
// 		e := new(errors.Error)
// 		if er := json.Unmarshal([]byte(ee.Error()), e); er != nil {
// 			e.Detail = ee.Error()
// 		}
// 		if e.Id == "" {
// 			e.Id = id
// 		}
// 		if e.Code == 0 {
// 			e.Code = CodeUnknown
// 			e.Status = StatusText(CodeUnknown)
// 		}
// 		return e
// 	}
// }

// IsCode 检查错误码
func IsCode(err error, code int32) bool {
	return Parse(Cause(err)).Code == code
}

// Is 比较两个错误
func Is(a, b error) bool {
	if Parse(Cause(a)).Code == Parse(Cause(b)).Code {
		return true
	}
	return false
}

// Universal 通用错误
func Universal(args ...interface{}) *Error {
	return New(CodeUniversal, args...)
}

// Unknown 未知错误
func Unknown(args ...interface{}) *Error {
	return New(CodeUnknown, args...)
}

// Invalid 无效
func Invalid(args ...interface{}) *Error {
	return New(CodeInvalid, args...)
}

// Exists 已存在
func Exists(args ...interface{}) *Error {
	return New(CodeExists, args...)
}

// Started 已开始
func Started(args ...interface{}) *Error {
	return New(CodeStarted, args...)
}

// Finished 已结束
func Finished(args ...interface{}) *Error {
	return New(CodeFinished, args...)
}

// Other 其它错误
func Other(args ...interface{}) *Error {
	return New(CodeOther, args...)
}

// Unauthorized 未授权
func Unauthorized(args ...interface{}) *Error {
	return New(CodeUnauthorized, args...)
}

// Forbidden 拒绝访问
func Forbidden(args ...interface{}) *Error {
	return New(CodeForbidden, args...)
}

// NotFound 不存在
func NotFound(args ...interface{}) *Error {
	return New(CodeNotFound, args...)
}

// Timeout 请求超时
func Timeout(args ...interface{}) *Error {
	return New(CodeTimeout, args...)
}

// Server 服务器错误
func Server(args ...interface{}) *Error {
	return New(CodeServerError, args...)
}

// Unavailable 不可用
func Unavailable(args ...interface{}) *Error {
	return New(CodeServiceUnavailable, args...)
}
