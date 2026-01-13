package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Simple test client for the file management API
func main() {
	baseURL := "http://localhost:8080/file"

	fmt.Println("File Management API Test Client")
	fmt.Println("================================")

	// Test 1: List files in root directory
	fmt.Println("\n1. Testing file listing...")
	testListFiles(baseURL, "/")

	// Test 2: Get file details (if any files exist)
	fmt.Println("\n2. Testing file details...")
	testFileDetails(baseURL, "/test.txt") // This might fail if file doesn't exist

	// Test 3: Try to open a file
	fmt.Println("\n3. Testing file opening...")
	testOpenFile(baseURL, "/test.txt") // This might fail if file doesn't exist

	// Test 4: Test error handling with invalid path
	fmt.Println("\n4. Testing error handling...")
	testErrorHandling(baseURL)
}

func testListFiles(baseURL, path string) {
	url := fmt.Sprintf("%s/list?path=%s", baseURL, url.QueryEscape(path))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	fmt.Printf("Status: %s\n", resp.Status)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		fmt.Printf("Raw response: %s\n", string(body))
		return
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("Response: %s\n", string(prettyJSON))
}

func testFileDetails(baseURL, filePath string) {
	url := fmt.Sprintf("%s/details?path=%s", baseURL, url.QueryEscape(filePath))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Response: %s\n", string(body))
}

func testOpenFile(baseURL, filePath string) {
	url := fmt.Sprintf("%s/open?path=%s", baseURL, url.QueryEscape(filePath))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Response: %s\n", string(body))
}

func testErrorHandling(baseURL string) {
	// Test with invalid path containing ../
	url := fmt.Sprintf("%s/list?path=%s", baseURL, url.QueryEscape("../etc"))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Response: %s\n", string(body))
}

// Helper function to create a test file for testing
func createTestFile() {
	content := fmt.Sprintf("Test file created at %s\nThis is a test file for the file management API.", time.Now().Format(time.RFC3339))

	err := os.WriteFile("/data/test.txt", []byte(content), 0644)
	if err != nil {
		fmt.Printf("Could not create test file: %v\n", err)
	} else {
		fmt.Println("Test file created at /data/test.txt")
	}
}