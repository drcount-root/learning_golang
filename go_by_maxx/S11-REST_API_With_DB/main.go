package main

import (
	"example.com/rest_api_with_db/db"
	"example.com/rest_api_with_db/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
