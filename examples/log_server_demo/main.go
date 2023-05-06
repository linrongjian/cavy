package main

import (
	"CavyGo/component/logserver"
	"log"
)

func main() {
	ls, err := logserver.NewLogServer()
	if err != nil {
		log.Panicln("log server start failure")
	}
	ls.Run()
}
