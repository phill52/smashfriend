package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"smashfriend/controllers"
	"smashfriend/database"
	"smashfriend/models"
)

func main() {
	db, err := database.Connect(database.DefaultConfig())
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = database.AutoMigrate(db, &models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	router := gin.Default()

	router.GET("/users", controllers.GetUsers)
	router.GET("/users/page/:page", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
