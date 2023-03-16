package main

import "fastserver/service/logicserver"

func main() {
	g := logicserver.NewLogicServer()
	g.Init()
	g.Run()
}
