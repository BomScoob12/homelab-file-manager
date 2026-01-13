<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
    <div class="bg-white rounded-lg shadow-xl max-w-4xl w-full max-h-[90vh] flex flex-col">
      <!-- Header -->
      <div class="flex items-center justify-between p-6 border-b border-gray-200">
        <div class="flex items-center space-x-3">
          <DocumentIcon class="w-6 h-6 text-gray-400" />
          <div>
            <h3 class="text-lg font-semibold text-gray-900">{{ file.name }}</h3>
            <p class="text-sm text-gray-500">{{ file.path }}</p>
          </div>
        </div>
        <button
          @click="$emit('close')"
          class="p-2 text-gray-400 hover:text-gray-600 rounded-md hover:bg-gray-100"
        >
          <XMarkIcon class="w-5 h-5" />
        </button>
      </div>

      <!-- File Info -->
      <div class="px-6 py-4 bg-gray-50 border-b border-gray-200">
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4 text-sm">
          <div>
            <span class="text-gray-500">Size:</span>
            <span class="ml-2 font-medium">{{ formatFileSize(file.size) }}</span>
          </div>
          <div>
            <span class="text-gray-500">Modified:</span>
            <span class="ml-2 font-medium">{{ formatDate(file.modTime) }}</span>
          </div>
          <div>
            <span class="text-gray-500">Type:</span>
            <span class="ml-2 font-medium">{{ file.mimeType || 'Unknown' }}</span>
          </div>
          <div>
            <span class="text-gray-500">Permissions:</span>
            <span class="ml-2 font-medium font-mono">{{ file.permissions }}</span>
          </div>
        </div>
      </div>

      <!-- Content -->
      <div class="flex-1 overflow-hidden">
        <div v-if="isTextFile" class="h-full">
          <div class="p-4 border-b border-gray-200 bg-gray-50">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">File Content</span>
              <button
                @click="copyToClipboard"
                class="btn btn-secondary text-xs"
              >
                <ClipboardIcon class="w-4 h-4 mr-1" />
                Copy
              </button>
            </div>
          </div>
          <div class="p-6 overflow-auto h-full max-h-96">
            <pre class="text-sm text-gray-800 whitespace-pre-wrap font-mono leading-relaxed bg-gray-50 p-4 rounded border">{{ content }}</pre>
          </div>
        </div>

        <div v-else-if="isImageFile" class="p-6 flex items-center justify-center h-full">
          <div class="text-center">
            <PhotoIcon class="w-16 h-16 text-gray-400 mx-auto mb-4" />
            <p class="text-gray-600 mb-4">Image preview not available</p>
            <p class="text-sm text-gray-500">{{ file.name }}</p>
          </div>
        </div>

        <div v-else class="p-6 flex items-center justify-center h-full">
          <div class="text-center">
            <DocumentIcon class="w-16 h-16 text-gray-400 mx-auto mb-4" />
            <p class="text-gray-600 mb-2">Binary file cannot be displayed</p>
            <p class="text-sm text-gray-500">{{ file.name }} ({{ formatFileSize(file.size) }})</p>
          </div>
        </div>
      </div>

      <!-- Footer -->
      <div class="flex justify-end space-x-3 p-6 border-t border-gray-200">
        <button @click="$emit('close')" class="btn btn-secondary">
          Close
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useToast } from '../composables/useToast'
import {
  DocumentIcon,
  XMarkIcon,
  ClipboardIcon,
  PhotoIcon
} from '@heroicons/vue/24/outline'

export default {
  name: 'FileContentModal',
  components: {
    DocumentIcon,
    XMarkIcon,
    ClipboardIcon,
    PhotoIcon
  },
  props: {
    file: {
      type: Object,
      required: true
    },
    content: {
      type: String,
      default: ''
    }
  },
  emits: ['close'],
  setup(props) {
    const toast = useToast()
    const isTextFile = computed(() => {
      const textTypes = [
        'text/',
        'application/json',
        'application/xml',
        'application/javascript',
        'application/typescript',
        'text/x-go',
        'text/x-python',
        'text/x-java-source',
        'text/x-c',
        'text/x-c++',
        'text/markdown',
        'application/x-yaml',
        'application/toml'
      ]
      return textTypes.some(type => props.file.mimeType?.startsWith(type))
    })

    const isImageFile = computed(() => {
      return props.file.mimeType?.startsWith('image/')
    })

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

    const copyToClipboard = async () => {
      try {
        await navigator.clipboard.writeText(props.content)
        toast.success('Copied to clipboard', 'File content has been copied to clipboard')
      } catch (err) {
        console.error('Failed to copy content:', err)
        toast.error('Copy failed', 'Failed to copy content to clipboard')
      }
    }

    return {
      isTextFile,
      isImageFile,
      formatFileSize,
      formatDate,
      copyToClipboard
    }
  }
}
</script>