package main

import (
	"github.com/AlessandroCinque/GoAPI-Practice/db"
	"github.com/AlessandroCinque/GoAPI-Practice/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	
	server.Run(":8080") //localhost:8080
}