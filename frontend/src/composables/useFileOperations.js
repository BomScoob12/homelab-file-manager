import { fileAPI } from '../services/api'

export function useFileOperations() {
  const listFiles = async (path) => {
    return await fileAPI.listFiles(path)
  }

  const openFileContent = async (path) => {
    return await fileAPI.openFile(path)
  }

  const deleteFileItem = async (path) => {
    return await fileAPI.deleteFile(path)
  }

  const downloadFileItem = async (file) => {
    const response = await fileAPI.openFile(file.path)
    const blob = new Blob([response.content], { type: file.mimeType || 'text/plain' })
    const url = window.URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = file.name
    document.body.appendChild(a)
    a.click()
    window.URL.revokeObjectURL(url)
    document.body.removeChild(a)
  }

  return {
    listFiles,
    openFileContent,
    deleteFileItem,
    downloadFileItem
  }
}