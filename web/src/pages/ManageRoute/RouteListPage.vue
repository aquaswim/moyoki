<script setup>
import { ref, onMounted } from "vue";
import { RouterLink } from "vue-router";
import Swal from "sweetalert2";

const routes = ref([]);
const loading = ref(true);

const fetchRoutes = async () => {
  try {
    loading.value = true;
    const response = await fetch("/api/routes");
    const data = await response.json();
    routes.value = data.data;
  } catch (error) {
    console.error("Error fetching routes:", error);
    Swal.fire({
      title: "Error",
      text: "Error fetching routes",
    });
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchRoutes();
});

const deleteRoute = (id) => {
  Swal.fire({
    title: "Are you sure?",
    text: "You won't be able to revert this!",
    icon: "warning",
    showCancelButton: true,
    confirmButtonText: "Yes, delete it!",
    showLoaderOnConfirm: true,
    preConfirm: () => {
      return new Promise((resolve) => {
        fetch(`/api/routes/${id}`, {
          method: "DELETE",
        })
          .catch((err) => {
            console.log("delete route error", err);
            Swal.fire({
              title: "Error",
              text: "Error delete routes",
            });
          })
          .finally(() => {
            resolve();
            return fetchRoutes();
          });
      });
    },
    allowOutsideClick: () => !Swal.isLoading(),
  }).then((result) => {
    if (result.isConfirmed) {
      Swal.fire("Deleted!", "Route has been deleted.", "success");
    }
  });
};
</script>

<template>
  <div class="container mx-auto">
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-bold">Manage Routes</h1>
      <RouterLink
        to="/routes/new"
        class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
      >
        Create
      </RouterLink>
    </div>

    <div class="overflow-x-auto">
      <div v-if="loading" class="text-center py-4">Loading...</div>
      <table v-else class="min-w-full bg-white border border-gray-300">
        <thead>
          <tr class="bg-gray-100">
            <th class="py-2 px-4 border-b text-left">ID</th>
            <th class="py-2 px-4 border-b text-left">Method</th>
            <th class="py-2 px-4 border-b text-left">Path</th>
            <th class="py-2 px-4 border-b text-left">Description</th>
            <th class="py-2 px-4 border-b text-left">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="route in routes" :key="route.id" class="hover:bg-gray-50">
            <td class="py-2 px-4 border-b-gray-100">{{ route.id }}</td>
            <td class="py-2 px-4 border-b-gray-100">{{ route.method }}</td>
            <td class="py-2 px-4 border-b-gray-100">{{ route.path }}</td>
            <td class="py-2 px-4 border-b-gray-100">{{ route.description }}</td>
            <td class="py-2 px-4 border-b-gray-100">
              <div class="flex gap-2">
                <RouterLink
                  :to="`/routes/${route.id}/edit`"
                  class="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-1 px-3 rounded"
                >
                  Edit
                </RouterLink>
                <button
                  class="bg-red-500 hover:bg-red-700 text-white font-bold py-1 px-3 rounded"
                  @click="deleteRoute(route.id)"
                >
                  Delete
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
table {
  @apply w-full table-auto;
}
</style>
