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

    There's an `env.example` that has everything we need for now. We may need to add more thing and instructions to it later,
    but for now this is all we need. You can just run

    ```sh
    cp env.example .env
    ```

    to create the env file.

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
