package routes

import (
	"github.com/AlessandroCinque/GoAPI-Practice/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	autheticated := server.Group("/")

	autheticated.Use(middleware.Authenticate)
	autheticated.PUT("/events/:id",UpdateEvent)
	autheticated.DELETE("/events/:id",DeleteEvent)
	autheticated.GET("/events", GetEvents)
	autheticated.GET("/events/:id", GetEvent)
	autheticated.POST("/events", CreateEvent)
	
	server.POST("/signup", Signup)
	server.POST("/login", Login)
}