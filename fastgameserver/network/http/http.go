package http

import (
	"context"
	"fastgameserver/fastgameserver/logger"
	"net/http"
	"time"

	"fmt"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

const (
	versionCode = 3.0
)

var (
	rootRouter = httprouter.New()
	rootPath   = ""

	Ctx    context.Context
	cancel context.CancelFunc
)

// GetVersion 版本号
func GetVersion() int {
	return versionCode
}

func echoVersion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("version:%v", versionCode)))
}

// CreateHTTPServer 启动服务器
func CreateHTTPServer() {
	rootRouter.Handle("GET", rootPath+"/version", echoVersion)

	go acceptHTTPRequest()
}

// acceptHTTPRequest 监听和接受HTTP
func acceptHTTPRequest() {
	var h http.Handler
	if Opts.Test {
		// 支持客户端跨域访问
		c := cors.New(cors.Options{
			AllowOriginFunc: func(origin string) bool {
				return true
			},
			AllowCredentials: true,
			AllowedHeaders:   []string{"*"},           // we need this line for cors to allow cross-origin
			ExposedHeaders:   []string{"Set-Session"}, // we need this line for cors to set Access-Control-Expose-Headers
		})
		h = c.Handler(rootRouter)
	} else {
		// 对外服务器不应该允许跨域访问
		h = rootRouter
	}

	http.Handle("/", h)

	portStr := fmt.Sprintf(":%d", Opts.Port)
	s := &http.Server{
		Addr:           portStr,
		Handler:        nil,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 10,
	}

	logger.Infof("Http server listen at:%d\n", Opts.Port)

	err := s.ListenAndServe()
	if err != nil {
		logger.Errorf("Http server ListenAndServe %d failed:%s\n", Opts.Port, err)
	}
}
