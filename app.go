package main

import (
	"kardashian_api/config"
	"kardashian_api/database"
	"kardashian_api/routes"
	"log"
)

func init() {
	errEnv := config.LoadEnv()
	if errEnv != nil {
		log.Fatalf("Fatal Error: %v", errEnv)
	}
	errDB := database.LoadClient()
	if errDB != nil {
		log.Fatalf("Fatal Error: %v", errDB)
	}

}

func main() {
	routes.LoadRoutes()
	err := routes.Run()
	if err != nil {
		log.Fatalf("Fatal Error: %v", err)
	}

}
