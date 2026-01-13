import { ref, computed } from 'vue'

export function useNavigation(route, router) {
  const currentPath = ref('/')

  const pathSegments = computed(() => {
    if (!currentPath.value || currentPath.value === '/') return []
    return currentPath.value.split('/').filter(segment => segment.length > 0)
  })

  const navigateToPath = (path) => {
    currentPath.value = path
    router.push(path === '/' ? '/' : `/files${path}`)
  }

  const navigateToParent = () => {
    const parentPath = currentPath.value.split('/').slice(0, -1).join('/') || '/'
    navigateToPath(parentPath)
  }

  const getPathUpTo = (index) => {
    const segments = pathSegments.value.slice(0, index + 1)
    return '/' + segments.join('/')
  }

  return {
    currentPath,
    pathSegments,
    navigateToPath,
    navigateToParent,
    getPathUpTo
  }
}