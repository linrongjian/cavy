package main

import "CavyGo/component/logicserver"

func main() {
	g := logicserver.NewLogicServer()
	g.Init()
	g.Run()
}
