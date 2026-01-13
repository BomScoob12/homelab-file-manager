<template>
  <div class="space-y-6">
    <!-- Breadcrumb Navigation -->
    <nav class="flex" aria-label="Breadcrumb">
      <ol class="flex items-center space-x-2">
        <li>
          <button
            @click="navigateToPath('/')"
            class="text-primary-600 hover:text-primary-700 font-medium"
          >
            Home
          </button>
        </li>
        <li v-for="(segment, index) in pathSegments" :key="index" class="flex items-center">
          <svg class="w-4 h-4 text-gray-400 mx-2" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 111.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
          </svg>
          <button
            @click="navigateToPath(getPathUpTo(index))"
            class="text-primary-600 hover:text-primary-700 font-medium"
          >
            {{ segment }}
          </button>
        </li>
      </ol>
    </nav>

    <!-- Current Path Display -->
    <div class="card p-4">
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-2">
          <FolderIcon class="w-5 h-5 text-primary-600" />
          <span class="font-medium text-gray-900">{{ currentPath || '/' }}</span>
        </div>
        <button
          @click="refreshFiles"
          :disabled="loading"
          class="btn btn-secondary"
        >
          <ArrowPathIcon class="w-4 h-4 mr-2" :class="{ 'animate-spin': loading }" />
          Refresh
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="card p-8">
      <div class="flex items-center justify-center">
        <ArrowPathIcon class="w-6 h-6 animate-spin text-primary-600 mr-2" />
        <span class="text-gray-600">Loading files...</span>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="card p-6 border-red-200 bg-red-50">
      <div class="flex items-center">
        <ExclamationTriangleIcon class="w-5 h-5 text-red-600 mr-2" />
        <span class="text-red-800">{{ error }}</span>
      </div>
    </div>

    <!-- File List -->
    <div v-else class="card">
      <div class="p-4 border-b border-gray-200">
        <h2 class="text-lg font-semibold text-gray-900">
          Files and Directories
          <span v-if="files.length > 0" class="text-sm font-normal text-gray-500">
            ({{ files.length }} items)
          </span>
        </h2>
      </div>

      <!-- Empty State -->
      <div v-if="files.length === 0" class="p-8 text-center">
        <FolderIcon class="w-12 h-12 text-gray-400 mx-auto mb-4" />
        <p class="text-gray-500">This directory is empty</p>
      </div>

      <!-- File List -->
      <div v-else class="divide-y divide-gray-200">
        <div
          v-for="file in files"
          :key="file.path"
          class="p-4 hover:bg-gray-50 transition-colors duration-150"
        >
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-3 flex-1 min-w-0">
              <!-- File Icon -->
              <div class="flex-shrink-0">
                <FolderIcon v-if="file.isDir" class="w-6 h-6 text-blue-500" />
                <DocumentIcon v-else class="w-6 h-6 text-gray-400" />
              </div>

              <!-- File Info -->
              <div class="flex-1 min-w-0">
                <button
                  @click="handleFileClick(file)"
                  class="text-left w-full"
                >
                  <p class="text-sm font-medium text-gray-900 truncate hover:text-primary-600">
                    {{ file.name }}
                  </p>
                  <p class="text-xs text-gray-500">
                    {{ formatFileSize(file.size) }} • {{ formatDate(file.modTime) }}
                  </p>
                </button>
              </div>
            </div>

            <!-- Actions -->
            <div class="flex items-center space-x-2">
              <button
                v-if="!file.isDir"
                @click="openFile(file)"
                class="p-2 text-gray-400 hover:text-primary-600 rounded-md hover:bg-gray-100"
                title="Open file"
              >
                <EyeIcon class="w-4 h-4" />
              </button>
              <button
                @click="deleteFile(file)"
                class="p-2 text-gray-400 hover:text-red-600 rounded-md hover:bg-gray-100"
                title="Delete"
              >
                <TrashIcon class="w-4 h-4" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- File Content Modal -->
    <FileContentModal
      v-if="showFileContent"
      :file="selectedFile"
      :content="fileContent"
      @close="closeFileContent"
    />

    <!-- Delete Confirmation Modal -->
    <DeleteConfirmModal
      v-if="showDeleteConfirm"
      :file="fileToDelete"
      @confirm="confirmDelete"
      @cancel="cancelDelete"
    />
  </div>
</template>

