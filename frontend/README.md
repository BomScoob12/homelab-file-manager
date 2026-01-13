# File Manager Frontend

A modern Vue.js frontend for the file management system.

## Features

- 📁 **Directory Navigation**: Browse through directories with breadcrumb navigation
- 📄 **File Preview**: View text file contents in a modal
- 🗑️ **File Operations**: Delete files and directories with confirmation
- 🔄 **Real-time Updates**: Refresh file listings
- 📱 **Responsive Design**: Works on desktop and mobile devices
- 🎨 **Modern UI**: Clean interface built with Tailwind CSS

## Tech Stack

- **Vue 3** - Progressive JavaScript framework
- **Vue Router** - Client-side routing
- **Vite** - Fast build tool and dev server
- **Tailwind CSS** - Utility-first CSS framework
- **Heroicons** - Beautiful SVG icons
- **Axios** - HTTP client for API calls

## Project Structure

```
frontend/
├── src/
│   ├── components/          # Vue components
│   │   ├── FileManager.vue      # Main file browser
│   │   ├── FileContentModal.vue # File content viewer
│   │   └── DeleteConfirmModal.vue # Delete confirmation
│   ├── services/           # API services
│   │   └── api.js          # Backend API calls
│   ├── App.vue             # Root component
│   ├── main.js             # Application entry point
│   └── style.css           # Global styles
├── public/                 # Static assets
├── index.html              # HTML template
├── package.json            # Dependencies and scripts
├── vite.config.js          # Vite configuration
└── tailwind.config.js      # Tailwind configuration
```

## Installation

1. **Install Node.js** (version 16 or higher)

2. **Install dependencies**:
   ```bash
   cd frontend
   npm install
   ```

## Development

1. **Start the development server**:
   ```bash
   npm run dev
   ```

2. **Open your browser** to `http://localhost:3000`

3. **Make sure the backend is running** on `http://localhost:8080`

## Build for Production

1. **Build the application**:
   ```bash
   npm run build
   ```

2. **Preview the production build**:
   ```bash
   npm run preview
   ```

The built files will be in the `dist/` directory.

## API Integration

The frontend communicates with the backend through these endpoints:

- `GET /file/list?path=/` - List files and directories
- `GET /file/details?path=/file.txt` - Get file metadata
- `GET /file/open?path=/file.txt` - Read file content
- `DELETE /file/delete?path=/file.txt` - Delete file or directory

## Configuration

### Vite Configuration (`vite.config.js`)

- **Dev Server**: Runs on port 3000
- **Proxy**: API calls to `/api/*` are proxied to `http://localhost:8080`
- **Build Output**: Files are built to `dist/` directory

### API Configuration (`src/services/api.js`)

- **Base URL**: `http://localhost:8080`
- **Timeout**: 10 seconds
- **Error Handling**: Automatic error logging and user-friendly messages

## Components

### FileManager.vue
Main component that handles:
- Directory navigation with breadcrumbs
- File listing with icons and metadata
- File operations (open, delete)
- Loading and error states

### FileContentModal.vue
Modal component for displaying:
- Text file contents with syntax highlighting
- File metadata (size, date, permissions)
- Copy to clipboard functionality
- Binary file indicators

### DeleteConfirmModal.vue
Confirmation dialog for:
- File and directory deletion
- Warning messages for directories
- File information display

## Styling

The application uses Tailwind CSS with custom components:

- **Buttons**: `.btn`, `.btn-primary`, `.btn-secondary`, `.btn-danger`
- **Cards**: `.card` for content containers
- **Inputs**: `.input` for form elements

## Browser Support

- Chrome/Edge 88+
- Firefox 85+
- Safari 14+

## Development Tips

1. **Hot Reload**: Changes are automatically reflected in the browser
2. **Vue DevTools**: Install the browser extension for debugging
3. **API Debugging**: Check the browser console for API call logs
4. **Responsive Testing**: Use browser dev tools to test mobile layouts

## Troubleshooting

### Common Issues

1. **API Connection Failed**
   - Ensure backend is running on port 8080
   - Check CORS settings in backend
   - Verify API endpoints are accessible

2. **Build Errors**
   - Clear node_modules and reinstall: `rm -rf node_modules && npm install`
   - Check Node.js version compatibility

3. **Styling Issues**
   - Ensure Tailwind CSS is properly configured
   - Check for conflicting CSS rules

### Error Messages

- **"Failed to list files"**: Backend API is not responding
- **"File not found"**: Requested file doesn't exist
- **"Access denied"**: File permissions or path security issue

## Future Enhancements

- 📤 **File Upload**: Drag and drop file uploads
- ✏️ **File Editing**: In-browser text file editing
- 🔍 **Search**: File and content search functionality
- 📊 **File Analytics**: Storage usage and file type statistics
- 🎨 **Themes**: Dark mode and custom themes
- 📱 **Mobile App**: Progressive Web App (PWA) features