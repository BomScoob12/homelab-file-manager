package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("🧪 MIME Type Test")
	fmt.Println("=================")

	// Wait a moment for server to be ready
	time.Sleep(2 * time.Second)

	// Test the list endpoint
	resp, err := http.Get("http://localhost:8080/file/list?path=/")
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("❌ JSON decode error: %v\n", err)
		return
	}

	fmt.Printf("📡 Response Status: %s\n", resp.Status)
	
	if items, ok := result["items"].([]interface{}); ok {
		fmt.Printf("📁 Found %d items:\n", len(items))
		for i, item := range items {
			if itemMap, ok := item.(map[string]interface{}); ok {
				name := itemMap["name"]
				mimeType := itemMap["mimeType"]
				isDir := itemMap["isDir"]
				fmt.Printf("  %d. %s (dir: %v, mime: %s)\n", i+1, name, isDir, mimeType)
			}
		}
	} else {
		fmt.Printf("❌ No items found in response\n")
		prettyJSON, _ := json.MarshalIndent(result, "", "  ")
		fmt.Printf("Raw response: %s\n", string(prettyJSON))
	}
}