package main

import "fastgameserver/fastgameserver/logicserver"

func main() {
	g := logicserver.NewLogicServer()
	g.Init()
	g.Run()
}
