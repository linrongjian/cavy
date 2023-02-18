package main

import (
	"fastgameserver/framework/gamegate"
	"fastgameserver/framework/logger"
)

func main() {
	g := gamegate.NewGateServer()
	if g.Init() != nil {
		logger.Error("gate init err")
	}
	g.Run()
}
