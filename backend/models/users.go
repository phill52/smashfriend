package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ClerkID string `json:"clerk_id" gorm:"uniqueIndex;not null"`

	Username    string `json:"username" gorm:"uniqueIndex;not null"`
	DisplayName string `json:"display_name"`

	ProfilePicture string `json:"profile_picture"`

	IsActive bool      `json:"is_active" gorm:"default:true"`
	LastSeen *time.Time `json:"last_seen"`
}
