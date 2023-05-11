package main

import (
	"cavy/component/auth"
	"cavy/core/app"
)

func main() {
	g := auth.NewLoginServer()
	g.Init()
	app.Run(g)
}
