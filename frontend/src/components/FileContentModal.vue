<template>
  <BaseModal
    :model-value="modelValue"
    title="File Content"
    size="lg"
    @update:model-value="handleClose"
    @close="handleClose"
  >
    <template #header>
      <div class="flex items-center space-x-3">
        <DocumentIcon class="w-6 h-6 text-gray-400" />
        <div>
          <h3 class="text-lg font-semibold text-gray-900">{{ file?.name }}</h3>
          <p class="text-sm text-gray-500">{{ file?.path }}</p>
        </div>
      </div>
    </template>

    <!-- File Info -->
    <div class="px-6 py-4 bg-gray-50 border-b border-gray-200">
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4 text-sm">
        <div>
          <span class="text-gray-500">Size:</span>
          <span class="ml-2 font-medium">{{ formatFileSize(file?.size || 0) }}</span>
        </div>
        <div>
          <span class="text-gray-500">Modified:</span>
          <span class="ml-2 font-medium">{{ formatDate(file?.modTime) }}</span>
        </div>
        <div>
          <span class="text-gray-500">Type:</span>
          <span class="ml-2 font-medium">{{ file?.mimeType || 'Unknown' }}</span>
        </div>
        <div>
          <span class="text-gray-500">Permissions:</span>
          <span class="ml-2 font-medium font-mono">{{ file?.permissions }}</span>
        </div>
      </div>
    </div>

    <!-- Content -->
    <div class="flex-1 overflow-hidden">
      <!-- Text Files -->
      <div v-if="isTextFile" class="h-full">
        <div class="p-4 border-b border-gray-200 bg-gray-50">
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-600">File Content</span>
            <BaseButton
              variant="secondary"
              size="sm"
              :icon="ClipboardIcon"
              @click="copyToClipboard"
            >
              Copy
            </BaseButton>
          </div>
        </div>
        <div class="p-6 overflow-auto h-full max-h-96">
          <pre class="text-sm text-gray-800 whitespace-pre-wrap font-mono leading-relaxed bg-gray-50 p-4 rounded border">{{ content }}</pre>
        </div>
      </div>

      <!-- Images -->
      <div v-else-if="isImageFile" class="p-6 flex items-center justify-center h-full">
        <div class="text-center max-w-full">
          <img 
            :src="getFileUrl(file?.path)" 
            :alt="file?.name"
            class="max-w-full max-h-96 object-contain rounded shadow-lg"
            @error="handleImageError"
            @load="handleImageLoad"
          />
          <p class="text-sm text-gray-500 mt-4">{{ file?.name }}</p>
        </div>
      </div>

      <!-- PDFs -->
      <div v-else-if="isPdfFile" class="h-full">
        <iframe 
          :src="getFileUrl(file?.path)"
          class="w-full h-96 border-0"
          title="PDF Viewer"
        />
        <div class="p-4 text-center border-t">
          <p class="text-sm text-gray-500">{{ file?.name }}</p>
          <BaseButton
            variant="primary"
            size="sm"
            class="mt-2"
            @click="downloadFile"
          >
            Download PDF
          </BaseButton>
        </div>
      </div>

      <!-- Videos -->
      <div v-else-if="isVideoFile" class="p-6 flex items-center justify-center h-full">
        <div class="text-center max-w-full">
          <video 
            :src="getFileUrl(file?.path)"
            controls
            class="max-w-full max-h-96 rounded shadow-lg"
          >
            Your browser does not support the video tag.
          </video>
          <p class="text-sm text-gray-500 mt-4">{{ file?.name }}</p>
        </div>
      </div>

      <!-- Audio -->
      <div v-else-if="isAudioFile" class="p-6 flex items-center justify-center h-full">
        <div class="text-center max-w-full">
          <div class="mb-4">
            <MusicalNoteIcon class="w-16 h-16 text-gray-400 mx-auto mb-4" />
          </div>
          <audio 
            :src="getFileUrl(file?.path)"
            controls
            class="w-full max-w-md"
          >
            Your browser does not support the audio tag.
          </audio>
          <p class="text-sm text-gray-500 mt-4">{{ file?.name }}</p>
        </div>
      </div>

      <!-- Other Files -->
      <div v-else class="p-6 flex items-center justify-center h-full">
        <div class="text-center">
          <DocumentIcon class="w-16 h-16 text-gray-400 mx-auto mb-4" />
          <p class="text-gray-600 mb-2">{{ getFileTypeDescription() }}</p>
          <p class="text-sm text-gray-500 mb-4">{{ file?.name }} ({{ formatFileSize(file?.size || 0) }})</p>
          <BaseButton
            variant="primary"
            size="sm"
            @click="downloadFile"
          >
            Download File
          </BaseButton>
        </div>
      </div>
    </div>

    <template #footer>
      <BaseButton variant="secondary" @click="handleClose">
        Close
      </BaseButton>
    </template>
  </BaseModal>
</template>

<script setup>
import { computed } from 'vue'
import {
  DocumentIcon,
  ClipboardIcon,
  PhotoIcon,
  MusicalNoteIcon
} from '@heroicons/vue/24/outline'
import BaseModal from './common/BaseModal.vue'
import BaseButton from './common/BaseButton.vue'
import { useToast } from '../composables/useToast'
import { formatFileSize, formatDate } from '../utils/formatters'

const props = defineProps({
  modelValue: {
    type: Boolean,
    required: true
  },
  file: {
    type: Object,
    default: null
  },
  content: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue', 'close'])

const toast = useToast()

const handleClose = () => {
  emit('update:modelValue', false)
  emit('close')
}

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
  return textTypes.some(type => props.file?.mimeType?.startsWith(type))
})

const isImageFile = computed(() => {
  return props.file?.mimeType?.startsWith('image/')
})

const isPdfFile = computed(() => {
  return props.file?.mimeType === 'application/pdf'
})

const isVideoFile = computed(() => {
  return props.file?.mimeType?.startsWith('video/')
})

const isAudioFile = computed(() => {
  return props.file?.mimeType?.startsWith('audio/')
})

const getFileUrl = (path) => {
  const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'
  return `${API_BASE_URL}/file/raw?path=${encodeURIComponent(path)}`
}

const getFileTypeDescription = () => {
  const mimeType = props.file?.mimeType
  if (!mimeType) return 'Unknown file type'
  
  if (mimeType.includes('zip') || mimeType.includes('archive')) return 'Archive file'
  if (mimeType.includes('executable')) return 'Executable file'
  if (mimeType.includes('office') || mimeType.includes('word') || mimeType.includes('excel')) return 'Office document'
  if (mimeType === 'application/octet-stream') return 'Binary file'
  
  return `${mimeType.split('/')[0]} file`.replace(/^\w/, c => c.toUpperCase())
}

const downloadFile = () => {
  const url = getFileUrl(props.file?.path)
  const link = document.createElement('a')
  link.href = url
  link.download = props.file?.name || 'download'
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
  toast.success('Download started', `Downloading ${props.file?.name}`)
}

const handleImageError = () => {
  toast.error('Image load failed', 'Could not load the image file')
}

const handleImageLoad = () => {
  console.log('Image loaded successfully')
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
</script>