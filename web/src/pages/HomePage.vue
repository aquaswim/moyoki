<script setup>
import { ref, onMounted } from "vue";

// Hardcoded routes count
const routesCount = ref(0);
const isLoading = ref(true);

// Simulate loading state
onMounted(() => {
  setTimeout(() => {
    routesCount.value = 2; // Hardcoded value matching the count in RouteListPage
    isLoading.value = false;
  }, 1500); // Simulate network delay
});
</script>

<template>
  <div class="container mx-auto p-4">
    <h1 class="text-2xl font-bold mb-6">HomePage</h1>

    <div class="bg-white p-6 rounded-lg shadow-md">
      <h2 class="text-xl font-semibold mb-2">Routes Statistics</h2>
      <div class="flex items-center">
        <div v-if="isLoading" class="flex items-center">
          <div
            class="animate-spin rounded-full h-5 w-5 border-b-2 border-blue-500 mr-3"
          ></div>
          <span>Loading routes count...</span>
        </div>
        <div v-else class="text-3xl font-bold text-blue-600">
          {{ routesCount }}
          <span class="text-gray-600 text-lg font-normal"
            >routes configured</span
          >
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  max-width: 1200px;
}
</style>
