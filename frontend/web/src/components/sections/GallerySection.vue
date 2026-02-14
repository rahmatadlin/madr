<template>
  <section v-if="items.length > 0 || isLoading" class="py-20 bg-white">
    <div class="container mx-auto px-4">
      <div class="text-center mb-12">
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">Galeri Foto</h2>
        <p class="text-lg text-gray-600">Dokumentasi kegiatan dan aktivitas masjid</p>
      </div>
      <p v-if="isLoading" class="text-center text-gray-600">Memuat galeri...</p>
      <div v-else class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
        <div v-for="item in items" :key="item.id" class="relative aspect-square rounded-lg overflow-hidden group">
          <img :src="imgUrl(item.image_url)" :alt="item.title" class="w-full h-full object-cover transition-transform duration-300 group-hover:scale-110" />
          <div class="absolute inset-0 bg-black/0 group-hover:bg-black/50 transition-colors duration-300 flex items-center justify-center">
            <p class="text-white opacity-0 group-hover:opacity-100 transition-opacity text-sm font-medium px-4 text-center">{{ item.title }}</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { galleryApi } from '@/api/gallery'
import type { GalleryItem } from '@/api/gallery'
import { resolveMediaUrl } from '@/utils/media'
const items = ref<GalleryItem[]>([])
const isLoading = ref(true)
function imgUrl(url: string) {
  return url.startsWith('http') ? url : resolveMediaUrl(url)
}
onMounted(async () => {
  try { const res = await galleryApi.getAll(12, 0); items.value = res.data || [] } finally { isLoading.value = false }
})
</script>
