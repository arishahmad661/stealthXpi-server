package main

import (
	"github.com/arishahmad661/stealth_x_pi/config"
	"github.com/arishahmad661/stealth_x_pi/internal/db"
	"log"
)

func main() {
	config.LoadEnv()
	db.InitMongoDB()

	r := config.SetupRoutes()

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server.")
	}
}
