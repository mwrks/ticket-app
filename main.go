package main

import (
	"ticket-app/initializers"
	"ticket-app/routes"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
}
func main() {
	r := routes.SetupRouter()

	r.Run() // listen and serve on 0.0.0.0:8080
}
