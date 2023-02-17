package main

import "fastgameserver/fastgameserver/loginserver"

func main() {
	g := loginserver.NewLoginServer()
	g.Init()
	g.Run()
}
