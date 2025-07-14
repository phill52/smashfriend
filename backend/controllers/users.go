package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"smashfriend/models"
	"smashfriend/repositories"
	"smashfriend/utils"
	"strconv"
)

func GetUsers(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	pagination := utils.GetPaginationData(page, limit, models.User{})

	users, err := repositories.GetUsers(pagination.Offset, pagination.Limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := repositories.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	username := c.PostForm("username")

	if len(username) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username must be at least 3 characters long"})
		return
	}

	existing_user, err := repositories.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if existing_user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A user with this username already exists"})
		return
	}

	user, err := repositories.CreateUser(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
