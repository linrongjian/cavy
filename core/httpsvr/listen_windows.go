package httpsvr

import (
	"fmt"
	"net/http"

	"github.com/linrongjian/cavy/common/api"
)

////"battlePlatform/wjrgit.qianz.com/common/api")

// 监听https
func listenTLS(conf *api.Conf, certfile string, keyfile string) {
	fmt.Printf("Http Server Listen TLS %s\n", conf.Addr)

	http.Handle("/", &httpHandler{})
	err := http.ListenAndServeTLS(conf.Addr, certfile, keyfile, nil)
	if err != nil {
		msg := fmt.Sprintf("Listen addr %s err:%v", conf.Addr, err)
		panic(msg)
	}
}

// 监听http
func listen(addr string) {
	fmt.Printf("Http Server Listen %s\n", addr)

	http.Handle("/", &httpHandler{})
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		msg := fmt.Sprintf("Listen addr %s err:%v", addr, err)
		panic(msg)
	}
}
