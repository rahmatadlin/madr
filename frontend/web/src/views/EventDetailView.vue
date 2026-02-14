<template>
  <main class="min-h-screen">
    <div class="container mx-auto px-4 py-20">
      <div v-if="loading" class="text-center text-gray-600">Memuat...</div>
      <div v-else-if="event" class="max-w-2xl mx-auto">
        <h1 class="text-3xl font-bold text-gray-900 mb-4">{{ event.title }}</h1>
        <p class="text-gray-600 mb-4">{{ event.description }}</p>
        <p class="text-sm text-gray-600">{{ formatDate(event.date) }}</p>
        <p v-if="event.location" class="text-sm text-gray-600">{{ event.location }}</p>
        <router-link to="/" class="inline-block mt-6 text-blue-600 hover:underline">Kembali</router-link>
      </div>
      <div v-else class="text-center text-gray-600">Acara tidak ditemukan.</div>
    </div>
    <Footer />
  </main>
</template>
<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { eventApi } from '@/api/events'
import type { Event } from '@/api/events'
import Footer from '@/components/layouts/Footer.vue'

const route = useRoute()
const event = ref<Event | null>(null)
const loading = ref(true)

function formatDate(d: string) {
  return new Date(d).toLocaleDateString('id-ID', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' })
}

onMounted(async () => {
  const id = Number(route.params.id)
  if (!id) { loading.value = false; return }
  try {
    event.value = await eventApi.getById(id)
  } catch {
    event.value = null
  } finally {
    loading.value = false
  }
})
</script>
