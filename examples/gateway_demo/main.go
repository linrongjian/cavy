package main

import (
	"fastgameserver/core/logger"
	"fastgameserver/service/gateway"
)

func main() {
	g := gateway.NewGateServer()
	if g.Init() != nil {
		logger.Error("gate init err")
	}
	g.Run()
}
