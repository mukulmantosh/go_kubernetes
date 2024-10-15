package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

type DBClient interface {
	Ready() bool
	RunMigration() error
}

type Client struct {
	db *gorm.DB
}

func (c Client) Ready() bool {
	var ready string
	result := c.db.Raw("SELECT 1 as ready").Scan(&ready)
	if result.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}

func (c Client) RunMigration() error {
	if !c.Ready() {
		log.Fatal("Database is not ready")
	}
	err := c.db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}

func NewDBClient() (DBClient, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	databasePort, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatal("Invalid DB Port")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbHost, dbUsername, dbPassword, dbName, databasePort, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	client := Client{db}
	return client, nil
}
