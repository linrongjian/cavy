package httpsvr

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strings"

	"bytes"
	"encoding/json"
	"time"

	"github.com/linrongjian/cavy/common/api"
	"github.com/linrongjian/cavy/common/consul"
	"github.com/linrongjian/cavy/common/gerrors"
	"github.com/linrongjian/cavy/common/hook"
	"github.com/linrongjian/cavy/common/mlog"
	"github.com/linrongjian/cavy/common/servercfg"
	"github.com/linrongjian/cavy/proto/pb"
	"google.golang.org/protobuf/proto"
)

type httpHandler struct{}

type PbReply func(r *Request) (proto.Message, int32)
type JsonReply func(r *Request) (int, interface{})
type StrReply func(r *Request) string
type ByteReply func(r *Request) []byte

var (
	intercept      = map[string]bool{"/health": true}
	handlerMap     = make(map[string]func(*Request) []byte)
	tokenWhiteList = map[string]bool{}
	serverID       = ""
	disableSign    = false
	signWhiteList  = map[string]bool{}
	contentPath    = map[string]bool{}
)

func init() {
	flag.BoolVar(&disableSign, "disableSign", false, "是否开启签名")
	serverID = api.NewUUID().String()
}

func Register(paths map[string]PbReply, check bool) {
	for k, fn := range paths {
		if _, ok := handlerMap[k]; ok {
			msg := fmt.Sprintf("path %s is exists!!!", k)
			panic(msg)
		}
		handlerMap[k] = func(r *Request) []byte {
			data, errcode := fn(r)
			reply := &pb.HttpReply{
				Errcode: errcode,
			}
			if errcode == 0 {
				if r.isSetCookie {
					r.w.Header().Set("Set-Session", r.cookie.Encode())
				}
				if data != nil {
					var err error
					reply.Data, err = proto.Marshal(data)
					if err != nil {
						r.Log.Errorf("SendSuccess Marshal data err:%v", err)
					}
				}
			}
			buf, err := proto.Marshal(reply)
			if err != nil {
				r.Log.Errorf("SendSuccess Marshal resp err:%v", err)
				return []byte(err.Error())
			}
			return buf
		}
		if !check {
			tokenWhiteList[k] = true
			signWhiteList[k] = true
		}
	}
}

// RegisterJSON 注册函数
func RegisterJSON(paths map[string]JsonReply) {
	for k, v := range paths {
		if _, ok := handlerMap[k]; ok {
			msg := fmt.Sprintf("path %s is exists!!!", k)
			panic(msg)
		}
		f := v
		handlerMap[k] = func(r *Request) []byte {
			var resp struct {
				Code int         `json:"result"`
				Msg  string      `json:"msg"`
				Data interface{} `json:"data"`
			}
			code, data := f(r)
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
				return []byte(err.Error())
			}
			return buf
		}

		tokenWhiteList[k] = true
		signWhiteList[k] = true
	}
}

// RegisterContent 注册函数
func RegisterContent(paths map[string]ByteReply) {
	for k, v := range paths {
		if _, has := handlerMap[k]; has {
			msg := fmt.Sprintf("path %s is exists!!!", k)
			panic(msg)
		}
		f := v
		handlerMap[k] = func(r *Request) []byte {
			buf := f(r)
			if buf != nil {
				http.ServeContent(r.w, r.r, r.Path, time.Now(), bytes.NewReader(buf))
			} else {
				http.NotFound(r.w, r.r)
			}
			return nil
		}
		contentPath[k] = true
		tokenWhiteList[k] = true
		signWhiteList[k] = true
	}
}

// RegisterString 注册函数
func RegisterString(paths map[string]StrReply) {
	for k, v := range paths {
		if _, ok := handlerMap[k]; ok {
			msg := fmt.Sprintf("path %s is exists!!!", k)
			panic(msg)
		}
		f := v
		handlerMap[k] = func(r *Request) []byte {
			data := f(r)
			return []byte(data)
		}

		tokenWhiteList[k] = true
		signWhiteList[k] = true
	}
}

// wait 等待结束
func wait(svc api.Service) {
	mlog.PanicToError()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	if err := svc.Stop(); err != nil {
		msg := fmt.Sprintf("svc stop err:%v", err)
		panic(msg)
	}
}

// Listen 监听HTTP
func Listen(server *api.Conf, svc api.Service) {
	// 回调钩子后再监听
	hook.CallHook(server)
	//监听服务
	go listen(server.Addr)
	//等待进程
	wait(svc)
}

// Run 启动服务
func Run(cfg *api.Conf, service api.Service) {
	// 回调钩子后再监听
	hook.CallHook(cfg)
	// 启动服务
	run(cfg, service)
	//等待进程
	wait(service)
}

func run(cfg *api.Conf, s api.Service) {
	//监听服务
	// go listenTLS(cfg, cfg.CertFile, cfg.Keyfile)

	go listen(cfg.Addr)
	if err := s.Init(cfg); err != nil {
		msg := fmt.Sprintf("service init err:%v", err)
		panic(msg)
	}
	if err := s.Start(); err != nil {
		msg := fmt.Sprintf("service start err:%v", err)
		panic(msg)
	}

	consul.RegisterToConsul(cfg)

	mlog.NewLogger(nil).Info("服务启动成功")
}

