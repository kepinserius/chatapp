package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Content  string `json:"content"`
	UserID   uint   `json:"user_id"`
	User     User   `json:"user"`
	RoomID   uint   `json:"room_id"`
	Room     Room   `json:"room"`
}
