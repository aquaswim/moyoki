<script setup>
import { ref, onMounted } from "vue";
import { useRouter, useRoute } from "vue-router";
import Swal from "sweetalert2";

const router = useRouter();
const route = useRoute();
const routeId = route.params.id;
const isLoading = ref(true);
const httpMethods = [
  "GET",
  "POST",
  "PUT",
  "DELETE",
  "PATCH",
  "OPTIONS",
  "HEAD",
];

const routeData = ref({
  id: null,
  method: "GET",
  path: "",
  description: "",
  responseCode: 200,
  responseHeaders: [{ key: "", value: "" }],
  responseBody: "",
});

// Fetch route data on component mount
onMounted(async () => {
  try {
    const res = await fetch(`/api/routes/${routeId}`);
    const resp = await res.json();

    routeData.value = resp.data;

    isLoading.value = false;
  } catch (error) {
    console.error("Error fetching route data:", error);
    isLoading.value = false;
  }
});

const addHeader = () => {
  if (!routeData.value.responseHeaders) {
    routeData.value.responseHeaders = [];
  }
  routeData.value.responseHeaders.push({ key: "", value: "" });
};

const removeHeader = (index) => {
  routeData.value.responseHeaders.splice(index, 1);
};

const updateRoute = async () => {
  try {
    isLoading.value = true;
    const response = await fetch(`/api/routes/${routeData.value.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(routeData.value),
    });

    if (!response.ok) {
      throw new Error("Failed to update route");
    }

    isLoading.value = false;
    router.push("/routes");
  } catch (error) {
    isLoading.value = false;
    Swal.fire({
      icon: "error",
      title: "Error",
      text: error.message || "Failed to update route",
    });
  }
};
</script>

<template>
  <div class="container mx-auto px-4 py-6">
    <div class="mb-6">
      <h1 class="text-2xl font-bold">Update Route</h1>
      <p class="text-gray-600">Edit route ID: {{ routeId }}</p>
    </div>

    <div v-if="isLoading" class="flex justify-center items-center h-64">
      <div
        class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500"
      ></div>
    </div>

    <form
      v-else
      class="bg-white rounded-lg shadow-md p-6"
      @submit.prevent="updateRoute"
    >
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
        <!-- HTTP Method -->
        <div>
          <label
            for="method"
            class="block text-sm font-medium text-gray-700 mb-1"
            >HTTP Method</label
          >
          <select
            id="method"
            v-model="routeData.method"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option v-for="method in httpMethods" :key="method" :value="method">
              {{ method }}
            </option>
          </select>
        </div>

        <!-- Path -->
        <div>
          <label
            for="path"
            class="block text-sm font-medium text-gray-700 mb-1"
          >
            Path
          </label>
          <input
            id="path"
            v-model="routeData.path"
            type="text"
            placeholder="/api/resource/{id}"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          />
          <p class="mt-1 text-sm text-gray-500">
            Use {param} for dynamic path segments
          </p>
        </div>
      </div>

      <!-- Description -->
      <div class="mb-6">
        <label
          for="description"
          class="block text-sm font-medium text-gray-700 mb-1"
          >Description</label
        >
        <textarea
          id="description"
          v-model="routeData.description"
          rows="3"
          class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          placeholder="Describe what this route does"
        ></textarea>
      </div>

      <!-- Response Code -->
      <div class="mb-6">
        <label
          for="responseCode"
          class="block text-sm font-medium text-gray-700 mb-1"
          >Response Code</label
        >
        <input
          id="responseCode"
          v-model="routeData.responseCode"
          type="number"
          class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
        />
      </div>

      <!-- Response Headers -->
      <div class="mb-6">
        <label class="block text-sm font-medium text-gray-700 mb-2"
          >Response Headers</label
        >

        <div
          v-for="(header, index) in routeData.responseHeaders"
          :key="index"
          class="flex gap-2 mb-2"
        >
          <input
            v-model="header.key"
            type="text"
            placeholder="Header name"
            class="w-1/2 px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          />
          <input
            v-model="header.value"
            type="text"
            placeholder="Header value"
            class="w-1/2 px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          />
          <button
            type="button"
            class="bg-red-500 text-white p-2 rounded-md hover:bg-red-600 focus:outline-none"
            @click="removeHeader(index)"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                clip-rule="evenodd"
              />
            </svg>
          </button>
        </div>

        <button
          type="button"
          class="mt-2 inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="addHeader"
        >
          Add Header
        </button>
      </div>

      <!-- Response Body -->
      <div class="mb-6">
        <label
          for="responseBody"
          class="block text-sm font-medium text-gray-700 mb-1"
          >Response Body</label
        >
        <textarea
          id="responseBody"
          v-model="routeData.responseBody"
          rows="6"
          class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 font-mono"
          placeholder='{"data": "example response"}'
        ></textarea>
      </div>

      <!-- Form Actions -->
      <div class="flex justify-end gap-3">
        <button
          type="button"
          class="px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="router.push('/routes')"
        >
          Cancel
        </button>
        <button
          type="submit"
          class="px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          Update Route
        </button>
      </div>
    </form>
  </div>
</template>

<style scoped>
/* Additional styles can be added here if needed */
</style>
