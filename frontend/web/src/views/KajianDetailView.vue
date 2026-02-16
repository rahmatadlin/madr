<template>
  <main class="min-h-screen bg-gray-50">
    <div class="container mx-auto px-4 py-12">
      <div v-if="loading" class="text-center py-12 text-gray-600">Memuat...</div>
      <template v-else-if="video">
        <div class="max-w-4xl mx-auto">
          <router-link to="/" class="inline-flex items-center gap-2 text-gray-600 hover:text-gray-900 mb-6">
            ← Kembali ke Beranda
          </router-link>
          <p class="text-sm text-gray-500 mb-2">{{ formatDate(video.published_at) }}</p>
          <h1 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">{{ video.title }}</h1>
          <div class="aspect-video w-full rounded-xl overflow-hidden bg-black mb-6">
            <iframe
              :src="embedUrl"
              title="YouTube video"
              class="w-full h-full"
              allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
              allowfullscreen
            />
          </div>
          <div v-if="video.description" class="prose prose-gray max-w-none">
            <p class="text-gray-600 whitespace-pre-wrap">{{ video.description }}</p>
          </div>
        </div>
      </template>
      <div v-else class="text-center py-12 text-gray-600">Video tidak ditemukan.</div>
    </div>
    <Footer />
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { kajianApi } from '@/api/kajian'
import type { Kajian } from '@/api/kajian'
import Footer from '@/components/layouts/Footer.vue'

const route = useRoute()
const video = ref<Kajian | null>(null)
const loading = ref(true)

const embedUrl = computed(() => {
  if (!video.value) return ''
  const videoId = video.value.youtube_url.match(/(?:youtube\.com\/watch\?v=|youtu\.be\/)([^&\n?#]+)/)?.[1] || video.value.video_id
  return `https://www.youtube.com/embed/${videoId}`
})

function formatDate(d: string) {
  return new Date(d).toLocaleDateString('id-ID', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' })
}

onMounted(async () => {
  const id = route.params.id as string
  if (!id) {
    loading.value = false
    return
  }
  try {
    const res = await kajianApi.getById(parseInt(id, 10))
    video.value = res.data || null
  } catch {
    video.value = null
  } finally {
    loading.value = false
  }
})
</script>
