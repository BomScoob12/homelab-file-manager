import axios from 'axios'

// Use environment variable or default to backend service in Docker
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

// Create axios instance with default config
const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  }
})

// Request interceptor
api.interceptors.request.use(
  (config) => {
    console.log(`Making ${config.method?.toUpperCase()} request to ${config.url}`)
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Response interceptor
api.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    console.error('API Error:', error.response?.data || error.message)
    return Promise.reject(error)
  }
)

export const fileAPI = {
  // List files in a directory
  async listFiles(path = '/') {
    try {
      const response = await api.get('/file/list', {
        params: { path }
      })
      return response.data
    } catch (error) {
      throw new Error(error.response?.data?.error || 'Failed to list files')
    }
  },

  // Get file details
  async getFileDetails(path) {
    try {
      const response = await api.get('/file/details', {
        params: { path }
      })
      return response.data
    } catch (error) {
      throw new Error(error.response?.data?.error || 'Failed to get file details')
    }
  },

  // Open/read file content
  async openFile(path) {
    try {
      const response = await api.get('/file/open', {
        params: { path }
      })
      return response.data
    } catch (error) {
      throw new Error(error.response?.data?.error || 'Failed to open file')
    }
  },

  // Delete file or directory
  async deleteFile(path) {
    try {
      const response = await api.delete('/file/delete', {
        params: { path }
      })
      return response.data
    } catch (error) {
      throw new Error(error.response?.data?.error || 'Failed to delete file')
    }
  }
}

export default api