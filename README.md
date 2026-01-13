# File Manager

A full-stack file management system with a Go backend and Vue.js frontend, designed for Docker deployment.

## Project Structure

```
.
├── backend/                 # Go API server
│   ├── cmd/app/            # Application entry point
│   ├── internal/           # Internal packages
│   ├── Dockerfile          # Backend Docker image
│   └── .env                # Environment configuration
├── frontend/               # Vue.js web application
│   ├── src/                # Source code
│   └── Dockerfile          # Frontend Docker image
├── docker-compose.yml      # Production Docker Compose
├── docker-compose.dev.yml  # Development Docker Compose
├── start.sh               # Unix startup script
├── dev.bat                # Windows startup script
└── data/                  # Mounted data directory
```

## Features

### Backend (Go)
- 🚀 **Fast HTTP API** built with Go's net/http
- 📁 **File Operations**: List, read, delete files and directories
- 🔒 **Security**: Path validation and access control
- 🐳 **Docker Native**: Designed for containerized deployment
- 📊 **Health Checks**: Built-in health monitoring
- ⚡ **High Performance**: Efficient file system operations

### Frontend (Vue.js)
- 🎨 **Modern UI**: Clean interface with Tailwind CSS
- 📱 **Responsive Design**: Works on desktop and mobile
- 🗂️ **File Browser**: Navigate directories with breadcrumbs
- 👁️ **File Preview**: View text files in modal
- 🗑️ **File Management**: Delete files with confirmation
- 🔄 **Real-time Updates**: Refresh file listings

## Quick Start with Docker

### Prerequisites
- **Docker** and **Docker Compose**

