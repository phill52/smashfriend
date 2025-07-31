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
		response := utils.GetResponse(nil, nil, http.StatusBadRequest, "Page parameter is not a number")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response := utils.GetResponse(nil, nil, http.StatusBadRequest, "Limit parameter is not a number")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var paginationErr *utils.PaginationError
	paginatedUsers, err := repositories.GetPaginatedUsers(page, limit)
	if err != nil {
		if errors.As(err, &paginationErr) {
			response := utils.GetResponse(nil, nil, http.StatusBadRequest, paginationErr.Error())
			c.JSON(http.StatusBadRequest, response)
			return
		} else {
			response := utils.GetResponse(nil, nil, http.StatusInternalServerError, "")
			c.JSON(http.StatusInternalServerError, *response)
			return
		}
	}

	users := paginatedUsers.Users
	paginationData := paginatedUsers.Pagination
	response := utils.GetResponse(users, paginationData, http.StatusOK, "Success")
	c.JSON(http.StatusOK, response)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := repositories.GetUser(id)
	if err != nil {
		response := utils.GetResponse(nil, nil, http.StatusInternalServerError, "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.GetResponse(user, nil, http.StatusOK, "Success")
	c.JSON(http.StatusOK, response)
}

func CreateUser(c *gin.Context) {
	username := c.PostForm("username")

	if len(username) < 3 {
		response := utils.GetResponse(nil, nil, http.StatusBadRequest, "Username must be at least 3 characters long")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	existing_user, err := repositories.GetUserByUsername(username)
	if err != nil {
		response := utils.GetResponse(nil, nil, http.StatusInternalServerError, "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if existing_user != nil {
		response := utils.GetResponse(nil, nil, http.StatusBadRequest, "A user with this username already exists")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := repositories.CreateUser(username)
	if err != nil {
		response := utils.GetResponse(nil, nil, http.StatusInternalServerError, "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := utils.GetResponse(user, nil, http.StatusOK, "Success")
	c.JSON(http.StatusOK, response)
}
