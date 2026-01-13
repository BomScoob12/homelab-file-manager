@echo off
setlocal enabledelayedexpansion

title File Manager - Development Environment

echo.
echo 🚀 File Manager Development Environment
echo =======================================
echo.

REM Function to check if command exists
call :check_command docker "Docker"
call :check_command docker-compose "Docker Compose"

echo ✅ All dependencies are installed
echo.

REM Setup data directory with sample files
call :setup_data

echo 🔧 Building and starting services...
echo    This may take a few minutes on first run...
echo.

REM Start Docker Compose with better error handling
docker-compose up --build
set compose_exit=%errorlevel%

echo.
if %compose_exit% equ 0 (
    echo ✅ Services stopped successfully
) else (
    echo ❌ Services stopped with errors (exit code: %compose_exit%)
)

echo.
echo 👋 Thanks for using File Manager!
pause
exit /b %compose_exit%

REM ============================================================================
REM Functions
REM ============================================================================

:check_command
where %1 >nul 2>nul
if %errorlevel% neq 0 (
    echo ❌ %2 is not installed or not in PATH
    echo    Please install %2 and try again
    echo    Visit: https://docs.docker.com/get-docker/
    echo.
    pause
    exit /b 1
)
goto :eof

:setup_data
if exist "data" (
    echo 📁 Data directory already exists
    goto :eof
)

echo 📁 Creating data directory with sample files...

REM Create directory structure
mkdir data 2>nul
mkdir data\documents 2>nul
mkdir data\images 2>nul
mkdir data\logs 2>nul

REM Create sample files
(
echo Sample text file for testing the file manager.
echo You can edit, delete, or create new files through the web interface.
) > data\sample.txt

(
echo # File Manager Sample
echo.
echo This is a **markdown** file with some content.
echo.
echo ## Features
echo - Browse files and directories
echo - View file contents
echo - Delete files
echo - Responsive web interface
echo.
echo ## Usage
echo Navigate to http://localhost:3000 to start using the file manager.
) > data\README.md

(
echo {
echo   "name": "file-manager-config",
echo   "version": "1.0.0",
echo   "description": "Sample configuration file",
echo   "settings": {
echo     "theme": "dark",
echo     "autoSave": true,
echo     "maxFileSize": "10MB"
echo   }
echo }
) > data\config.json

echo [%date% %time%] File Manager started > data\logs\app.log
echo [%date% %time%] Sample data created >> data\logs\app.log

echo Document content for testing > data\documents\document.txt
echo Meeting notes from today's standup > data\documents\notes.txt

REM Create a simple SVG image for testing
(
echo ^<svg width="200" height="100" xmlns="http://www.w3.org/2000/svg"^>
echo   ^<rect width="200" height="100" fill="lightblue"/^>
echo   ^<text x="100" y="50" text-anchor="middle" dy=".3em" font-family="Arial" font-size="16"^>Sample Image^</text^>
echo ^</svg^>
) > data\images\sample.svg

(
echo package main
echo.
echo import "fmt"
echo.
echo func main^(^) {
echo 	fmt.Println^("Hello, File Manager!"^)
echo 	fmt.Println^("This is a sample Go file"^)
echo }
) > data\main.go

echo ✅ Sample data created in .\data
echo.
goto :eof