package httpsvr

import (
	"fmt"
	"net/http"

	////"battlePlatform/wjrgit.qianz.com/common/api"
	"cavy/common/api"

	"github.com/fvbock/endless"
)

// 监听https
func listenTLS(conf *api.Conf, certfile string, keyfile string) {
	fmt.Printf("Http Server Listen TLS %s\n", conf.Addr)

	http.Handle("/", &httpHandler{})
	err := endless.ListenAndServeTLS(conf.Addr, certfile, keyfile, nil)
	if err != nil {
		msg := fmt.Sprintf("Listen addr %s err:%v", conf.Addr, err)
		panic(msg)
	}
}

// 监听http
func listen(addr string) {
	fmt.Printf("Http Server Listen %s\n", addr)

	http.Handle("/", &httpHandler{})
	err := endless.ListenAndServe(addr, nil)
	if err != nil {
		msg := fmt.Sprintf("Listen addr %s err:%v", addr, err)
		panic(msg)
	}
}
