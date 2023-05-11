package main

import (
	"cavy/component/logicserver"
	"cavy/core/app"
)

func main() {
	g := logicserver.NewLogicServer()
	g.Init()
	app.Run(g)
}
