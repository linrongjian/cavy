package httpsvr

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/linrongjian/cavy/common/api"
	"github.com/linrongjian/cavy/common/mlog"
	"github.com/linrongjian/cavy/proto/pb"

	"google.golang.org/protobuf/proto"
)

// Request 请求结构
type Request struct {
	w                   http.ResponseWriter
	r                   *http.Request
	vals                url.Values
	isSetCookie         bool
	cookie              *api.Cookie
	Path                string
	Method              string
	Log                 *mlog.Logger
	checkOriginCallBack func(r *Request) bool
	token               *api.Token
	inblacklist         bool
}

const (
	sessionName = "Session="
)

// NewRequest 创建请求对象
func NewRequest(w http.ResponseWriter, r *http.Request) *Request {
	req := &Request{
		w:           w,
		r:           r,
		vals:        r.URL.Query(),
		cookie:      &api.Cookie{},
		Path:        r.URL.Path,
		Method:      r.Method,
		token:       nil,
		isSetCookie: false,
	}
	session := r.Header.Get("Session")
	req.cookie.Decode(session)

	req.Log = mlog.NewLogger(map[string]interface{}{
		"URL": r.URL.String(),
	})
	//req.Log.Debug("Request:", req.Body())
	return req
}

// SetCookie set cookie header
func (r *Request) SetCookie(cookie *api.Cookie) {
	r.isSetCookie = true
	r.cookie = cookie
}

// GetCookie  get cookie header
func (r *Request) GetCookie() *api.Cookie {
	return r.cookie
}

// Body request body
func (r *Request) Body() []byte {
	body, err := ioutil.ReadAll(r.r.Body)
	if err != nil {
		return nil
	}
	r.r.Body = ioutil.NopCloser(bytes.NewReader(body))
	return body
}

// SendFailedProto 发送失败
func (r *Request) sendFailed(errcode int32) error {
	r.w.WriteHeader(200)

	resp := &pb.HttpReply{
		Errcode: errcode,
	}

	buf, err := proto.Marshal(resp)
	if err != nil {
		r.Log.Errorf("SendFailed Marshal err:%v", err)
		return err
	}
	r.w.Write(buf)

	return nil
}

// sendData 发送数据
func (r *Request) sendData(code int, data interface{}) error {
	r.w.WriteHeader(200)
	var resp struct {
		Code int         `json:"result"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}

	if code == 0 {
		resp.Code = 0
		resp.Data = data
	} else {
		resp.Code = code
		resp.Msg = data.(string)
	}

	buf, err := json.Marshal(&resp)
	if err != nil {
		r.Log.Errorf("Marshal err:%v", err)
		return err
	}
	r.w.Write(buf)
	r.Log.Debug("Response:", string(buf))
	return nil
}

// send 发送数据
func (r *Request) send(code int, data []byte) {
	r.w.WriteHeader(200)
	if data != nil {
		r.w.Write(data)
	}
}

// SendSuccess 发送成功
func (r *Request) sendSuccess(data proto.Message) error {
	if r.isSetCookie {
		r.w.Header().Set("Set-Session", r.cookie.Encode())
	}
	r.w.WriteHeader(200)

	dataBuf := []byte{}
	if data != nil {
		buf, err := proto.Marshal(data)
		if err != nil {
			r.Log.Errorf("SendSuccess Marshal data err:%v", err)
			return err
		}
		dataBuf = buf
	}

	result := int32(0)
	resp := &pb.HttpReply{
		Errcode: result,
		Data:    dataBuf,
	}

	buf, err := proto.Marshal(resp)
	if err != nil {
		r.Log.Errorf("SendSuccess Marshal resp err:%v", err)
		return err
	}
	r.w.Write(buf)

	return nil
}

// SendNotFind 发送没有找到
func (r *Request) SendNotFind() {
	r.w.WriteHeader(404)
}

// SendSUCCESS 发送没有找到
func (r *Request) SendSUCCESS() {
	r.w.Write([]byte("SUCCESS"))
}

// SendCode qq回调返回
func (r *Request) SendCode(data []byte) {
	r.w.Write(data)
}

// StringParam 获得字符串参数
func (r *Request) StringParam(key string) (string, bool) {
	val := r.vals.Get(key)
	return val, val != ""
}

// IntParam 获得int参数
func (r *Request) IntParam(key string) (int, bool) {
	val, ok := r.StringParam(key)
	if !ok {
		return 0, false
	}
	vall, err := strconv.Atoi(val)
	if err != nil {
		return 0, false
	}
	return vall, true
}

// Int32Param 获得int32参数
func (r *Request) Int32Param(key string) (int32, bool) {
	val, ok := r.IntParam(key)
	return int32(val), ok
}

// Int64Param 获得int64参数
func (r *Request) Int64Param(key string) (int64, bool) {
	val, ok := r.StringParam(key)
	if !ok {
		return 0, false
	}
	vall, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, false
	}
	return vall, true
}

// AllStringParam 获得所有参数字符串
func (r *Request) AllStringParam() string {
	val := r.vals.Encode()
	return val
}

// VerifyCookie 验证Cookie
func (r *Request) verifyCookie() bool {
	if r.cookie == nil {
		return false
	}

	result := r.VerifyToken()
	if !result {
		return false
	}

	return true
}

// VerifyToken 验证token
func (r *Request) VerifyToken() bool {
	if r.cookie.Token == "" {
		return false
	}

	info, err := api.ParseToken(r.cookie.Token)
	if err != nil {
		return false
	}

	return info.IsExpire()
}

// GetToken 获得用户Token
func (r *Request) getToken() string {
	return r.cookie.Token
}

// Token 获得请求token
func (r *Request) Token() *api.Token {
	return r.token
}

// 检查来源
func (r *Request) checkOrigin(hr *http.Request) bool {
	if r.checkOriginCallBack != nil {
		return r.checkOriginCallBack(r)
	}
	return false
}

// GetRemoteAddr 获得远端地址
func (r *Request) GetRemoteAddr() string {
	return r.r.RemoteAddr
}

// GetRequest 获得http请求
func (r *Request) GetRequest() *http.Request {
	return r.r
}

// GetXForwardedFor 获得远端地址
func (r *Request) GetXForwardedFor() string {
	return r.r.Header.Get("X-Forwarded-For")
}

// GetHTTPRequest 获取http请求结构体
func (r *Request) GetHTTPRequest() *http.Request {
	return r.r
}

// GetVals 获取http请求结构体
func (r *Request) GetVals() url.Values {
	return r.vals
}
