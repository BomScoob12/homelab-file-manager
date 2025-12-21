# File Management API Documentation

## Overview
A comprehensive file management REST API built with Go's net/http package. The service provides secure file operations within a containerized environment.

## Base URL
```
http://localhost:8080/file
```

## Architecture Flow
```
User Request → Handler Layer → Service Layer → File System Layer → Response
```

### Layer Responsibilities:
- **Handler Layer**: HTTP request/response handling, validation, error mapping
- **Service Layer**: Business logic, path validation, security checks
- **File System Layer**: Direct file operations, abstracted through interface

## Endpoints

### 1. List Files and Directories
**Endpoint**: `GET /file/list`

Lists all files and directories in the specified path.

**Query Parameters**:
- `path` (optional): Directory path to list. Defaults to "/"

**Example Requests**:
```bash
# List root directory
curl "http://localhost:8080/file/list?path=/"

# List subdirectory
curl "http://localhost:8080/file/list?path=/documents"

# List with no path (defaults to root)
curl "http://localhost:8080/file/list"
```

**Success Response** (200 OK):
```json
{
  "success": true,
  "path": "/documents",
  "items": [
    {
      "name": "report.pdf",
      "path": "/documents/report.pdf",
      "isDir": false,
      "fileType": "regular",
      "size": 2048576,
      "modTime": "2024-01-15T10:30:00Z",
      "permissions": "-rw-r--r--",
      "extension": ".pdf"
    },
    {
      "name": "images",
      "path": "/documents/images",
      "isDir": true,
      "fileType": "directory",
      "size": 0,
      "modTime": "2024-01-14T15:20:00Z",
      "permissions": "drwxr-xr-x",
      "extension": ""
    }
  ],
  "totalItems": 2,
  "totalSize": 2048576,
  "requestTime": "2024-01-15T12:00:00Z"
}
```

### 2. Open File Content
**Endpoint**: `GET /file/open`

Opens and reads the content of a file. Supports text files up to 10MB.

**Query Parameters**:
- `path` (required): File path to open

**Example Requests**:
```bash
# Open text file
curl "http://localhost:8080/file/open?path=/documents/readme.txt"

# Open configuration file
curl "http://localhost:8080/file/open?path=/config/app.json"
```

**Success Response** (200 OK):
```json
{
  "success": true,
  "name": "readme.txt",
  "path": "/documents/readme.txt",
  "content": "This is the content of the file...",
  "size": 1024,
  "mimeType": "text/plain",
  "encoding": "utf-8",
  "requestTime": "2024-01-15T12:00:00Z"
}
```

### 3. Get File Details
**Endpoint**: `GET /file/details`

Retrieves detailed metadata about a file or directory without reading content.

**Query Parameters**:
- `path` (required): File or directory path

**Example Requests**:
```bash
# Get file details
curl "http://localhost:8080/file/details?path=/documents/report.pdf"

# Get directory details
curl "http://localhost:8080/file/details?path=/documents"
```

**Success Response** (200 OK):
```json
{
  "success": true,
  "name": "report.pdf",
  "path": "/documents/report.pdf",
  "fullPath": "/WorkDir/documents/report.pdf",
  "isDir": false,
  "size": 2048576,
  "modTime": "2024-01-15T10:30:00Z",
  "mimeType": "application/pdf",
  "permissions": "-rw-r--r--",
  "extension": ".pdf",
  "requestTime": "2024-01-15T12:00:00Z"
}
```

### 4. Delete File or Directory
**Endpoint**: `DELETE /file/delete`

Deletes a specified file or directory (including all contents for directories).

**Query Parameters**:
- `path` (required): File or directory path to delete

**Example Requests**:
```bash
# Delete file
curl -X DELETE "http://localhost:8080/file/delete?path=/documents/old_file.txt"

# Delete directory
curl -X DELETE "http://localhost:8080/file/delete?path=/temp"
```

**Success Response** (200 OK):
```json
{
  "success": true,
  "message": "File deleted successfully",
  "path": "/documents/old_file.txt"
}
```

## Error Responses

All error responses follow this format:
```json
{
  "success": false,
  "error": "Error message description",
  "code": 400
}
```

