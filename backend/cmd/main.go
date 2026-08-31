package main

import (
	"log"

	"github.com/SA-1-69/T20/backend/internal/config"
)

func main() {
	if err := config.ConnectDatabase(); err != nil {
		log.Fatal(err)
	}
	log.Println("database connected and migrated")
}
