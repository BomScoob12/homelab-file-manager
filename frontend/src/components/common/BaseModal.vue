<template>
  <Teleport to="body">
    <Transition
      enter-active-class="transition-opacity duration-300"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition-opacity duration-300"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div
        v-if="modelValue"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50"
        @click="handleBackdropClick"
      >
        <Transition
          enter-active-class="transition-all duration-300"
          enter-from-class="opacity-0 scale-95"
          enter-to-class="opacity-100 scale-100"
          leave-active-class="transition-all duration-300"
          leave-from-class="opacity-100 scale-100"
          leave-to-class="opacity-0 scale-95"
        >
          <div
            v-if="modelValue"
            :class="modalClasses"
            @click.stop
          >
            <!-- Header -->
            <div v-if="$slots.header || title" class="flex items-center justify-between p-6 border-b border-gray-200">
              <div class="flex items-center space-x-3">
                <slot name="header">
                  <h3 class="text-lg font-semibold text-gray-900">{{ title }}</h3>
                </slot>
              </div>
              <BaseButton
                v-if="closable"
                variant="ghost"
                size="sm"
                :icon="XMarkIcon"
                @click="close"
              />
            </div>

            <!-- Body -->
            <div :class="bodyClasses">
              <slot />
            </div>

            <!-- Footer -->
            <div v-if="$slots.footer" class="flex justify-end space-x-3 p-6 border-t border-gray-200">
              <slot name="footer" />
            </div>
          </div>
        </Transition>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { computed } from 'vue'
import { XMarkIcon } from '@heroicons/vue/24/outline'
import BaseButton from './BaseButton.vue'

const props = defineProps({
  modelValue: {
    type: Boolean,
    required: true
  },
  title: {
    type: String,
    default: ''
  },
  size: {
    type: String,
    default: 'md',
    validator: value => ['sm', 'md', 'lg', 'xl', 'full'].includes(value)
  },
  closable: {
    type: Boolean,
    default: true
  },
  closeOnBackdrop: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits(['update:modelValue', 'close'])

const modalClasses = computed(() => {
  const baseClasses = 'bg-white rounded-lg shadow-xl flex flex-col'
  
  const sizeClasses = {
    sm: 'max-w-md w-full max-h-[80vh]',
    md: 'max-w-2xl w-full max-h-[85vh]',
    lg: 'max-w-4xl w-full max-h-[90vh]',
    xl: 'max-w-6xl w-full max-h-[95vh]',
    full: 'w-full h-full max-w-none max-h-none m-0 rounded-none'
  }
  
  return [baseClasses, sizeClasses[props.size]]
})

const bodyClasses = computed(() => {
  return 'flex-1 overflow-hidden'
})

const close = () => {
  emit('update:modelValue', false)
  emit('close')
}

const handleBackdropClick = () => {
  if (props.closeOnBackdrop) {
    close()
  }
}
</script>