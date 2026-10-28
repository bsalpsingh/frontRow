package main

import (
	"log"

	"github.com/bsalpsingh/frontRow/internal/db"
	"github.com/bsalpsingh/frontRow/internal/router"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if _, err := db.Init(); err != nil {
		panic("failed to init db")
	}
	router.Init()

}
