package models


import (
	"gorm.io/gorm"
)

type ChatRoom struct {
	gorm.Model

	Name        string  `json:"name" gorm:"uniqueIndex;not null"`
	Description *string `json:"description`

	Users []User `gorm:"many2many:chat_room_users;"`
}

type Chat struct {
	gorm.Model

	ChatRoom   ChatRoom
	ChatRoomID uint

	CreatedBy   User
	CreatedByID uint

	Message string `json:"message" gorm:"not null"`
}
