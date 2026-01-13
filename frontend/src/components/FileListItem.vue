<template>
  <div class="p-4 hover:bg-gray-50 transition-colors duration-150">
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
            class="text-left w-full group"
            @click="$emit('click', file)"
          >
            <p class="text-sm font-medium text-gray-900 truncate group-hover:text-primary-600 transition-colors">
              {{ file.name }}
            </p>
            <p class="text-xs text-gray-500">
              <span v-if="!file.isDir">{{ formatFileSize(file.size) }} • </span>
              {{ formatDate(file.modTime) }}
              <span v-if="file.extension" class="ml-1 text-gray-400">{{ file.extension }}</span>
            </p>
          </button>
        </div>
      </div>

      <!-- Actions -->
      <div class="flex items-center space-x-2">
        <BaseButton
          v-if="!file.isDir"
          variant="ghost"
          size="sm"
          :icon="EyeIcon"
          title="Open file"
          @click.stop="$emit('open', file)"
        />
        <BaseButton
          v-if="!file.isDir"
          variant="ghost"
          size="sm"
          :icon="ArrowDownTrayIcon"
          title="Download file"
          @click.stop="$emit('download', file)"
        />
        <BaseButton
          variant="ghost"
          size="sm"
          :icon="TrashIcon"
          title="Delete"
          class="text-gray-400 hover:text-red-600"
          @click.stop="$emit('delete', file)"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import {
  FolderIcon,
  DocumentIcon,
  EyeIcon,
  TrashIcon,
  ArrowDownTrayIcon
} from '@heroicons/vue/24/outline'
import BaseButton from './common/BaseButton.vue'
import { formatFileSize, formatDate } from '../utils/formatters'

defineProps({
  file: {
    type: Object,
    required: true
  }
})

defineEmits(['click', 'open', 'download', 'delete'])
</script>