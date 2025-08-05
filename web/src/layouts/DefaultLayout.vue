<script setup>
import { RouterLink } from "vue-router";
</script>

<template>
  <div class="min-h-screen flex flex-col">
    <nav class="bg-gray-800 text-white p-4">
      <div class="text-xl font-bold">Moyoki</div>
    </nav>
    <div class="flex flex-1">
      <aside class="w-64 bg-gray-100 p-4">
        <nav class="space-y-2">
          <RouterLink
            to="/"
            class="block px-4 py-2 text-gray-700 hover:bg-gray-200 rounded"
            >Home
          </RouterLink>
          <RouterLink
            to="/routes"
            class="block px-4 py-2 text-gray-700 hover:bg-gray-200 rounded"
            >Routes
          </RouterLink>
          <RouterLink
            to="/about"
            class="block px-4 py-2 text-gray-700 hover:bg-gray-200 rounded"
            >About
          </RouterLink>
        </nav>
      </aside>
      <main class="flex-1 p-4">
        <slot />
      </main>
    </div>
  </div>
</template>

<style scoped></style>
