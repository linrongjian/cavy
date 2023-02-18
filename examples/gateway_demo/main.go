package main

import "fastgameserver/framework/gamegate"

func main() {
	g := gamegate.NewGateServer()
	g.Init()
	g.Run()
}
