<template>
  <section v-if="events.length > 0 || isLoading" class="py-20 bg-gray-50">
    <div class="container mx-auto px-4">
      <div class="text-center mb-12">
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">Program dan Acara</h2>
        <p class="text-lg text-gray-600">Ikuti berbagai kegiatan dan program yang kami adakan</p>
      </div>
      <p v-if="isLoading" class="text-center text-gray-600">Memuat acara...</p>
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="event in events" :key="event.id" class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow">
          <div class="p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-2 line-clamp-2">{{ event.title }}</h3>
            <p class="text-gray-600 text-sm mb-4 line-clamp-2">{{ event.description }}</p>
            <p class="text-sm text-gray-600 mb-4">{{ formatDate(event.date) }}</p>
            <p v-if="event.location" class="text-sm text-gray-600 mb-4">{{ event.location }}</p>
            <router-link :to="'/events/' + event.id" class="block w-full text-center py-2 rounded-lg border border-gray-300 hover:bg-gray-50">Lihat Detail</router-link>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { eventApi } from '@/api/events'
import type { Event } from '@/api/events'
const events = ref<Event[]>([])
const isLoading = ref(true)
function formatDate(d: string) {
  return new Date(d).toLocaleDateString('id-ID', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' })
}
onMounted(async () => {
  try { const res = await eventApi.getAll(6, 0); events.value = res.data || [] } finally { isLoading.value = false }
})
</script>
