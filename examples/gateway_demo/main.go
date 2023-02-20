package main

import (
	"fastserver/core/logger"
	"fastserver/service/gateway"
)

func main() {
	g := gateway.NewGateServer()
	if g.Init() != nil {
		logger.Error("gate init err")
	}
	g.Run()
}
