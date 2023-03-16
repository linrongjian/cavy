package main

import "fastserver/service/auth"

func main() {
	g := auth.NewLoginServer()
	g.Init()
	g.Run()
}
