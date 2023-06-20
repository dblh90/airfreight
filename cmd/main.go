package main

import (
	"github.com/dbl90/airfreight/internal/models/config"
	"github.com/dbl90/airfreight/internal/models/db"
	"github.com/dbl90/airfreight/internal/services"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	log.Println("Starting server...")

	log.Println("Loading config...")
	loadedConfig, err := config.LoadConfig("config/env-dev.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}
	log.Println("Config is loaded successfully")

	log.Println("Connecting to DB...")
	dbClient, err := db.NewDBClient(loadedConfig.App.DBConfig)
	defer dbClient.Close()
	if err != nil {
		log.Fatalf("Error connecting to DB: %s", err)
	}
	log.Println("DB is connected")
	apiClient := services.NewAPIClient(loadedConfig, dbClient)

	log.Println("Listening on port " + loadedConfig.App.Port)
	if apiClient.App.Listen(":" + loadedConfig.App.Port); err != nil {
		log.Fatalf("Error listening on port %s: %s", loadedConfig.App.Port, err)
	}
}
