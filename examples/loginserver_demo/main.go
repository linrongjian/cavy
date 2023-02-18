package main

import "fastgameserver/framework/gamelogin"

func main() {
	g := gamelogin.NewLoginServer()
	g.Init()
	g.Run()
}
