#!/bin/bash

# File Manager Docker Startup Script

set -e

echo "🚀 Starting File Manager with Docker"
echo "===================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo -e "${RED}❌ Docker is not installed. Please install Docker first.${NC}"
    exit 1
fi

# Check if Docker Compose is installed
if ! command -v docker-compose &> /dev/null; then
    echo -e "${RED}❌ Docker Compose is not installed. Please install Docker Compose first.${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Docker and Docker Compose are installed${NC}"

# Create data directory if it doesn't exist
if [ ! -d "./data" ]; then
    echo -e "${BLUE}📁 Creating data directory...${NC}"
    mkdir -p ./data
    
    # Create sample files
    echo "Sample text file content" > ./data/sample.txt
    echo "# Sample Markdown File" > ./data/README.md
    echo "This is a **markdown** file with some content." >> ./data/README.md
    echo '{"name": "sample", "version": "1.0", "description": "Sample JSON file"}' > ./data/config.json
    echo "Sample log entry" > ./data/app.log
    
    # Create subdirectories
    mkdir -p ./data/documents
    mkdir -p ./data/images
    echo "Document content" > ./data/documents/document.txt
    
    # Create a Go file
    cat > ./data/main.go << 'EOF'
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
EOF
    
    echo -e "${GREEN}✅ Sample data created in ./data${NC}"
fi

# Function to cleanup on exit
cleanup() {
    echo -e "\n${YELLOW}🛑 Shutting down File Manager...${NC}"
    docker-compose down
    exit 0
}

# Set trap to cleanup on script exit
trap cleanup SIGINT SIGTERM

# Start the services
echo -e "${BLUE}🔧 Starting File Manager services...${NC}"
docker-compose up --build

# This line should never be reached due to the trap, but just in case
cleanup