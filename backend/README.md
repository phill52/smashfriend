# Smashfriend Backend

## Prerequisites

-   Go 1.24.3 or higher
-   Docker

## Setup

1. **Install dependencies:**

    ```bash
    go mod tidy
    ```

2. **Set up the env:**

    There's an `env.example`. Run

    ```sh
    cp env.example .env
    ```

    to create the env file, and put in the API keys (from the discord channel). **Make sure the API keys are protected and NEVER committed**

    **Note**: The `.env` files for the frontend and backend are different in development. You need to have the file in both directories
    for them to apply when you run the application.

3. **Set up the database:**
   The database is set up in docker for now. If you want to run your own database, that's fine, just change the env file accordingly.
   To run the docker database, just run

    ```sh
    docker-compose up -d
    ```

    Whenever you are done with development, make sure you run this command to stop it so you don't waste resources on your computer.

    ```sh
    docker-compose stop
    ```

    If you ever need to check stuff in the database directly, there is a script, `dbshell.sh`, that runs postgres in the command line.

4. **Run the application**

    ```sh
    go run main.go
    ```
