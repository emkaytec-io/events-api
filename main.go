package main

import (
	"emkaytec.io/events-api/v2/db"
	"emkaytec.io/events-api/v2/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
