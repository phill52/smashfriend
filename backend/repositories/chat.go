package repositories

import (
	"errors"
	"fmt"
	"smashfriend/database"
	"smashfriend/models"
	"smashfriend/utils"

	"gorm.io/gorm"
)

type PaginatedMessages struct {
	Messages []models.Message
	Pagination utils.PaginationData
}

func GetChatRooms() ([]models.ChatRoom, error) {
	var chatRooms []models.ChatRoom
	result := database.DB.Find(&chatRooms)

	return chatRooms, result.Error
}

func GetMessagesForRoom(page, limit int) (*PaginatedMessages, error) {
	var messages []models.Message
	query := database.DB.Model(&messages)

	paginatedQuery, paginationData, err := utils.PaginateData(query, page, limit)
	if err != nil {
		return nil, err
	}

	result := paginatedQuery.Order("created_at desc").Find(&messages)

	paginatedMessages := &PaginatedMessages{
		Messages: messages,
		Pagination: *paginationData,
	}
	return paginatedMessages, result.Error
}

func CreateChatRoom(name string, description *string) (*models.ChatRoom, error) {
	if name == "" {
		return nil, errors.New("chatroom name is required")
	}

	chatRoom := models.ChatRoom{
		Name: name,
		Description: description,
		Users: []models.User{},
	}
	result := database.DB.Create(&chatRoom)
	if result.Error != nil {
		return nil, result.Error
	}
	return &chatRoom, nil
}

func AddUserToRoom(user *models.User, chatRoom *models.ChatRoom) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Error; err != nil {
			return err
		}

		model := tx.Model(&chatRoom)
		count := model.Where("id = ?", user.ID).Association("chat_room_users").Count()
		if count > 0 {
			return fmt.Errorf("user with id %d already exists in this chat room", user.ID)
		}

		if err := model.Association("chat_room_users").Append(&user); err != nil {
			return err
		}
		return nil
	})
	return err
}

func RemoveUserFromRoom(user *models.User, chatRoom *models.ChatRoom) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Error; err != nil {
			return err
		}

		model := tx.Model(&chatRoom)
		count := model.Where("id = ?", user.ID).Association("chat_room_users").Count()
		if count == 0 {
			return fmt.Errorf("user with id %d does not exist in this chat room", user.ID)
		}

		if err := model.Association("chat_room_users").Delete(&user); err != nil {
			return err
		}

		return nil

	})

	return err
}

func CreateMessageForChatRoom(message string, user *models.User, chatRoom *models.ChatRoom) (*models.Message, error) {
	if message == "" {
		return nil, errors.New("message must not be empty")
	}

	chatMessage := models.Message{
		ChatRoom: *chatRoom,
		ChatRoomID: chatRoom.ID,
		CreatedBy: *user,
		CreatedByID: user.ID,
		Message: message,
	}

	result := database.DB.Create(&chatMessage)
	if result.Error != nil {
		return nil, result.Error
	}

	return &chatMessage, nil
}

