package main

import (
	"github.com/talles-morais/gin-rest-api/database"
	"github.com/talles-morais/gin-rest-api/routes"
)

func main() {
	database.Connect()

	routes.HandleRequests()
}
