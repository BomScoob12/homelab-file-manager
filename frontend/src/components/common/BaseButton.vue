<template>
  <button
    :class="buttonClasses"
    :disabled="disabled || loading"
    @click="$emit('click', $event)"
  >
    <component
      :is="loadingIcon"
      v-if="loading"
      class="w-4 h-4 animate-spin"
      :class="{ 'mr-2': $slots.default }"
    />
    <component
      :is="icon"
      v-else-if="icon"
      class="w-4 h-4"
      :class="{ 'mr-2': $slots.default }"
    />
    <slot />
  </button>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowPathIcon } from '@heroicons/vue/24/outline'

const props = defineProps({
  variant: {
    type: String,
    default: 'primary',
    validator: value => ['primary', 'secondary', 'danger', 'ghost'].includes(value)
  },
  size: {
    type: String,
    default: 'md',
    validator: value => ['sm', 'md', 'lg'].includes(value)
  },
  disabled: {
    type: Boolean,
    default: false
  },
  loading: {
    type: Boolean,
    default: false
  },
  icon: {
    type: Object,
    default: null
  }
})

defineEmits(['click'])

const loadingIcon = ArrowPathIcon

const buttonClasses = computed(() => {
  const baseClasses = 'inline-flex items-center justify-center font-medium rounded-md transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2'
  
  const variantClasses = {
    primary: 'bg-primary-600 text-white hover:bg-primary-700 focus:ring-primary-500 disabled:bg-primary-300',
    secondary: 'bg-gray-200 text-gray-700 hover:bg-gray-300 focus:ring-gray-500 disabled:bg-gray-100',
    danger: 'bg-red-600 text-white hover:bg-red-700 focus:ring-red-500 disabled:bg-red-300',
    ghost: 'text-gray-400 hover:text-gray-600 hover:bg-gray-100 focus:ring-gray-500'
  }
  
  const sizeClasses = {
    sm: 'px-2 py-1 text-xs',
    md: 'px-4 py-2 text-sm',
    lg: 'px-6 py-3 text-base'
  }
  
  return [
    baseClasses,
    variantClasses[props.variant],
    sizeClasses[props.size],
    {
      'opacity-50 cursor-not-allowed': props.disabled || props.loading
    }
  ]
})
</script>