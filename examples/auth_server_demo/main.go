package main

import "eventgo/component/auth"

func main() {
	g := auth.NewLoginServer()
	g.Init()
	g.Run()
}
