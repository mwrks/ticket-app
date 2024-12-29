package main

import (
	"fmt"
	"ticket-app/initializers"
	"ticket-app/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Ticket{}, &models.Order{})
	if err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}
}
