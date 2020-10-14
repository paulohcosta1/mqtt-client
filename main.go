package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()


	if err != nil {
		log.Fatal("Something is wrong with .env file")
	}

	api()
}
