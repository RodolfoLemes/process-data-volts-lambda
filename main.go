package main

import (
	"log"
	"process-data-volts-lambda/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// lambda.Start(handlers.HandleAPI)
	handlers.HandleAPI()
}