### Easy Setup

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd file-manager
   ```

2. **Start with Docker Compose** (Recommended):
   ```bash
   # Unix/Linux/macOS
   ./start.sh
   
   # Windows
   dev.bat
   
   # Or manually
   docker-compose up --build
   ```

3. **Access the application**:
   - **Frontend**: http://localhost:3000
   - **Backend API**: http://localhost:8080
   - **Data Directory**: `./data` (auto-created)

### Development Mode

For development with hot reload:
```bash
docker-compose -f docker-compose.dev.yml up --build
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/file/list?path=/` | List files and directories |
| GET | `/file/details?path=/file.txt` | Get file metadata |
| GET | `/file/open?path=/file.txt` | Read file content |
| DELETE | `/file/delete?path=/file.txt` | Delete file or directory |

## Configuration

### Backend Configuration (.env file)
Create a `.env` file in the `backend/` directory:

```bash
# Copy the example file
cp backend/.env.example backend/.env
```

Available environment variables:
- **FILE_MANAGER_BASE_PATH**: Base directory for file operations (default: `C:/temp/test-localfile-manager` on Windows)
- **PORT**: Server port (default: 8080)
- **HOST**: Server host (default: empty - all interfaces)
- **LOG_LEVEL**: Logging level (default: info)

### Frontend Configuration
- **Dev Server**: Port 3000
- **API Proxy**: `/api/*` → `http://localhost:8080`
- **Build Output**: `frontend/dist/`

## Docker Deployment

### Backend Only
```bash
cd backend
docker build -t file-manager-backend .
docker run -p 8080:8080 -v /host/files:/WorkDir file-manager-backend
```

### Full Stack (Docker Compose)
```yaml
version: '3.8'
services:
  backend:
    build: ./backend
    ports:
      - "8080:8080"
    volumes:
      - /host/files:/WorkDir
  
  frontend:
    build: ./frontend
    ports:
      - "3000:3000"
    depends_on:
      - backend
```

## Development Commands

### Backend Commands
```bash
cd backend
make run          # Start development server
make test         # Run tests
make build        # Build binary
make setup-test   # Create test environment
make test-client  # Test API endpoints
```

### Frontend Commands
```bash
cd frontend
npm run dev       # Start development server
npm run build     # Build for production
npm run preview   # Preview production build
```

## File System Structure

The application operates on a base directory (`/WorkDir` in Docker):

```
/WorkDir/                    # Base directory (Docker mount)
├── documents/               # Example directory
│   ├── report.pdf          # Files
│   └── notes.txt
├── images/
│   └── photo.jpg
└── config.json
```

## Security Features

- **Path Validation**: Prevents directory traversal attacks
- **Base Directory Restriction**: All operations within `/WorkDir`
- **File Size Limits**: Prevents memory exhaustion
- **Input Sanitization**: Clean and validate all inputs
- **CORS Configuration**: Controlled cross-origin access

## API Response Format

### Success Response
```json
{
  "success": true,
  "path": "/documents",
  "items": [
    {
      "name": "file.txt",
      "path": "/documents/file.txt",
      "isDir": false,
      "size": 1024,
      "modTime": "2024-01-15T10:30:00Z",
      "permissions": "-rw-r--r--",
      "mimeType": "text/plain"
    }
  ],
  "totalItems": 1,
  "totalSize": 1024,
  "requestTime": "2024-01-15T12:00:00Z"
}
```

### Error Response
```json
{
  "success": false,
  "error": "File not found",
  "code": 404
}
```

## Testing

### Backend Testing
```bash
cd backend
make test                    # Run all tests
make test-client            # Test API with client
curl "http://localhost:8080/file/list?path=/"  # Manual test
```

### Frontend Testing
- Open browser developer tools
- Check network tab for API calls
- Test file operations in the UI

## Troubleshooting

### Common Issues

1. **Backend not starting**:
   - Check if port 8080 is available
   - Verify Go installation and version
   - Check file permissions for `/WorkDir`

2. **Frontend can't connect to backend**:
   - Ensure backend is running on port 8080
   - Check CORS configuration
   - Verify proxy settings in `vite.config.js`

3. **File operations failing**:
   - Check file permissions
   - Verify path exists and is accessible
   - Check backend logs for detailed errors

### Debug Mode

Enable debug logging:
```bash
# Backend
cd backend
LOG_LEVEL=debug make run

# Frontend
cd frontend
DEBUG=true npm run dev
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.

## Roadmap

- [ ] File upload functionality
- [ ] File editing capabilities
- [ ] Search and filtering
- [ ] User authentication
- [ ] File sharing and permissions
- [ ] Bulk operations
- [ ] File versioning
- [ ] Integration with cloud storage
##
 Configuration

### Environment Variables (.env)
The backend uses environment variables for configuration:

```bash
# Base path for file operations (Docker mount point)
FILE_MANAGER_BASE_PATH=/data

# Server Configuration
PORT=8080
HOST=0.0.0.0

# Development Settings
LOG_LEVEL=info
```

### Docker Volumes
- **Data Directory**: `./data` is mounted to `/data` in the container
- **All file operations** happen within this mounted directory
- **Persistent storage** across container restarts

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/file/list?path=/` | List files and directories |
| GET | `/file/details?path=/file.txt` | Get file metadata |
| GET | `/file/open?path=/file.txt` | Read file content |
| DELETE | `/file/delete?path=/file.txt` | Delete file or directory |

## Docker Deployment

### Production Deployment
```bash
# Build and start services
docker-compose up -d --build

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Custom Data Directory
```bash
# Use custom data directory
docker run -v /your/data/path:/data -p 3000:3000 file-manager
```

### Environment Customization
Create a `.env` file in the project root:
```bash
# Custom configuration
FILE_MANAGER_BASE_PATH=/data
PORT=8080
LOG_LEVEL=debug
```

## Development

### Local Development (without Docker)
```bash
# Backend
cd backend
make setup-data  # Create sample data
make run         # Start Go server

# Frontend  
cd frontend
npm install
npm run dev      # Start Vite dev server
```

### Docker Development
```bash
# Development mode with hot reload
docker-compose -f docker-compose.dev.yml up --build
```

## File System Structure

The application operates on a data directory mounted at `/data`:

```
/data/                      # Base directory (Docker mount)
├── documents/              # Example directory
│   ├── report.pdf         # Files
│   └── notes.txt
├── images/
│   └── photo.jpg
└── config.json
```

## Security Features

- **Path Validation**: Prevents directory traversal attacks
- **Base Directory Restriction**: All operations within `/data`
- **Input Sanitization**: Clean and validate all inputs
- **CORS Configuration**: Controlled cross-origin access
- **Non-root User**: Containers run as non-root user

## Troubleshooting

### Common Issues

1. **Port already in use**:
   ```bash
   # Check what's using the port
   docker ps
   # Stop conflicting containers
   docker-compose down
   ```

2. **Permission issues with data directory**:
   ```bash
   # Fix permissions
   sudo chown -R $USER:$USER ./data
   ```

3. **Container build fails**:
   ```bash
   # Clean Docker cache
   docker system prune -a
   # Rebuild without cache
   docker-compose build --no-cache
   ```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test with Docker
5. Submit a pull request

## License

This project is licensed under the MIT License.