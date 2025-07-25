package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"smashfriend/repositories"
	"smashfriend/utils"
)

func GetUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		response, err := repositories.GetResponseWithoutPagination(nil, http.StatusBadRequest, "page parameter not valid")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit parameter not valid"})
		return
	}

	var paginationErr *utils.PaginationError
	users, err := repositories.GetUsers(page, limit)
	if err != nil {
		if errors.As(err, &paginationErr) {
			c.JSON(http.StatusBadRequest, gin.H{"error": paginationErr.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	response, err := repositories.GetResponse(users, page, limit, http.StatusOK, "Success")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := repositories.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response, err := repositories.GetResponseWithoutPagination(user, http.StatusOK, "Success")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
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
