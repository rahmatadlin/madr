<template>
  <section v-if="videos.length > 0 || isLoading" class="py-20 bg-gray-50">
    <div class="container mx-auto px-4">
      <div class="text-center mb-12">
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">Kajian Minggu Lalu</h2>
        <p class="text-lg text-gray-600">Rekaman kajian dan tausiyah dari YouTube</p>
      </div>

      <p v-if="isLoading" class="text-center text-gray-600">Memuat kajian...</p>
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <router-link
          v-for="video in videos"
          :key="video.video_id"
          :to="'/kajian/' + video.video_id"
          class="group block rounded-xl border bg-white overflow-hidden shadow-sm hover:shadow-lg transition-shadow"
        >
          <div class="relative aspect-video bg-gray-200">
            <img
              :src="video.thumbnail_url"
              :alt="video.title"
              class="absolute inset-0 w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
            />
            <div class="absolute inset-0 flex items-center justify-center bg-black/30 opacity-0 group-hover:opacity-100 transition-opacity">
              <span class="rounded-full bg-red-600 p-3 text-white">▶</span>
            </div>
          </div>
          <div class="p-4">
            <p class="text-sm text-gray-500 mb-1">{{ formatDate(video.published_at) }}</p>
            <h3 class="font-semibold text-gray-900 mb-2 line-clamp-2 group-hover:text-blue-600 transition-colors">{{ video.title }}</h3>
            <p class="text-sm text-gray-600 line-clamp-2">{{ video.description || 'Tonton kajian selengkapnya.' }}</p>
          </div>
        </router-link>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { youtubeApi } from '@/api/youtube'
import type { YouTubeVideo } from '@/api/youtube'

const videos = ref<YouTubeVideo[]>([])
const isLoading = ref(true)

function formatDate(d: string) {
  return new Date(d).toLocaleDateString('id-ID', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' })
}

onMounted(async () => {
  try {
    const res = await youtubeApi.getKajianVideos()
    videos.value = res.data ?? []
  } finally {
    isLoading.value = false
  }
})
</script>
