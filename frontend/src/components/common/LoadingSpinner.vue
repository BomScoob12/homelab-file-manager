<template>
  <div :class="containerClasses">
    <ArrowPathIcon :class="spinnerClasses" />
    <span v-if="text" :class="textClasses">{{ text }}</span>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowPathIcon } from '@heroicons/vue/24/outline'

const props = defineProps({
  size: {
    type: String,
    default: 'md',
    validator: value => ['sm', 'md', 'lg'].includes(value)
  },
  text: {
    type: String,
    default: ''
  },
  center: {
    type: Boolean,
    default: false
  }
})

const containerClasses = computed(() => {
  const baseClasses = 'flex items-center'
  const centerClasses = props.center ? 'justify-center' : ''
  const spacingClasses = props.text ? 'space-x-2' : ''
  
  return [baseClasses, centerClasses, spacingClasses]
})

const spinnerClasses = computed(() => {
  const baseClasses = 'animate-spin text-primary-600'
  
  const sizeClasses = {
    sm: 'w-4 h-4',
    md: 'w-6 h-6',
    lg: 'w-8 h-8'
  }
  
  return [baseClasses, sizeClasses[props.size]]
})

const textClasses = computed(() => {
  const sizeClasses = {
    sm: 'text-sm',
    md: 'text-base',
    lg: 'text-lg'
  }
  
  return ['text-gray-600', sizeClasses[props.size]]
})
</script>