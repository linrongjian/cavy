package main

import (
	"github.com/linrongjian/cavy/component/auth"
	"github.com/linrongjian/cavy/core/app"
)

func main() {
	g := auth.NewLoginServer()
	g.Init()
	app.Run(g)
}
