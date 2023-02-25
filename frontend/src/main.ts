import { createApp } from 'vue'
import App from './App.vue'
import { anu } from 'anu-vue'

// UnoCSS import
import 'uno.css'

// anu styles
import 'anu-vue/dist/style.css'

// default theme styles
import '@anu-vue/preset-theme-default/dist/style.css'

//import './assets/main.css'

createApp(App)
    .use(anu)
    .mount('#app')


