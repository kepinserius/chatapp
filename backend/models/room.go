package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name     string    `json:"name" gorm:"unique"`
	Users    []User    `json:"users" gorm:"many2many:user_rooms;"`
	Messages []Message `json:"messages"`
}
