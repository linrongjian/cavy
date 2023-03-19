package main

import "servergo/service/auth"

func main() {
	g := auth.NewLoginServer()
	g.Init()
	g.Run()
}
