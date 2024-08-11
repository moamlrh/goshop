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
		panic(err)
	}

	defer db.Close()
	log.Println("Successfully connected to the database")
}
