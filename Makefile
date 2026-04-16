.PHONY: build run test clean install dev

build:
	@echo "Building server..."
	CGO_ENABLED=1 go build -o bookstore-server cmd/server/main.go

run: build
	@echo "Starting server..."
	./bookstore-server

dev:
	@echo "Running in development mode..."
	CGO_ENABLED=1 go run cmd/server/main.go

test:
	@echo "Running API tests..."
	./test_api.sh

install:
	@echo "Installing dependencies..."
	go mod download

clean:
	@echo "Cleaning up..."
	rm -f bookstore-server
	rm -rf data/

deps:
	@echo "Checking dependencies..."
	go mod tidy
	go mod verify

help:
	@echo "Available commands:"
	@echo "  make build   - Build the server binary"
	@echo "  make run     - Build and run the server"
	@echo "  make dev     - Run in development mode"
	@echo "  make test    - Run API tests"
	@echo "  make install - Install dependencies"
	@echo "  make clean   - Clean build artifacts and database"
	@echo "  make deps    - Tidy and verify dependencies"
