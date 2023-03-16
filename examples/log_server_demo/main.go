package main

import (
	"fastserver/core/logger"
	"fastserver/service/logserver"
)

func main() {
	g := logserver.NewLogServer()
	if g.Init() != nil {
		logger.Error("gate init err")
	}
	g.Run()
}
