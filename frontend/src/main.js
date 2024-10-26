import './assets/main.css'

import router from './routes/index'
import stores from './store'

import { createApp } from 'vue'
import App from './App.vue'

createApp(App).use(router).use(stores).mount('#app')
