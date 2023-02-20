import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Unocss from 'unocss/vite'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [Unocss(),vue()],
  resolve: {
    alias: {
      
    }
  },
  server: {
    proxy: {
      '^/api/.*': {
        target: 'http://localhost:8474',
        changeOrigin: true,
        //rewrite: (path) => path.replace(/^\/api/, ''),
      },
    }
  }
})
