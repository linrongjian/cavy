package main

import (
	"github.com/linrongjian/cavy/component/logicserver"
	"github.com/linrongjian/cavy/core/app"
)

func main() {
	g := logicserver.NewLogicServer()
	g.Init()
	app.Run(g)
}