func Start(addr string) {
	go listen(addr)
}

// 验证token
func verifyToken(r *Request, path string) bool {
	if _, ok := tokenWhiteList[path]; !ok {
		token := r.getToken()
		r.Log.Debug("VerifyToken ", token)
		info, err := api.ParseToken(token)
		if err != nil {
			r.Log.Errorf("ParseToken %s err:%v", token, err)
			r.sendFailed(gerrors.TokenErr)
			return false
		}
		r.Log = mlog.NewLogger(mlog.Fields{
			"URL": r.r.URL.String(),
		})
		r.Log.Info("Request Logic Start")
		if info.IsExpire() {
			r.Log.Errorf("Token %s Expire", token)
			r.sendFailed(gerrors.TokenExpire)
			return false
		}

		//小于更新时间时重新生成token并更新
		if info.RemainingTime() < api.UpdateTokenTime {
			newToken := api.GenToken(info.PlayerID)
			r.SetCookie(&api.Cookie{
				Token: newToken,
			})
		}
		r.token = info
	} else {
		r.token = &api.Token{}
	}

	return true
}

// 验证签名
func verifySign(r *Request, path string) bool {
	if disableSign {
		return true
	}
	if signWhiteList[path] {
		return true
	}

	strs := []string{}
	r.r.ParseForm()
	if len(r.r.Form) > 0 {
		for k, v := range r.r.Form {
			strs = append(strs, fmt.Sprintf("%s=%s", k, strings.Join(v, "")))
		}
		sort.Strings(strs)
	}
	heads := []string{r.Method, r.r.URL.Path}
	strs = append(heads, strs...)

	if r.Method == "POST" {
		var b strings.Builder
		// builder类型，比+方式跟buffer的性能高
		for _, v := range r.Body() {
			b.WriteString(fmt.Sprintf("%d", v))
		}
		strs = append(strs, b.String())
		bodyStr := ""
		for _, v := range r.Body() {
			bodyStr += fmt.Sprintf("%d", v)
		}
	}

	sign := r.genSign(strs...)
	if sign != r.r.Header.Get("Sign") {
		r.Log.Errorf("Sign Err Client Sign:%s Server Sign:%s", r.r.Header.Get("Sign"), sign)
		r.sendFailed(gerrors.SignErr)
		return false
	}

	return true
}

// Makesign 设置请求签名
func Makesign(r *http.Request, path string) string {
	strs := []string{}
	r.ParseForm()
	if len(r.Form) > 0 {
		for k, v := range r.Form {
			strs = append(strs, fmt.Sprintf("%s=%s", k, strings.Join(v, "")))
		}
		sort.Strings(strs)
	}
	heads := []string{r.Method, path}
	strs = append(heads, strs...)

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return ""
		}
		r.Body = ioutil.NopCloser(bytes.NewReader(body))

		bodyStr := ""
		for _, v := range body {
			bodyStr += fmt.Sprintf("%d", v)
		}
		strs = append(strs, bodyStr)
	}

	rn := NewRequest(nil, r)
	sign := rn.genSign(strs...)
	return sign
}

// ServeHTTP http回调
func (*httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Sign, Session, Accept, Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Max-Age", fmt.Sprintf("%d", 60*60*24))
	w.Header().Set("Access-Control-Expose-Headers", "Set-Session")

	req := NewRequest(w, r)
	if intercept[req.Path] {
		handleIntercept(req)
		return
	}
	if r.Method == "OPTIONS" {
		req.Log.Debugf("Request Options")
		w.WriteHeader(204)
		return
	}
	now := time.Now()
	req.Log.Info("Request Start")
	defer func() {
		deltatime := time.Now().Sub(now)
		req.Log.Info("Request End ", deltatime)
		if deltatime > 50*time.Millisecond {
			req.Log.Info("-- slowly -- Request End ", deltatime)
		}
	}()

	if handler, ok := handlerMap[req.Path]; ok {
		if req.Path == "/" || (verifyToken(req, req.Path) && verifySign(req, req.Path)) {
			if !contentPath[req.Path] {
				req.send(200, handler(req))
			} else {
				handler(req)
			}
		}

		if req.inblacklist {
			req.sendFailed(gerrors.InBlackList)
			return
		}
	} else {
		paths := strings.Split(req.Path, "/")
		if len(paths) > 0 {
			paths[len(paths)-1] = "*"
		}
		wildcardPath := strings.Join(paths, "/")
		if handler, ok := handlerMap[wildcardPath]; ok {
			if req.Path == "/" || (verifyToken(req, wildcardPath) && verifySign(req, wildcardPath)) {
				if !contentPath[wildcardPath] {
					req.send(200, handler(req))
				} else {
					handler(req)
				}
			}

			if req.inblacklist {
				req.sendFailed(gerrors.InBlackList)
				return
			}
		} else {
			req.SendNotFind()
		}
	}
}

func handleIntercept(req *Request) {
	if req.Path == "/health" {
		req.send(200, []byte("consulCheck"))
	}
}

func GetServerID() string {
	return serverID
}

func GetLocalServer() string {
	return fmt.Sprintf("https://%s:%d", servercfg.Cfg.GetIPAddr(), servercfg.Cfg.GetIPPort())
}
