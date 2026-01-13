@echo off
setlocal enabledelayedexpansion

echo 🚀 Starting File Manager with Docker
echo ====================================

REM Check if Docker is installed
where docker >nul 2>nul
if %errorlevel% neq 0 (
    echo ❌ Docker is not installed. Please install Docker first.
    pause
    exit /b 1
)

REM Check if Docker Compose is installed
where docker-compose >nul 2>nul
if %errorlevel% neq 0 (
    echo ❌ Docker Compose is not installed. Please install Docker Compose first.
    pause
    exit /b 1
)

echo ✅ Docker and Docker Compose are installed

REM Create data directory if it doesn't exist
if not exist "data" (
    echo 📁 Creating data directory...
    mkdir data
    mkdir data\documents
    mkdir data\images
    
    echo Sample text file content > data\sample.txt
    echo # Sample Markdown File > data\README.md
    echo This is a **markdown** file with some content. >> data\README.md
    echo {"name": "sample", "version": "1.0", "description": "Sample JSON file"} > data\config.json
    echo Sample log entry > data\app.log
    echo Document content > data\documents\document.txt
    
    echo package main > data\main.go
    echo. >> data\main.go
    echo import "fmt" >> data\main.go
    echo. >> data\main.go
    echo func main() { >> data\main.go
    echo     fmt.Println("Hello, World!") >> data\main.go
    echo } >> data\main.go
    
    echo ✅ Sample data created in ./data
)

echo 🔧 Starting File Manager services...
docker-compose up --build

echo.
echo File Manager is running!
echo 📱 Frontend: http://localhost:3000
echo 🔧 Backend API: http://localhost:8080
echo 📁 Data directory: ./data
echo.
echo Press Ctrl+C to stop the services.
pause