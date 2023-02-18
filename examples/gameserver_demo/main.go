package main

import "fastgameserver/framework/gamelogic"

func main() {
	g := gamelogic.NewLogicServer()
	g.Init()
	g.Run()
}
