<script setup>
import { ref } from "vue";
import { RouterLink } from "vue-router";
import Swal from 'sweetalert2';

const routes = ref([
  { id: 1, path: "/api/users", description: "Get all users" },
  { id: 2, path: "/api/products", description: "Get all products" },
]);

const deleteRoute = (id) => {
  Swal.fire({
    title: "Are you sure?",
    text: "You won't be able to revert this!",
    icon: "warning",
    showCancelButton: true,
    confirmButtonText: "Yes, delete it!"
  }).then((result) => {
    if (result.isConfirmed) {
      // todo do remove
      console.log("delete", id);
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
      <table class="min-w-full bg-white border border-gray-300">
        <thead>
          <tr class="bg-gray-100">
            <th class="py-2 px-4 border-b text-left">ID</th>
            <th class="py-2 px-4 border-b text-left">Path</th>
            <th class="py-2 px-4 border-b text-left">Description</th>
            <th class="py-2 px-4 border-b text-left">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="route in routes" :key="route.id" class="hover:bg-gray-50">
            <td class="py-2 px-4 border-b-gray-100">{{ route.id }}</td>
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
