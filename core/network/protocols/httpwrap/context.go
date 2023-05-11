package httpwrap

import (
	"cavy/core/logger"
	"cavy/core/network/protocols/mqwrap"
	"cavy/core/protocol/pb"
	"cavy/core/store/mysql"
	"cavy/core/util"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/go-redis/redis/v8"
	"github.com/go-xorm/xorm"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/protobuf/proto"
)

type Context struct {
	Rds       redis.Conn
	Db        *xorm.Engine
	MqChannel *mqwrap.MqChannel
	UserID    string
	Query     url.Values
	Params    httprouter.Params
	Body      []byte
	r         *http.Request
	W         http.ResponseWriter
	Ctx       context.Context
}

type Handle func(*Context)

func newReqContext(r *http.Request, requiredUserID bool) (*Context, error) {
	userID := ""

	query := r.URL.Query()
	tk := r.Header.Get("Session")
	if requiredUserID {
		id, err := util.ParseTK(tk)
		if err != nil {
			return nil, err
		}
		userID = id
	}

	ctx := &Context{}
	ctx.UserID = userID
	ctx.Query = query
	return ctx, nil
}

// WriteRsp send protobuf messsage to peer
func (ctx *Context) WriteRsp(m *pb.HTTPResponse) {
	if m.GetResult() != 0 {
		if m.GetMsg() == "" {
			m.Msg = proto.String(util.StatusText(m.GetResult()))
		}
	}
	buf, err := proto.Marshal(m)
	if err != nil {
		logger.Errorf("WriteRsp panic:", err)
	}
	ctx.W.Write(buf)
}

// GetHTTPRequest 获取请求信息
func (ctx *Context) GetHTTPRequest() *http.Request {
	return ctx.r
}

// replyClientWithTokenError reply to client for decoding token failed
func replyClientWithTokenError(w http.ResponseWriter, err error) {
	// result := int32(ErrTokenFormat)
	// switch errCode {
	// case ErrTokenEmpty:
	// 	result = ErrTokenEmpty
	// case ErrTokenDecrypt:
	// 	result = int32(ErrTokenDecrypt)
	// case ErrTokenFormat:
	// 	result = int32(ErrTokenFormat)
	// case ErrTokenExpired:
	// 	result = int32(ErrTokenExpired)
	// }

	httpRsp := pb.HTTPResponse{}
	httpRsp.Result = proto.Int32(-1)
	httpRsp.Msg = proto.String(err.Error())

	// 退出函数时发送回复给客户端
	bytes, err := proto.Marshal(&httpRsp)
	if err != nil {
		logger.Errorf("WriteRsp panic:", err)
	}

	w.Write(bytes)
}

// wrapGetHandleInternal 包装 get handle
func wrapGetHandleInternal(h Handle, requiredUserID bool) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		ctx, err := newReqContext(r, requiredUserID)
		if err != nil {
			replyClientWithTokenError(w, err)
			return
		}

		ctx.Params = params
		ctx.r = r
		ctx.W = w
		ctx.Ctx = Ctx
		ctx.Db = mysql.DbConnect
		h(ctx)
	}
}

// wrapPostHandleInternal 包装 post handle
func wrapPostHandleInternal(h Handle, requiredUserID bool) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		ctx, err := newReqContext(r, requiredUserID)
		if err != nil {
			replyClientWithTokenError(w, err)
			return
		}

		ctx.Params = params

		// read all body bytes
		// Read body
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		ctx.Body = b
		ctx.r = r
		ctx.W = w
		h(ctx)
	}
}

// RegisterGetHandle 注册http get handle
func RegisterGetHandle(subPath string, handle Handle) {
	logger.Info("RegisterGetHandle:", subPath)
	if subPath[0] != '/' {
		logger.Errorf("subPath must begin with '/', :", subPath)
	}

	path := rootPath + subPath
	h, _, _ := rootRouter.Lookup("GET", path)
	if h != nil {
		logger.Errorf("subPath with 'GET' has been register, subPath:", subPath)
	}

	rootRouter.GET(path, wrapGetHandleInternal(handle, true))
}

// RegisterPostHandle 注册http post handle
func RegisterPostHandle(subPath string, handle Handle) {
	logger.Infof("RegisterPostHandle:", subPath)
	if subPath[0] != '/' {
		logger.Errorf("RegisterPostHandle subPath must begin with '/', :", subPath)
	}

	path := rootPath + subPath
	h, _, _ := rootRouter.Lookup("POST", path)
	if h != nil {
		logger.Errorf("RegisterPostHandle subPath with 'POST' has been register, subPath:", subPath)
	}

	rootRouter.POST(path, wrapPostHandleInternal(handle, true))
}

// RegisterPostHandleNoUserID 注册http post handle
func RegisterPostHandleNoUserID(subPath string, handle Handle) {
	logger.Info("RegisterPostHandleNoUserID:", subPath)
	if subPath[0] != '/' {
		logger.Errorf("RegisterPostHandleNoUserID subPath must begin with '/', :", subPath)
	}

	path := rootPath + subPath
	h, _, _ := rootRouter.Lookup("POST", path)
	if h != nil {
		logger.Errorf("RegisterPostHandleNoUserID subPath with 'POST' has been register, subPath:", subPath)
	}

	rootRouter.POST(path, wrapPostHandleInternal(handle, false))
}

// RegisterGetHandleNoUserID 注册http post handle
func RegisterGetHandleNoUserID(subPath string, handle Handle) {
	logger.Info("RegisterGetHandleNoUserID:", subPath)
	if subPath[0] != '/' {
		logger.Errorf("RegisterGetHandleNoUserID subPath must begin with '/', :", subPath)
	}

	path := rootPath + subPath
	h, _, _ := rootRouter.Lookup("GET", path)
	if h != nil {
		logger.Errorf("subPath with 'GET' has been register, subPath:", subPath)
	}

	rootRouter.GET(path, wrapGetHandleInternal(handle, false))
}
