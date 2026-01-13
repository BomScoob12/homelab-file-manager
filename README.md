# File Manager

A simple web-based file manager built with Go backend and Vue.js frontend.

## Features

- 📁 Browse files and directories
- 👀 View file contents (text, images, PDFs, videos)
- 🗑️ Delete files and folders
- 📱 Responsive web interface
- 🐳 Docker ready

## Quick Start

```bash
# Clone and run
git clone https://github.com/BomScoob12/homelab-file-manager.git file-manager
cd file-manager

# Windows
dev.bat

# Linux/macOS
./start.sh
```

Open http://localhost:3000

## What's Included

- **Frontend**: Vue.js app on port 3000
- **Backend**: Go API on port 8080
- **Data**: Files stored in `./data` directory

## File Support

| Type | Support |
|------|---------|
| Text files | ✅ View/edit |
| Images | ✅ Display |
| PDFs | ✅ Embedded viewer |
| Videos/Audio | ✅ Media player |
| Other files | ✅ Download |

## Development

```bash
# Backend only
cd backend && go run cmd/app/main.go

# Frontend only  
cd frontend && npm run dev

# Docker
docker compose up --build
```

## Tech Stack

- **Backend**: Go, net/http
- **Frontend**: Vue 3, Tailwind CSS, Vite
- **Deployment**: Docker, Docker Compose