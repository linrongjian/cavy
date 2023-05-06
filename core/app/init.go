package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	log.Println("load env file")
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	for _, v := range os.Environ() {
		log.Println(v)
	}
}
