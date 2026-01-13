<template>
  <BaseModal
    :model-value="modelValue"
    title="Delete Confirmation"
    size="sm"
    @update:model-value="$emit('update:modelValue', $event)"
    @close="$emit('cancel')"
  >
    <template #header>
      <div class="flex items-center">
        <div class="flex-shrink-0">
          <ExclamationTriangleIcon class="w-6 h-6 text-red-600" />
        </div>
        <div class="ml-3">
          <h3 class="text-lg font-semibold text-gray-900">
            Delete {{ file?.isDir ? 'Directory' : 'File' }}
          </h3>
        </div>
      </div>
    </template>

    <!-- Content -->
    <div class="px-6 pb-6">
      <p class="text-sm text-gray-600 mb-4">
        Are you sure you want to delete 
        <span class="font-medium text-gray-900">{{ file?.name }}</span>?
        <span v-if="file?.isDir" class="text-red-600">
          This will delete the directory and all its contents.
        </span>
        This action cannot be undone.
      </p>

      <BaseCard padding="sm" class="bg-gray-50">
        <div class="flex items-center space-x-2">
          <FolderIcon v-if="file?.isDir" class="w-4 h-4 text-blue-500" />
          <DocumentIcon v-else class="w-4 h-4 text-gray-400" />
          <span class="text-sm font-medium text-gray-700">{{ file?.path }}</span>
        </div>
        <div class="mt-1 text-xs text-gray-500">
          {{ formatFileSize(file?.size || 0) }} • {{ formatDate(file?.modTime) }}
        </div>
      </BaseCard>
    </div>

    <template #footer>
      <BaseButton variant="secondary" @click="$emit('cancel')">
        Cancel
      </BaseButton>
      <BaseButton variant="danger" :icon="TrashIcon" @click="$emit('confirm')">
        Delete {{ file?.isDir ? 'Directory' : 'File' }}
      </BaseButton>
    </template>
  </BaseModal>
</template>

<script setup>
import {
  ExclamationTriangleIcon,
  FolderIcon,
  DocumentIcon,
  TrashIcon
} from '@heroicons/vue/24/outline'
import BaseModal from './common/BaseModal.vue'
import BaseButton from './common/BaseButton.vue'
import BaseCard from './common/BaseCard.vue'
import { formatFileSize, formatDate } from '../utils/formatters'

defineProps({
  modelValue: {
    type: Boolean,
    required: true
  },
  file: {
    type: Object,
    default: null
  }
})

defineEmits(['update:modelValue', 'confirm', 'cancel'])
</script>