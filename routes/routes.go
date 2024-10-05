package routes

import (
	"emkaytec.io/events-api/v2/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// event routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// server.POST("/events", middleware.Authenticate, createEvent)
	// server.PUT("/events/:id", middleware.Authenticate, updateEvent)
	// server.DELETE("/events/:id", middleware.Authenticate, deleteEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// user routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
