package main

import "fastgameserver/fastgameserver/gateserver"

func main() {
	g := gateserver.NewGateServer()
	g.Init()
	g.Run()
}
