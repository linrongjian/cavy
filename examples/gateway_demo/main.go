package main

import (
	"servergo/core/logger"
	"servergo/service/gateway"
)

func main() {
	g := gateway.NewGateServer()
	if g.Init() != nil {
		logger.Error("gate init err")
	}
	g.Run()
}
