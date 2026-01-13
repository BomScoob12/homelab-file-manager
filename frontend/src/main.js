import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import FileManager from './components/FileManager.vue'
import './style.css'

const routes = [
  { path: '/', component: FileManager },
  { path: '/files/:path*', component: FileManager, props: true }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

createApp(App).use(router).mount('#app')