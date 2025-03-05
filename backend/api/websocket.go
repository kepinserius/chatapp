package api

import (
	"chat-app/websocket"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	gorilla "github.com/gorilla/websocket"
)

var upgrader = gorilla.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections
	},
}

func serveWs(hub *websocket.Hub, c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}

	roomID, err := strconv.Atoi(c.Query("roomId"))
	if err != nil {
		log.Println("Invalid roomId:", err)
		return
	}

	userID, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		log.Println("Invalid userId:", err)
		return
	}

	username := c.Query("username")

	client := &websocket.Client{
		Hub:      hub,
		Conn:     conn,
		Send:     make(chan []byte, 256),
		RoomID:   uint(roomID),
		UserID:   uint(userID),
		Username: username,
	}

	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