### Common Error Codes:
- **400 Bad Request**: Missing required parameters, invalid path
- **403 Forbidden**: Access denied, path outside allowed directory
- **404 Not Found**: File or directory not found
- **500 Internal Server Error**: Server-side errors

### Example Error Responses:

**Missing Path Parameter** (400):
```json
{
  "success": false,
  "error": "File path is required",
  "code": 400
}
```

**File Not Found** (404):
```json
{
  "success": false,
  "error": "File or directory not found",
  "code": 404
}
```

**Access Denied** (403):
```json
{
  "success": false,
  "error": "Access denied",
  "code": 403
}
```

## Security Features

### Path Validation
- All paths are cleaned using `filepath.Clean()`
- Path traversal attacks prevented (no `../`, `..\\`, etc.)
- All operations restricted to `/WorkDir` base directory
- Relative paths converted to absolute within base directory

### File Size Limits
- File content reading limited to 10MB
- Prevents memory exhaustion attacks
- Binary files handled appropriately

### CORS Support
- Cross-origin requests supported
- Proper preflight handling for web applications

## Docker Integration

The service is designed for containerized deployment:

```dockerfile
# Mount host directory to container
VOLUME ["/WorkDir"]

# All file operations happen within this mounted directory
# Host path: /host/files -> Container path: /WorkDir
```

### Environment Setup:
```bash
# Run with Docker
docker run -v /host/files:/WorkDir -p 8080:8080 file-manager

# All API operations will work on files in /host/files
```

## MIME Type Support

The service recognizes 50+ file types including:

### Text Files
- `.txt`, `.md`, `.csv`, `.log` → `text/plain`
- `.json` → `application/json`
- `.xml` → `application/xml`
- `.yaml`, `.yml` → `application/x-yaml`

### Code Files
- `.go` → `text/x-go`
- `.js` → `application/javascript`
- `.py` → `text/x-python`
- `.java` → `text/x-java-source`

### Web Files
- `.html` → `text/html`
- `.css` → `text/css`
- `.scss` → `text/x-scss`

### Images
- `.png` → `image/png`
- `.jpg`, `.jpeg` → `image/jpeg`
- `.svg` → `image/svg+xml`

### Documents
- `.pdf` → `application/pdf`
- `.docx` → `application/vnd.openxmlformats-officedocument.wordprocessingml.document`

## Best Practices Implemented

### Error Handling
- Structured error responses
- Appropriate HTTP status codes
- Detailed error logging
- Graceful degradation

### Performance
- Efficient file operations
- Memory-conscious content reading
- Proper resource cleanup
- Request/response logging

### Security
- Input validation and sanitization
- Path traversal prevention
- Access control enforcement
- Safe file operations

### Code Organization
- Interface-based design
- Separation of concerns
- Dependency injection
- Testable architecture

## Testing Examples

### Using curl:
```bash
# Test file listing
curl -v "http://localhost:8080/file/list?path=/"

# Test file opening
curl -v "http://localhost:8080/file/open?path=/test.txt"

# Test file details
curl -v "http://localhost:8080/file/details?path=/test.txt"

# Test file deletion
curl -v -X DELETE "http://localhost:8080/file/delete?path=/test.txt"
```

### Using JavaScript (fetch):
```javascript
// List files
const listFiles = async (path = '/') => {
  const response = await fetch(`http://localhost:8080/file/list?path=${encodeURIComponent(path)}`);
  return await response.json();
};

// Open file
const openFile = async (path) => {
  const response = await fetch(`http://localhost:8080/file/open?path=${encodeURIComponent(path)}`);
  return await response.json();
};

// Delete file
const deleteFile = async (path) => {
  const response = await fetch(`http://localhost:8080/file/delete?path=${encodeURIComponent(path)}`, {
    method: 'DELETE'
  });
  return await response.json();
};
```

## Logging

The service provides comprehensive logging:
- Request start/completion times
- Error details with context
- File operation results
- Performance metrics

Example log output:
```
2024/01/15 12:00:00 Started GET /file/list
2024/01/15 12:00:00 Completed GET /file/list in 15.2ms
2024/01/15 12:00:05 Error opening file /nonexistent.txt: failed to get file info: stat /WorkDir/nonexistent.txt: no such file or directory
```