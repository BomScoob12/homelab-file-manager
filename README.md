# File Manager

A full-stack file management system with a Go backend and Vue.js frontend.

## Project Structure

```
.
├── backend/                 # Go API server
│   ├── cmd/app/            # Application entry point
│   ├── internal/           # Internal packages
│   │   ├── files/          # File management logic
│   │   ├── fs/             # File system utilities
│   │   └── routes/         # HTTP routes
│   ├── go.mod              # Go dependencies
│   └── Makefile            # Build commands
├── frontend/               # Vue.js web application
│   ├── src/                # Source code
│   │   ├── components/     # Vue components
│   │   └── services/       # API services
│   ├── package.json        # Node.js dependencies
│   └── vite.config.js      # Build configuration
└── README.md               # This file
```

## Features

### Backend (Go)
- 🚀 **Fast HTTP API** built with Go's net/http
- 📁 **File Operations**: List, read, delete files and directories
- 🔒 **Security**: Path validation and access control
- 🐳 **Docker Ready**: Designed for containerized deployment
- 📊 **Comprehensive Logging**: Request/response logging
- ⚡ **High Performance**: Efficient file system operations

### Frontend (Vue.js)
- 🎨 **Modern UI**: Clean interface with Tailwind CSS
- 📱 **Responsive Design**: Works on desktop and mobile
- 🗂️ **File Browser**: Navigate directories with breadcrumbs
- 👁️ **File Preview**: View text files in modal
- 🗑️ **File Management**: Delete files with confirmation
- 🔄 **Real-time Updates**: Refresh file listings

## Quick Start

### Prerequisites
- **Go 1.21+** for backend
- **Node.js 16+** for frontend

### Development Setup

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd file-manager
   ```

2. **Start the backend** (Terminal 1):
   ```bash
   cd backend
   make setup-test  # Create test files
   make run         # Start Go server on :8080
   ```

3. **Start the frontend** (Terminal 2):
   ```bash
   cd frontend
   npm install      # Install dependencies
   npm run dev      # Start Vite dev server on :3000
   ```

4. **Open your browser** to `http://localhost:3000`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/file/list?path=/` | List files and directories |
| GET | `/file/details?path=/file.txt` | Get file metadata |
| GET | `/file/open?path=/file.txt` | Read file content |
| DELETE | `/file/delete?path=/file.txt` | Delete file or directory |

## Configuration

### Backend Configuration
- **Port**: 8080 (configurable via environment)
- **Base Path**: `/WorkDir` (Docker mount point)
- **File Size Limit**: 10MB for content reading
- **CORS**: Enabled for frontend integration

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