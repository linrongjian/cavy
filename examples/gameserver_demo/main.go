package main

import "fastgameserver/service/business"

func main() {
	g := business.NewLogicServer()
	g.Init()
	g.Run()
}
