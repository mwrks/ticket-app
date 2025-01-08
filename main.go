package main

import (
	"github.com/mwrks/ticket-app/initializers"
	"github.com/mwrks/ticket-app/routes"
)

func init() {
	initializers.LoadEnv()         // Load environment variables
	initializers.ConnectDatabase() // Connecting to database
}

func main() {
	// Initialize router
	r := routes.SetupRouter()

	// Listen and serve on 0.0.0.0:8080
	r.Run()
}