<script>
import { ref, onMounted, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { fileAPI } from '../services/api'
import {
  FolderIcon,
  DocumentIcon,
  EyeIcon,
  TrashIcon,
  ArrowPathIcon,
  ExclamationTriangleIcon
} from '@heroicons/vue/24/outline'
import FileContentModal from './FileContentModal.vue'
import DeleteConfirmModal from './DeleteConfirmModal.vue'

export default {
  name: 'FileManager',
  components: {
    FolderIcon,
    DocumentIcon,
    EyeIcon,
    TrashIcon,
    ArrowPathIcon,
    ExclamationTriangleIcon,
    FileContentModal,
    DeleteConfirmModal
  },
  setup() {
    const route = useRoute()
    const router = useRouter()

    // Reactive state
    const files = ref([])
    const loading = ref(false)
    const error = ref(null)
    const currentPath = ref('/')
    const showFileContent = ref(false)
    const selectedFile = ref(null)
    const fileContent = ref('')
    const showDeleteConfirm = ref(false)
    const fileToDelete = ref(null)

    // Computed properties
    const pathSegments = computed(() => {
      if (!currentPath.value || currentPath.value === '/') return []
      return currentPath.value.split('/').filter(segment => segment.length > 0)
    })

    // Methods
    const loadFiles = async (path = '/') => {
      loading.value = true
      error.value = null
      
      try {
        const response = await fileAPI.listFiles(path)
        files.value = response.items || []
        currentPath.value = path
      } catch (err) {
        error.value = err.message
        files.value = []
      } finally {
        loading.value = false
      }
    }

    const refreshFiles = () => {
      loadFiles(currentPath.value)
    }

    const navigateToPath = (path) => {
      router.push(path === '/' ? '/' : `/files${path}`)
    }

    const getPathUpTo = (index) => {
      const segments = pathSegments.value.slice(0, index + 1)
      return '/' + segments.join('/')
    }

    const handleFileClick = (file) => {
      if (file.isDir) {
        navigateToPath(file.path)
      } else {
        openFile(file)
      }
    }

    const openFile = async (file) => {
      try {
        loading.value = true
        const response = await fileAPI.openFile(file.path)
        selectedFile.value = file
        fileContent.value = response.content || ''
        showFileContent.value = true
      } catch (err) {
        error.value = `Failed to open file: ${err.message}`
      } finally {
        loading.value = false
      }
    }

    const closeFileContent = () => {
      showFileContent.value = false
      selectedFile.value = null
      fileContent.value = ''
    }

    const deleteFile = (file) => {
      fileToDelete.value = file
      showDeleteConfirm.value = true
    }

    const confirmDelete = async () => {
      try {
        loading.value = true
        await fileAPI.deleteFile(fileToDelete.value.path)
        await loadFiles(currentPath.value) // Refresh the list
        showDeleteConfirm.value = false
        fileToDelete.value = null
      } catch (err) {
        error.value = `Failed to delete: ${err.message}`
      } finally {
        loading.value = false
      }
    }

    const cancelDelete = () => {
      showDeleteConfirm.value = false
      fileToDelete.value = null
    }

    const formatFileSize = (bytes) => {
      if (bytes === 0) return '0 B'
      const k = 1024
      const sizes = ['B', 'KB', 'MB', 'GB']
      const i = Math.floor(Math.log(bytes) / Math.log(k))
      return parseFloat((bytes / Math.pow(k, i)).toFixed(1)) + ' ' + sizes[i]
    }

    const formatDate = (dateString) => {
      return new Date(dateString).toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      })
    }

    // Watch route changes
    watch(
      () => route.params.path,
      (newPath) => {
        const path = newPath ? '/' + (Array.isArray(newPath) ? newPath.join('/') : newPath) : '/'
        loadFiles(path)
      },
      { immediate: true }
    )

    // Load files on mount
    onMounted(() => {
      const path = route.params.path 
        ? '/' + (Array.isArray(route.params.path) ? route.params.path.join('/') : route.params.path)
        : '/'
      loadFiles(path)
    })

    return {
      files,
      loading,
      error,
      currentPath,
      pathSegments,
      showFileContent,
      selectedFile,
      fileContent,
      showDeleteConfirm,
      fileToDelete,
      loadFiles,
      refreshFiles,
      navigateToPath,
      getPathUpTo,
      handleFileClick,
      openFile,
      closeFileContent,
      deleteFile,
      confirmDelete,
      cancelDelete,
      formatFileSize,
      formatDate
    }
  }
}
</script>