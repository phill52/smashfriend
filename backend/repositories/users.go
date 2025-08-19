package repositories

import (
	"errors"
	"smashfriend/database"
	"smashfriend/models"
	"smashfriend/utils"
	"time"

	"github.com/clerk/clerk-sdk-go/v2"
)

type PaginatedUsers struct {
	Users      []models.User
	Pagination utils.PaginationData
}

func GetPaginatedUsers(page, limit int) (*PaginatedUsers, error) {
	var users []models.User
	query := database.DB.Model(&users)

	paginatedQuery, paginationData, err := utils.PaginateData(query, page, limit)
	if err != nil {
		return nil, err
	}

	result := paginatedQuery.Find(&users)

	paginatedUsers := &PaginatedUsers{
		Users:      users,
		Pagination: *paginationData,
	}
	return paginatedUsers, result.Error
}

func GetUser(id string) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetOrCreateUserFromClerk(clerkUser *clerk.User) (*models.User, error) {
	var user models.User
	result := database.DB.Where("clerk_id = ?", clerkUser.ID).First(&user)
	if result.Error == nil {
		updates := map[string]interface{}{}
		if clerkUser.ImageURL != nil && *clerkUser.ImageURL != user.ProfilePicture {
			updates["profile_picture"] = clerkUser.ImageURL
		}

		database.DB.Model(&user).Updates(updates)
		return &user, nil
	}

	user = models.User{
		ClerkID: clerkUser.ID,
		Username: "",
		DisplayName: "",
		ProfilePicture: *clerkUser.ImageURL,
		IsActive: true,
		LastSeen: &time.Time{},
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	result := database.DB.Where("username = ?", username).Limit(1).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &user, nil
}

func CreateUser(username string) (*models.User, error) {
	if username == "" {
		return nil, errors.New("username is required")
	}

	if len(username) < 3 {
		return nil, errors.New("username must be at least 3 characters long")
	}

	existing_user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if existing_user != nil {
		return nil, errors.New("a user with this username already exists")
	}

	user := models.User{
		Username: username,
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
