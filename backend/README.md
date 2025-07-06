# SmashFriend Backend

A Go backend API using GORM and Gorilla Mux with PostgreSQL.

## Prerequisites

-   Go 1.24.3 or higher
-   PostgreSQL database

## Setup

1. **Install dependencies:**

    ```bash
    go mod tidy
    ```

2. **Set up PostgreSQL database:**

    - Create a database named `smashfriend`
    - Update the database configuration in `database/database.go` if needed:
        ```go
        func DefaultConfig() *Config {
            return &Config{
                Host:     "localhost",
                User:     "postgres",
                Password: "postgres",
                DBName:   "smashfriend",
                Port:     "5432",
                SSLMode:  "disable",
            }
        }
        ```

3. **Run the application:**
    ```bash
    go run main.go
    ```

## API Endpoints

-   `GET /health` - Health check
-   `GET /api/users` - Get all users
-   `POST /api/users` - Create a new user
-   `GET /api/games` - Get all games
-   `POST /api/games` - Create a new game

## Project Structure

```
backend/
├── main.go          # Application entry point
├── go.mod           # Go module dependencies
├── database/
│   └── database.go  # Database connection and configuration
├── models/
│   └── models.go    # GORM models
├── handlers/
│   └── handlers.go  # HTTP handlers
└── README.md        # This file
```

## Database Models

-   **User**: ID, Username, Email, Password, timestamps
-   **Game**: ID, Title, Description, Genre, timestamps

The database schema will be automatically migrated when the application starts.

## Database Configuration

The database connection is handled in the `database` package:

-   `database.DefaultConfig()` - Returns default PostgreSQL configuration
-   `database.Connect(config)` - Establishes database connection
-   `database.AutoMigrate(db, models...)` - Runs database migrations

You can customize the database configuration by modifying the `DefaultConfig()` function or creating a custom `Config` struct.
