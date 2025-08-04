import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import devManifest from "vite-plugin-dev-manifest"

// https://vite.dev/config/
export default defineConfig({
  server: {
    port: 5173,
  },
  plugins: [vue(), devManifest()],
  build: {
    manifest: true,
    rollupOptions: {
      input: {
        main: 'src/main.js',
      },
    },
  }
})
