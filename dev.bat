@echo off
setlocal enabledelayedexpansion

echo 🚀 Starting File Manager Development Environment
echo ================================================

REM Check if Go is installed
where go >nul 2>nul
if %errorlevel% neq 0 (
    echo ❌ Go is not installed. Please install Go 1.21 or later.
    pause
    exit /b 1
)

REM Check if Node.js is installed
where node >nul 2>nul
if %errorlevel% neq 0 (
    echo ❌ Node.js is not installed. Please install Node.js 16 or later.
    pause
    exit /b 1
)

REM Check if npm is installed
where npm >nul 2>nul
if %errorlevel% neq 0 (
    echo ❌ npm is not installed. Please install npm.
    pause
    exit /b 1
)

echo ✅ All prerequisites are installed

REM Setup test environment
echo Setting up test environment...
cd backend
call make setup-test
cd ..

REM Install frontend dependencies if needed
if not exist "frontend\node_modules" (
    echo Installing frontend dependencies...
    cd frontend
    call npm install
    cd ..
)

echo Starting backend server...
cd backend
start "Backend Server" cmd /k "go run cmd/app/main.go"
cd ..

REM Wait for backend to start
timeout /t 3 /nobreak >nul

echo Starting frontend development server...
cd frontend
start "Frontend Server" cmd /k "npm run dev"
cd ..

echo.
echo 🎉 Development environment is ready!
echo 📱 Frontend: http://localhost:3000
echo 🔧 Backend API: http://localhost:8080
echo 📁 Test files: /tmp/WorkDir
echo.
echo Press any key to open the frontend in your browser...
pause >nul

start http://localhost:3000

echo.
echo Both servers are running in separate windows.
echo Close those windows to stop the servers.
pause