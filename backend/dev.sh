#!/bin/bash

SERVER_PID=""

kill_server() {
    if [ -n "$SERVER_PID" ]; then
        kill $SERVER_PID
        wait $SERVER_PID
    fi
}

start_server() {
    go run main.go &
    SERVER_PID=$!
}

cleanup() {
    kill_server
    exit 0
}

trap cleanup SIGINT SIGTERM

start_server

fswatch -xr . | while read event; do
    if [[ "$event" == *go* ]]; then
        echo "Restarting server..."
        kill_server
        start_server
    fi
done
