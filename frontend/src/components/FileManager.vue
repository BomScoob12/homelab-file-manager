<template>
  <div class="space-y-6">
    <!-- Breadcrumb Navigation -->
    <BreadcrumbNavigation
      :current-path="currentPath"
      @navigate="navigateToPath"
    />

    <!-- Current Path Display -->
    <BaseCard>
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-2">
          <FolderIcon class="w-5 h-5 text-primary-600" />
          <span class="font-medium text-gray-900">{{ currentPath || '/' }}</span>
          <span v-if="files.length > 0" class="text-sm text-gray-500 ml-2">
            ({{ files.length }} {{ files.length === 1 ? 'item' : 'items' }})
          </span>
        </div>
        <div class="flex items-center space-x-2">
          <!-- Back Button -->
          <BaseButton
            v-if="currentPath !== '/'"
            variant="secondary"
            :icon="ArrowLeftIcon"
            @click="navigateToParent"
          >
            Back
          </BaseButton>
          <!-- Refresh Button -->
          <BaseButton
            variant="secondary"
            :icon="ArrowPathIcon"
            :loading="loading"
            @click="refreshFiles"
          >
            Refresh
          </BaseButton>
        </div>
      </div>
    </BaseCard>

    <!-- Loading State -->
    <BaseCard v-if="loading">
      <LoadingSpinner size="md" text="Loading files..." center />
    </BaseCard>

    <!-- Error State -->
    <BaseCard v-else-if="error" class="border-red-200 bg-red-50">
      <div class="flex items-center">
        <ExclamationTriangleIcon class="w-5 h-5 text-red-600 mr-2" />
        <span class="text-red-800">{{ error }}</span>
      </div>
    </BaseCard>

    <!-- File List -->
    <BaseCard v-else padding="none">
      <template #header>
        <h2 class="text-lg font-semibold text-gray-900">
          Files and Directories
          <span v-if="files.length > 0" class="text-sm font-normal text-gray-500">
            ({{ files.length }} items)
          </span>
        </h2>
      </template>

      <!-- Empty State -->
      <div v-if="files.length === 0" class="p-8 text-center">
        <FolderIcon class="w-12 h-12 text-gray-400 mx-auto mb-4" />
        <p class="text-gray-500">This directory is empty</p>
      </div>

      <!-- File List -->
      <div v-else class="divide-y divide-gray-200">
        <FileListItem
          v-for="file in files"
          :key="file.path"
          :file="file"
          @click="handleFileClick"
          @open="openFile"
          @download="downloadFile"
          @delete="deleteFile"
        />
      </div>
    </BaseCard>

    <!-- File Content Modal -->
    <FileContentModal
      v-model="showFileContent"
      :file="selectedFile"
      :content="fileContent"
    />

    <!-- Delete Confirmation Modal -->
    <DeleteConfirmModal
      v-model="showDeleteConfirm"
      :file="fileToDelete"
      @confirm="confirmDelete"
    />
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  FolderIcon,
  ArrowPathIcon,
  ArrowLeftIcon,
  ExclamationTriangleIcon
} from '@heroicons/vue/24/outline'

// Components
import BaseCard from './common/BaseCard.vue'
import BaseButton from './common/BaseButton.vue'
import LoadingSpinner from './common/LoadingSpinner.vue'
import BreadcrumbNavigation from './BreadcrumbNavigation.vue'
import FileListItem from './FileListItem.vue'
import FileContentModal from './FileContentModal.vue'
import DeleteConfirmModal from './DeleteConfirmModal.vue'

// Composables
import { useFileOperations } from '../composables/useFileOperations'
import { useNavigation } from '../composables/useNavigation'
import { useToast } from '../composables/useToast'

// Setup
const route = useRoute()
const router = useRouter()
const toast = useToast()

// State
const files = ref([])
const loading = ref(false)
const error = ref(null)
const showFileContent = ref(false)
const selectedFile = ref(null)
const fileContent = ref('')
const showDeleteConfirm = ref(false)
const fileToDelete = ref(null)

// Composables
const { currentPath, navigateToPath, navigateToParent } = useNavigation(route, router)
const { listFiles, openFileContent, deleteFileItem, downloadFileItem } = useFileOperations()

// Methods
const loadFiles = async (path = '/') => {
  loading.value = true
  error.value = null
  
  try {
    const response = await listFiles(path)
    files.value = response.items || []
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
    error.value = null
    const response = await openFileContent(file.path)
    selectedFile.value = file
    fileContent.value = response.content || ''
    showFileContent.value = true
  } catch (err) {
    error.value = `Failed to open file: ${err.message}`
    console.error('Error opening file:', err)
  } finally {
    loading.value = false
  }
}

const downloadFile = async (file) => {
  try {
    await downloadFileItem(file)
    toast.success('File downloaded', `${file.name} has been downloaded successfully`)
  } catch (err) {
    error.value = `Failed to download file: ${err.message}`
    toast.error('Download failed', err.message)
    console.error('Error downloading file:', err)
  }
}

const deleteFile = (file) => {
  fileToDelete.value = file
  showDeleteConfirm.value = true
}

const confirmDelete = async () => {
  try {
    loading.value = true
    await deleteFileItem(fileToDelete.value.path)
    await loadFiles(currentPath.value)
    showDeleteConfirm.value = false
    toast.success('File deleted', `${fileToDelete.value.name} has been deleted successfully`)
    fileToDelete.value = null
  } catch (err) {
    error.value = `Failed to delete: ${err.message}`
    toast.error('Delete failed', err.message)
  } finally {
    loading.value = false
  }
}

// Watchers
watch(
  () => route.params.path,
  (newPath) => {
    const path = newPath ? '/' + (Array.isArray(newPath) ? newPath.join('/') : newPath) : '/'
    currentPath.value = path
    loadFiles(path)
  },
  { immediate: true }
)

// Lifecycle
onMounted(() => {
  const path = route.params.path 
    ? '/' + (Array.isArray(route.params.path) ? route.params.path.join('/') : route.params.path)
    : '/'
  loadFiles(path)
})
</script>