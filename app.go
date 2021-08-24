package main

import (
	"kardashian_api/config"
	"kardashian_api/database"
	"kardashian_api/routes"
)

func init() {
	config.LoadEnv()
	database.Connect()

}

func main() {
	r := routes.Routes()
	r.Run()
}
