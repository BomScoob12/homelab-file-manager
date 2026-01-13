# File Management API Makefile - Docker focused

.PHONY: build run test clean dev docker-build docker-run docker-dev

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build parameters
BINARY_NAME=file-manager
BINARY_PATH=./cmd/app
MAIN_PATH=./cmd/app/main.go

# Docker parameters
DOCKER_IMAGE=file-manager
DOCKER_TAG=latest
DATA_DIR=./data

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PATH)

# Run the application locally
run: setup-data
	$(GOCMD) run $(MAIN_PATH)

# Run in development mode with auto-reload (requires air)
dev: setup-data
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "Air not installed. Install with: go install github.com/cosmtrek/air@latest"; \
		echo "Running without auto-reload..."; \
		$(GOCMD) run $(MAIN_PATH); \
	fi

# Test the application
test:
	$(GOTEST) -v ./...

# Test with coverage
test-coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html

# Download dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Format code
fmt:
	$(GOCMD) fmt ./...

# Lint code (requires golangci-lint)
lint:
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install from: https://golangci-lint.run/usage/install/"; \
	fi

# Run test client
test-client:
	$(GOCMD) run test_api.go

# Create local data directory and sample files
setup-data:
	@echo "🔧 Setting up local data directory..."
	@mkdir -p $(DATA_DIR)
	@mkdir -p $(DATA_DIR)/documents
	@mkdir -p $(DATA_DIR)/images
	@echo "Sample text file content" > $(DATA_DIR)/sample.txt
	@echo "# Sample Markdown File" > $(DATA_DIR)/README.md
	@echo "This is a **markdown** file with some content." >> $(DATA_DIR)/README.md
	@echo '{"name": "sample", "version": "1.0", "description": "Sample JSON file"}' > $(DATA_DIR)/config.json
	@echo "Sample log entry" > $(DATA_DIR)/app.log
	@echo "Document content" > $(DATA_DIR)/documents/document.txt
	@echo "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}" > $(DATA_DIR)/main.go
	@echo "✅ Local data directory created at $(DATA_DIR)"

# Build Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

# Run with Docker
docker-run: docker-build
	docker run --rm -p 8080:8080 -v $(PWD)/data:/data $(DOCKER_IMAGE):$(DOCKER_TAG)

# Run with Docker Compose
docker-dev:
	docker-compose up --build

# Stop Docker Compose
docker-stop:
	docker-compose down

# Show help
help:
	@echo "Available commands:"
	@echo "  build         - Build the application"
	@echo "  run           - Run the application locally (includes setup-data)"
	@echo "  dev           - Run in development mode with auto-reload"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  clean         - Clean build artifacts"
	@echo "  deps          - Download and tidy dependencies"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code"
	@echo "  test-client   - Run API test client"
	@echo "  setup-data    - Create local data directory and sample files"
	@echo "  docker-build  - Build Docker image"
	@echo "  docker-run    - Run with Docker"
	@echo "  docker-dev    - Run with Docker Compose (recommended)"
	@echo "  docker-stop   - Stop Docker Compose"
	@echo "  help          - Show this help message"

# Default target
all: deps fmt build test