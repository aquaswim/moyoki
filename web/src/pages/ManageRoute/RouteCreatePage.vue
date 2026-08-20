<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import Swal from "sweetalert2";
import RouteForm from "../../components/RouteForm.vue";

const router = useRouter();

const route = ref({
  method: "GET",
  path: "",
  description: "",
  responseCode: 200,
  responseHeaders: [{ key: "", value: "" }],
  responseBody: "",
});
const isLoading = ref(false);

const saveRoute = async () => {
  isLoading.value = true;
  try {
    // Here you would typically make an API call to save the route
    const { responseHeaders, ...reqBody } = route.value;
    reqBody.responseHeaders = [];
    responseHeaders.forEach((header) => {
      if (header.key) {
        reqBody.responseHeaders.push({
          key: header.key,
          value: header.value,
        });
      }
    });

    console.log("Saving route:", reqBody);

    // save to api
    const response = await fetch("/api/routes", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(reqBody),
    });

    if (!response.ok) {
      throw new Error("Failed to save route");
    }
    // Redirect back to the routes list
    router.push("/routes");
  } catch (error) {
    console.error("Error saving route:", error);
    Swal.fire({
      icon: "error",
      title: "Error",
      text: "Failed to save route. Please try again.",
    });
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <div class="container mx-auto px-4 py-6">
    <div class="mb-6">
      <h1 class="text-2xl font-bold">Create New Route</h1>
      <p class="text-gray-600">Define a new API route with response details</p>
    </div>

    <RouteForm
      v-model="route"
      submit-label="Save Route"
      :is-loading="isLoading"
      @submit="saveRoute"
      @cancel="router.push('/routes')"
    />
  </div>
</template>

<style scoped>
/* Additional styles can be added here if needed */
</style>
