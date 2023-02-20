package main

import "fastserver/service/business"

func main() {
	g := business.NewLogicServer()
	g.Init()
	g.Run()
}
