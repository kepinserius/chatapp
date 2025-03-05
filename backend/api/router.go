package api

import (
	"chat-app/websocket"

	"github.com/gin-gonic/gin"
)

var hub *websocket.Hub

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Initialize websocket hub
	hub = websocket.NewHub()
	go hub.Run()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API routes
	api := r.Group("/api")
	{
		api.POST("/register", registerUser)
		api.POST("/login", loginUser)
		api.GET("/rooms", authMiddleware(), getRooms)
		api.POST("/rooms", authMiddleware(), createRoom)
		api.GET("/rooms/:id/messages", authMiddleware(), getRoomMessages)
	}

	// WebSocket route
	r.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c)
	})

	return r
}
