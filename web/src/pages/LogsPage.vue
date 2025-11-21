<script setup>
import { ref, onMounted } from "vue";

const logs = ref([]);
const isLoading = ref(false);
const error = ref("");

// Helpers to format and parse datetime-local values
function toDateTimeLocalString(d) {
  const pad = (n) => String(n).padStart(2, "0");
  const year = d.getFullYear();
  const month = pad(d.getMonth() + 1);
  const day = pad(d.getDate());
  const hours = pad(d.getHours());
  const minutes = pad(d.getMinutes());
  return `${year}-${month}-${day}T${hours}:${minutes}`;
}

function toUnixSeconds(dtLocalStr) {
  // dtLocalStr like "2025-11-21T15:04"
  const d = new Date(dtLocalStr);
  return Math.floor(d.getTime() / 1000);
}

// Default range: start = now - 1 day, end = now
const now = new Date();
const oneDayMs = 24 * 60 * 60 * 1000;
const defaultStart = toDateTimeLocalString(new Date(now.getTime() - oneDayMs));
// set default end to end of this day
now.setHours(23, 59, 59, 999);
const defaultEnd = toDateTimeLocalString(now);

const startInput = ref(defaultStart);
const endInput = ref(defaultEnd);

async function fetchLogs() {
  error.value = "";
  isLoading.value = true;
  try {
    const startTs = toUnixSeconds(startInput.value);
    const endTs = toUnixSeconds(endInput.value);
    const res = await fetch(`/api/logs?start=${startTs}&end=${endTs}`);
    if (!res.ok) {
      throw new Error(`HTTP ${res.status}`);
    }
    const body = await res.json();
    logs.value = Array.isArray(body?.data) ? body.data : [];
  } catch (e) {
    error.value = e?.message || String(e);
  } finally {
    isLoading.value = false;
  }
}

onMounted(() => {
  // Auto-fetch with default params (end=now, start=now-1day)
  fetchLogs();
});
</script>

<template>
  <div class="container mx-auto p-4">
    <h1 class="text-2xl font-bold mb-4">Access Logs</h1>

    <div class="bg-white p-4 rounded shadow mb-4">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4 items-end">
        <div>
          <label class="block text-sm font-medium mb-1" for="start">Start</label>
          <input
            id="start"
            type="datetime-local"
            v-model="startInput"
            class="w-full border rounded px-3 py-2"
          />
        </div>
        <div>
          <label class="block text-sm font-medium mb-1" for="end">End</label>
          <input
            id="end"
            type="datetime-local"
            v-model="endInput"
            class="w-full border rounded px-3 py-2"
          />
        </div>
        <div class="md:col-span-2 flex gap-2">
          <button
            class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 disabled:opacity-50"
            :disabled="isLoading"
            @click="fetchLogs"
          >
            {{ isLoading ? "Loading..." : "Refresh" }}
          </button>
          <span v-if="error" class="text-red-600 self-center">{{ error }}</span>
        </div>
      </div>
    </div>

    <div class="bg-white rounded shadow overflow-x-auto">
      <table class="min-w-full border-collapse">
        <thead>
          <tr class="bg-gray-50 text-left">
            <th class="px-3 py-2 border-b">Time</th>
            <th class="px-3 py-2 border-b">Method</th>
            <th class="px-3 py-2 border-b">Path</th>
            <th class="px-3 py-2 border-b">Remote</th>
            <th class="px-3 py-2 border-b">Query</th>
            <th class="px-3 py-2 border-b">Header</th>
            <th class="px-3 py-2 border-b">Body</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="row in logs" :key="row.id" class="hover:bg-gray-50">
            <td class="px-3 py-2 border-b">
              {{ new Date(row.createdAt).toLocaleString() }}
            </td>
            <td class="px-3 py-2 border-b">{{ row.method }}</td>
            <td class="px-3 py-2 border-b">{{ row.path }}</td>
            <td class="px-3 py-2 border-b">{{ row.remoteAddr }}</td>
            <td class="px-3 py-2 border-b truncate max-w-[20ch]">
              {{ row.reqQuery || "-" }}
            </td>
            <td class="px-3 py-2 border-b truncate max-w-[40ch]">
              <pre v-if="row.reqHeaders" v-text="row.reqHeaders"></pre>
            </td>
            <td class="px-3 py-2 border-b truncate max-w-[40ch]">
              <pre v-if="row.reqBody" v-text="row.reqBody"></pre>
            </td>
          </tr>
          <tr v-if="!isLoading && logs.length === 0">
            <td class="px-3 py-4 text-center text-gray-500" colspan="6">
              No logs in this time range.
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.container { max-width: 1200px; }
</style>
