package main

import (
	"log"

	"github.com/linrongjian/cavy"
)

func main() {
	_, err := cavy.NewLogconsumer(handle)
	if err != nil {
		log.Panicf("logconsumer err: %s", err)
	}

	select {}
}

func handle(body []byte) {
	log.Println(string(body))
}
