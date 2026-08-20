<script setup>
import { ref, onMounted } from "vue";
import { useRouter, useRoute } from "vue-router";
import Swal from "sweetalert2";
import RouteForm from "../../components/RouteForm.vue";

const router = useRouter();
const route = useRoute();
const routeId = route.params.id;
const isLoading = ref(true);

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

const updateRoute = async () => {
  isLoading.value = true;
  try {
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

    router.push("/routes");
  } catch (error) {
    console.error("Error updating route:", error);
    Swal.fire({
      icon: "error",
      title: "Error",
      text: error.message || "Failed to update route",
    });
  } finally {
    isLoading.value = false;
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

    <RouteForm
      v-else
      v-model="routeData"
      submit-label="Update Route"
      :is-loading="isLoading"
      @submit="updateRoute"
      @cancel="router.push('/routes')"
    />
  </div>
</template>

<style scoped>
/* Additional styles can be added here if needed */
</style>
