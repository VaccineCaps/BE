package main

import (
	"BE/server"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	app := server.Server()
	err := app.Start(":8080")
	if err != nil {
		return
	}
}
