package main

import (
	"cavy"
	"log"
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
