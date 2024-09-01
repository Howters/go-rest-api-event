package main

import (
	"github.com/gin-gonic/gin"
	"github.com/howters/golang/db"
	"github.com/howters/golang/routes"
)


func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}


