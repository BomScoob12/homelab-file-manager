<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
    <div class="bg-white rounded-lg shadow-xl max-w-md w-full">
      <!-- Header -->
      <div class="p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <ExclamationTriangleIcon class="w-6 h-6 text-red-600" />
          </div>
          <div class="ml-3">
            <h3 class="text-lg font-semibold text-gray-900">
              Delete {{ file.isDir ? 'Directory' : 'File' }}
            </h3>
          </div>
        </div>
      </div>

      <!-- Content -->
      <div class="px-6 pb-6">
        <p class="text-sm text-gray-600 mb-4">
          Are you sure you want to delete 
          <span class="font-medium text-gray-900">{{ file.name }}</span>?
          <span v-if="file.isDir" class="text-red-600">
            This will delete the directory and all its contents.
          </span>
          This action cannot be undone.
        </p>

        <div class="bg-gray-50 rounded-md p-3">
          <div class="flex items-center space-x-2">
            <FolderIcon v-if="file.isDir" class="w-4 h-4 text-blue-500" />
            <DocumentIcon v-else class="w-4 h-4 text-gray-400" />
            <span class="text-sm font-medium text-gray-700">{{ file.path }}</span>
          </div>
          <div class="mt-1 text-xs text-gray-500">
            {{ formatFileSize(file.size) }} • {{ formatDate(file.modTime) }}
          </div>
        </div>
      </div>

      <!-- Footer -->
      <div class="flex justify-end space-x-3 p-6 border-t border-gray-200">
        <button @click="$emit('cancel')" class="btn btn-secondary">
          Cancel
        </button>
        <button @click="$emit('confirm')" class="btn btn-danger">
          <TrashIcon class="w-4 h-4 mr-2" />
          Delete {{ file.isDir ? 'Directory' : 'File' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import {
  ExclamationTriangleIcon,
  FolderIcon,
  DocumentIcon,
  TrashIcon
} from '@heroicons/vue/24/outline'

export default {
  name: 'DeleteConfirmModal',
  components: {
    ExclamationTriangleIcon,
    FolderIcon,
    DocumentIcon,
    TrashIcon
  },
  props: {
    file: {
      type: Object,
      required: true
    }
  },
  emits: ['confirm', 'cancel'],
  setup() {
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

    return {
      formatFileSize,
      formatDate
    }
  }
}
</script>