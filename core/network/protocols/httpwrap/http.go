package httpwrap

import (
	"context"
	"servergo/core/logger"
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
	Ctx        context.Context
)

// GetVersion 版本号
func GetVersion() int {
	return versionCode
}

func echoVersion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("version:%v", versionCode)))
}

func CreateHTTPServer() {
	rootRouter.Handle("GET", rootPath+"/version", echoVersion)

	go acceptHTTPRequest()
}

func acceptHTTPRequest() {
	var h http.Handler
	if Conf.Test {
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

	portStr := fmt.Sprintf(":%d", Conf.Port)
	s := &http.Server{
		Addr:           portStr,
		Handler:        nil,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 10,
	}

	logger.Infof("Http server listen at:%d\n", Conf.Port)

	err := s.ListenAndServe()
	if err != nil {
		logger.Errorf("Http server ListenAndServe %d failed:%s\n", Conf.Port, err)
	}
}
