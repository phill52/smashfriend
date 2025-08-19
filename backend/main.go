package main

import (
	"log"
	"net/http"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"smashfriend/controllers"
	"smashfriend/database"
	"smashfriend/models"
	"smashfriend/repositories"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	clerkSecretKey := os.Getenv("CLERK_SECRET_KEY")
	if clerkSecretKey == "" {
		log.Fatal("CLERK_SECRET_KEY environment variable is required")
	}
	clerk.SetKey(clerkSecretKey)

	db, err := database.Connect(database.DefaultConfig())
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = database.AutoMigrate(db, &models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	router := gin.Default()

	if os.Getenv("APP_ENV") == "development" {
		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders: []string{"Content-Type", "Authorization"},
		}))
	}

	protected := router.Group("/api")
	protected.Use(clerkAuthMiddleware())
	{
		protected.GET("/users", controllers.GetUsers)
		protected.GET("/users/:id", controllers.GetUser)
		protected.POST("/users", controllers.CreateUser)

		log.Println("Server starting on :8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	}
}

func clerkAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := clerkhttp.WithHeaderAuthorization()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := clerk.SessionClaimsFromContext(r.Context())
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}

			clerkUser, err := user.Get(r.Context(), claims.Subject)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user data from Clerk"})
				c.Abort()
				return
			}

			dbUser, err := repositories.GetOrCreateUserFromClerk(clerkUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get or create user"})
				c.Abort()
				return
			}

			c.Set("user", dbUser)
			c.Set("clerk_user", clerkUser)
			c.Set("user_id", claims.Subject)

			c.Next()
		}))

		handler.ServeHTTP(c.Writer, c.Request)
	}
}
