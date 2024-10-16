package main

import (
	"log"
)

func main() {
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
