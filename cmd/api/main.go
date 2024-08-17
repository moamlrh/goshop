package main

import (
	"log"

	"github.com/moamlrh/goshop/internal/config"
	"github.com/moamlrh/goshop/internal/database"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	db, err := database.NewPostgresDb()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()
	log.Println("Successfully connected to the database")
}
