package routes

import (
	"event-management/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deletEvent)
	authenticated.POST("/event/:id/register", registerForEvent)
	authenticated.DELETE("/event/:id/register", cancelRegistration)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// user routes
	server.POST("/signup", signUp)

	server.POST("/login", login)

}
