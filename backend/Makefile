# File Management API Makefile

.PHONY: build run test clean dev docker-build docker-run

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

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PATH)

# Run the application
run:
	$(GOCMD) run $(MAIN_PATH)

# Run in development mode with auto-reload (requires air)
dev:
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

# Create test directory and files
setup-test:
	@mkdir -p /tmp/WorkDir
	@echo "Test file content" > /tmp/WorkDir/test.txt
	@echo "Another test file" > /tmp/WorkDir/readme.md
	@mkdir -p /tmp/WorkDir/subdir
	@echo "Subdirectory file" > /tmp/WorkDir/subdir/file.txt
	@echo "Test environment created at /tmp/WorkDir"

# Build Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

# Run with Docker (requires Dockerfile)
docker-run:
	docker run -p 8080:8080 -v /tmp/WorkDir:/WorkDir $(DOCKER_IMAGE):$(DOCKER_TAG)

# Run with Docker and test directory
docker-run-test: setup-test docker-run

# Show help
help:
	@echo "Available commands:"
	@echo "  build         - Build the application"
	@echo "  run           - Run the application"
	@echo "  dev           - Run in development mode with auto-reload"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  clean         - Clean build artifacts"
	@echo "  deps          - Download and tidy dependencies"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code"
	@echo "  test-client   - Run API test client"
	@echo "  setup-test    - Create test directory and files"
	@echo "  docker-build  - Build Docker image"
	@echo "  docker-run    - Run with Docker"
	@echo "  help          - Show this help message"

# Default target
all: deps fmt build test