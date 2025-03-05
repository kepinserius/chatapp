package api

import (
	"chat-app/database"
	"chat-app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getRooms(c *gin.Context) {
	var rooms []models.Room
	if err := database.DB.Find(&rooms).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch rooms"})
		return
	}

	c.JSON(http.StatusOK, rooms)
}

func createRoom(c *gin.Context) {
	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create room"})
		return
	}

	c.JSON(http.StatusOK, room)
}

func getRoomMessages(c *gin.Context) {
	roomID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	var messages []models.Message
	if err := database.DB.Where("room_id = ?", roomID).Preload("User").Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}
