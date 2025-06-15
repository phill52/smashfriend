# SmashFriend Chat Server

A WebSocket-based chat server that supports multiple rooms and real-time messaging.

## Setup

1. Install Go if you haven't already [golang.org/dl](https://golang.org/dl)
2. Clone the repository
3. Install dependencies:

    ```bash
    go mod download
    ```

4. Start the server:

    ```bash
    go run main.go
    ```

    The server will start on `http://localhost:8080`

5. Verify it's running by visiting `http://localhost:8080` in your browser
    - You should see "Hello world"

## Server Details

-   **Base URL**: `http://localhost:8080/` (returns "Hello world")
-   **Websocket URL**: `ws://localhost:8080/ws`

## Connection

To establish a WebSocket connection, send a POST request to `/ws` with the following form data:

-   `username`: (required)

Example using TypeScript:

```ts
const formData = new FormData();
formData.append("username", "yourUsername");

const response = await fetch("http://localhost:8080/ws", {
    method: "POST",
    body: formData,
});

const ws = new WebSocket(response.url);
```

## Message Format

All messages are sent and received in JSON format with the following structure:

```json
{
    action: string; // "join", "leave", "message", "system", "get users"
    user: string; // username
    body: any; // Message content string or JSON
    room: string; // room ID
}
```

## Flow

### 1. Join a room

```json
{
    "action": "join",
    "room": "0", // Room ID (0-2)
    "user": "username",
    "body": null
}
```

Response: Broadcast to all users in the room

```json
{
    "action": "join",
    "user": "system",
    "body": "\"username joined the chat\"",
    "room": "0"
}
```

### 2. Send Message

```json
{
    "action": "message",
    "room": "0",
    "user": "username",
    "body": "Your message here"
}
```

Response: Broadcast to all users in the room

```json
{
    "action": "message",
    "user": "username",
    "body": "Your message here",
    "room": "0"
}
```

### 3. Leave Room

```json
{
    "action": "leave",
    "room": "0",
    "user": "username",
    "body": null
}
```

Response: Broadcast to all users in the room

```json
{
    "action": "leave",
    "user": "system",
    "body": "\"username left the chat\"",
    "room": "0"
}
```

### 4. Get Users in Room

```json
{
    "action": "get users",
    "room": "0",
    "user": "username",
    "body": null
}
```

Response: Only to requesting user

```json
{
    "action": "users",
    "user": "system",
    "body": ["user1", "user2", "user3"],
    "room": "0"
}
```

## Error Messages

The server may respond with error messages in the following format:

```json
{
    "action": "error",
    "body": "Error message here"
}
```
