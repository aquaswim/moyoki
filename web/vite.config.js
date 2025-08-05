import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import devManifest from "vite-plugin-dev-manifest";
import eslint from "vite-plugin-eslint";
import tailwindcss from "@tailwindcss/vite";

// https://vite.dev/config/
export default defineConfig({
  server: {
    port: 5173,
  },
  plugins: [eslint(), vue(), devManifest(), tailwindcss()],
  build: {
    manifest: true,
    rollupOptions: {
      input: {
        main: "src/main.js",
      },
    },
  },
});
