package repositories

import (
	"errors"
	"smashfriend/database"
	"smashfriend/models"
	"smashfriend/utils"
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

	result := paginatedQuery.Find(&messages)

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
	if chatRoom == nil || user == nil {
		return errors.New("chat room and user must not be nil")
	}

	if chatRoom.ID == 0 || user.ID == 0 {
		return errors.New("chat room and user must have valid IDs")
	}

	model := database.DB.Model(chatRoom)

	count := model.Where("id == ?", user.ID).Association("Users").Count()
	if count > 0 {
		return errors.New("user is already in this chat room")
	}

	err := model.Association("Users").Append(user)
	if err != nil {
		return err
	}

	return nil
}

func RemoveUserFromRoom(user *models.User, chatRoom *models.ChatRoom) error {
	if chatRoom == nil || user == nil {
		return errors.New("chat room and user must not be nil")
	}

	if chatRoom.ID == 0 || user.ID == 0 {
		return errors.New("chat room and user must have valid IDs")
	}

	model := database.DB.Model(&chatRoom)

	count := model.Where("id = ?", user.ID).Association("Users").Count()
	if count == 0 {
		return errors.New("user is not in this chat room")
	}

	err := model.Association("Users").Delete(user)
	if err != nil {
		return err
	}

	return nil
}

func CreateMessageForChatRoom(message string, user *models.User, chatRoom *models.ChatRoom) (*models.Message, error) {
	if message == "" {
		return nil, errors.New("message must not be empty")
	}

	if user == nil || chatRoom == nil {
		return nil, errors.New("chat room and user must not be nil")
	}

	if chatRoom.ID == 0 || user.ID == 0 {
		return nil, errors.New("chat room and user must have valid IDs")
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

