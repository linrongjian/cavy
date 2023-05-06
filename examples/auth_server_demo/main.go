package main

import "CavyGo/component/auth"

func main() {
	g := auth.NewLoginServer()
	g.Init()
	g.Run()
}
