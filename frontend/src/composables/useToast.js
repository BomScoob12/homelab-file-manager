import { ref, reactive } from 'vue'

const toasts = ref([])
let toastId = 0

export function useToast() {
  const addToast = (toast) => {
    const id = ++toastId
    const newToast = {
      id,
      type: 'info',
      duration: 5000,
      ...toast
    }
    
    toasts.value.push(newToast)
    
    // Auto remove after duration
    if (newToast.duration > 0) {
      setTimeout(() => {
        removeToast(id)
      }, newToast.duration)
    }
    
    return id
  }

  const removeToast = (id) => {
    const index = toasts.value.findIndex(toast => toast.id === id)
    if (index > -1) {
      toasts.value.splice(index, 1)
    }
  }

  const success = (title, message = '') => {
    return addToast({ type: 'success', title, message })
  }

  const error = (title, message = '') => {
    return addToast({ type: 'error', title, message })
  }

  const warning = (title, message = '') => {
    return addToast({ type: 'warning', title, message })
  }

  const info = (title, message = '') => {
    return addToast({ type: 'info', title, message })
  }

  const clear = () => {
    toasts.value = []
  }

  return {
    toasts,
    addToast,
    removeToast,
    success,
    error,
    warning,
    info,
    clear
  }
}