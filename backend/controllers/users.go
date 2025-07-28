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
		response := repositories.GetResponse(nil, nil, nil, http.StatusBadRequest, "Bad Request: page parameter not valid")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response := repositories.GetResponse(nil, nil, nil, http.StatusBadRequest, "Bad Request: limit parameter not valid")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var paginationErr *utils.PaginationError
	users, err := repositories.GetUsers(page, limit)
	if err != nil {
		if errors.As(err, &paginationErr) {
			response := repositories.GetResponse(nil, nil, nil, http.StatusBadRequest, paginationErr.Error())
			c.JSON(http.StatusBadRequest, response)
			return
		} else {
			response := repositories.GetResponse(nil, nil, nil, http.StatusInternalServerError, err.Error())
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	response := repositories.GetResponse(users, &page, &limit, http.StatusOK, "Success")
	c.JSON(http.StatusOK, response)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := repositories.GetUser(id)
	if err != nil {
		response := repositories.GetResponse(nil, nil, nil, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := repositories.GetResponse(user, nil, nil, http.StatusOK, "Success")
	c.JSON(http.StatusOK, response)
}

func CreateUser(c *gin.Context) {
	username := c.PostForm("username")

	if len(username) < 3 {
		response := repositories.GetResponse(nil, nil, nil, http.StatusBadRequest, "Username must be at least 3 characters long")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	existing_user, err := repositories.GetUserByUsername(username)
	if err != nil {
		response := repositories.GetResponse(nil, nil, nil, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if existing_user != nil {
		response := repositories.GetResponse(nil, nil, nil, http.StatusBadRequest, "A user with this username already exists")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := repositories.CreateUser(username)
	if err != nil {
		response := repositories.GetResponse(nil, nil, nil, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := repositories.GetResponse(user, nil, nil, http.StatusOK, "Success")
	c.JSON(http.StatusOK, response)
}
