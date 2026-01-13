<template>
  <nav class="flex" aria-label="Breadcrumb">
    <ol class="flex items-center space-x-2 flex-wrap">
      <li>
        <BaseButton
          variant="ghost"
          size="sm"
          :class="{ 'text-primary-800 font-semibold': currentPath === '/' }"
          @click="$emit('navigate', '/')"
        >
          <HomeIcon class="w-4 h-4 mr-1" />
          Home
        </BaseButton>
      </li>
      <li v-for="(segment, index) in pathSegments" :key="index" class="flex items-center">
        <ChevronRightIcon class="w-4 h-4 text-gray-400 mx-2" />
        <BaseButton
          variant="ghost"
          size="sm"
          :class="{ 'text-primary-800 font-semibold': getPathUpTo(index) === currentPath }"
          @click="$emit('navigate', getPathUpTo(index))"
        >
          {{ segment }}
        </BaseButton>
      </li>
    </ol>
  </nav>
</template>

<script setup>
import { computed } from 'vue'
import { HomeIcon, ChevronRightIcon } from '@heroicons/vue/24/outline'
import BaseButton from './common/BaseButton.vue'

const props = defineProps({
  currentPath: {
    type: String,
    required: true
  }
})

defineEmits(['navigate'])

const pathSegments = computed(() => {
  if (!props.currentPath || props.currentPath === '/') return []
  return props.currentPath.split('/').filter(segment => segment.length > 0)
})

const getPathUpTo = (index) => {
  const segments = pathSegments.value.slice(0, index + 1)
  return '/' + segments.join('/')
}
</script>