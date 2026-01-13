<template>
  <div :class="cardClasses">
    <div v-if="$slots.header" class="px-6 py-4 border-b border-gray-200">
      <slot name="header" />
    </div>
    <div :class="bodyClasses">
      <slot />
    </div>
    <div v-if="$slots.footer" class="px-6 py-4 border-t border-gray-200">
      <slot name="footer" />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  padding: {
    type: String,
    default: 'md',
    validator: value => ['none', 'sm', 'md', 'lg'].includes(value)
  },
  shadow: {
    type: String,
    default: 'sm',
    validator: value => ['none', 'sm', 'md', 'lg'].includes(value)
  }
})

const cardClasses = computed(() => {
  const baseClasses = 'bg-white rounded-lg border border-gray-200'
  
  const shadowClasses = {
    none: '',
    sm: 'shadow-sm',
    md: 'shadow-md',
    lg: 'shadow-lg'
  }
  
  return [baseClasses, shadowClasses[props.shadow]]
})

const bodyClasses = computed(() => {
  const paddingClasses = {
    none: '',
    sm: 'p-3',
    md: 'p-6',
    lg: 'p-8'
  }
  
  return paddingClasses[props.padding]
})
</script>