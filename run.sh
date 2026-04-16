#!/bin/bash

export CGO_ENABLED=1

if [ ! -f "bookstore-server" ]; then
    echo "Building server..."
    go build -o bookstore-server cmd/server/main.go
fi

echo "Starting BookStore API server..."
./bookstore-server
