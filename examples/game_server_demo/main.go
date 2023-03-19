package main

import "servergo/service/logicserver"

func main() {
	g := logicserver.NewLogicServer()
	g.Init()
	g.Run()
}
