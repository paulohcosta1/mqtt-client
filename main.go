package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Something is wrong with .env file")
	}

	uri, err := url.Parse(os.Getenv("RAW_URL"))

	if err != nil {
		log.Fatal(err)
	}

	api(uri)
}
