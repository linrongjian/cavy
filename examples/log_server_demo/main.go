package main

import (
	"eventgo/component/logserver"
	"eventgo/core/logger"
)

func main() {
	g := logserver.NewLogServer()
	if g.Init() != nil {
		logger.Error("gate init err")
	}
	g.Run()
}
