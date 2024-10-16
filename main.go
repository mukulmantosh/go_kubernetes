package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := NewDBClient()
	if err != nil {
		log.Fatalf("DB Error: %s\n", err)
		return
	}

	err = db.RunMigration()
	if err != nil {
		log.Fatalf("Migration Failed: %s\n", err)
		return
	}
	service := NewServer(db)
	log.Fatal(service.Start())
}
