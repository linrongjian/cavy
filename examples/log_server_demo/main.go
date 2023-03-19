package main

import (
	"servergo/core/logger"
	"servergo/service/logserver"
)

func main() {
	g := logserver.NewLogServer()
	if g.Init() != nil {
		logger.Error("gate init err")
	}
	g.Run()
}
