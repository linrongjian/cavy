package main

import "eventgo/component/logicserver"

func main() {
	g := logicserver.NewLogicServer()
	g.Init()
	g.Run()
}
