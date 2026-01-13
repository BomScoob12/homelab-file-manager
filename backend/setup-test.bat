@echo off
echo Setting up test environment...

REM Create test directory
if not exist "C:\tmp\WorkDir" mkdir "C:\tmp\WorkDir"
if not exist "C:\tmp\WorkDir\subdir" mkdir "C:\tmp\WorkDir\subdir"

REM Create test files
echo Test file content > "C:\tmp\WorkDir\test.txt"
echo Another test file > "C:\tmp\WorkDir\readme.md"
echo Subdirectory file > "C:\tmp\WorkDir\subdir\file.txt"
echo { "name": "test", "version": "1.0" } > "C:\tmp\WorkDir\config.json"
echo This is a log file > "C:\tmp\WorkDir\app.log"

REM Create a sample Go file
echo package main > "C:\tmp\WorkDir\main.go"
echo. >> "C:\tmp\WorkDir\main.go"
echo import "fmt" >> "C:\tmp\WorkDir\main.go"
echo. >> "C:\tmp\WorkDir\main.go"
echo func main() { >> "C:\tmp\WorkDir\main.go"
echo     fmt.Println("Hello, World!") >> "C:\tmp\WorkDir\main.go"
echo } >> "C:\tmp\WorkDir\main.go"

REM Create a sample HTML file
echo ^<!DOCTYPE html^> > "C:\tmp\WorkDir\index.html"
echo ^<html^> >> "C:\tmp\WorkDir\index.html"
echo ^<head^>^<title^>Test^</title^>^</head^> >> "C:\tmp\WorkDir\index.html"
echo ^<body^>^<h1^>Hello World^</h1^>^</body^> >> "C:\tmp\WorkDir\index.html"
echo ^</html^> >> "C:\tmp\WorkDir\index.html"

echo ✅ Test environment created at C:\tmp\WorkDir
echo Files created:
dir "C:\tmp\WorkDir" /b
echo.
echo Subdirectory files:
dir "C:\tmp\WorkDir\subdir" /b