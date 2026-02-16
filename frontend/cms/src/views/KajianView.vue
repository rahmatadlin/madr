<template>
  <div class="space-y-6">
    <div>
      <h1 class="text-3xl font-bold">Kajian Minggu Lalu</h1>
      <p class="text-gray-600">Video kajian dari YouTube Channel (30 hari terakhir)</p>
    </div>

    <div v-if="loading" class="text-gray-600">Memuat video dari YouTube...</div>
    <div v-else-if="error" class="text-red-600">{{ error }}</div>
    <div v-else-if="videos.length === 0" class="text-gray-600">Tidak ada video ditemukan.</div>
    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="video in videos"
        :key="video.video_id"
        class="rounded-lg border bg-white overflow-hidden shadow-sm hover:shadow-lg transition-shadow"
      >
        <a
          :href="'https://www.youtube.com/watch?v=' + video.video_id"
          target="_blank"
          rel="noopener noreferrer"
          class="block"
        >
          <div class="relative aspect-video bg-gray-200">
            <img
              :src="video.thumbnail_url"
              :alt="video.title"
              class="w-full h-full object-cover"
            />
            <div class="absolute inset-0 flex items-center justify-center bg-black/30 opacity-0 hover:opacity-100 transition-opacity">
              <span class="rounded-full bg-red-600 p-3 text-white">▶</span>
            </div>
          </div>
          <div class="p-4">
            <p class="text-sm text-gray-500 mb-1">{{ formatDate(video.published_at) }}</p>
            <h3 class="font-semibold text-gray-900 mb-2 line-clamp-2">{{ video.title }}</h3>
            <p class="text-xs text-gray-600 line-clamp-2">{{ video.description }}</p>
          </div>
        </a>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { youtubeApi } from '@/api/youtube'
import type { YouTubeVideo } from '@/api/youtube'

const videos = ref<YouTubeVideo[]>([])
const loading = ref(true)
const error = ref<string | null>(null)

function formatDate(d: string) {
  return new Date(d).toLocaleDateString('id-ID', { dateStyle: 'medium' })
}

onMounted(async () => {
  try {
    const res = await youtubeApi.getKajianVideos()
    videos.value = res.data ?? []
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Gagal memuat video'
  } finally {
    loading.value = false
  }
})
</script>
