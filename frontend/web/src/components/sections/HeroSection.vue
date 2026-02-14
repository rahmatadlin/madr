<template>
  <section class="relative min-h-screen flex items-center justify-center overflow-hidden">
    <div class="absolute inset-0 z-0">
      <video v-if="heroBanner?.type === 'video' && mediaUrl" autoplay muted loop playsinline class="w-full h-full object-cover">
        <source :src="mediaUrl" type="video/mp4" />
      </video>
      <div v-else class="w-full h-full bg-cover bg-center" :style="bgStyle"></div>
      <div class="absolute inset-0 bg-black/50"></div>
    </div>
    <div class="relative z-10 container mx-auto px-4 text-center text-white">
      <h1 class="text-4xl md:text-6xl lg:text-7xl font-bold mb-6">Masjid Al-Madr</h1>
      <p class="text-lg md:text-xl lg:text-2xl mb-8 text-gray-200">Pusat Kegiatan Keagamaan dan Sosial Masyarakat</p>
      <router-link to="/donate" class="inline-flex items-center justify-center px-8 py-6 text-lg rounded-lg bg-primary text-primary-foreground hover:opacity-90">Donasi Sekarang</router-link>
    </div>
    <div class="absolute bottom-8 left-1/2 -translate-x-1/2 z-10">
      <div class="w-6 h-10 border-2 border-white rounded-full flex items-start justify-center p-2"><div class="w-1 h-3 bg-white rounded-full"></div></div>
    </div>
  </section>
</template>
<script setup lang="ts">
import { computed, ref, onMounted } from 'vue'
import { bannerApi } from '@/api/banners'
import type { PaginatedResponse } from '@/api/client'
import type { Banner } from '@/api/banners'
import { resolveMediaUrl } from '@/utils/media'

const banners = ref<PaginatedResponse<Banner> | null>(null)
const heroBanner = computed(() => banners.value?.data?.[0])
const mediaUrl = computed(() => heroBanner.value?.media_url ? resolveMediaUrl(heroBanner.value.media_url) : undefined)
const bgStyle = computed(() => ({ backgroundImage: heroBanner.value && mediaUrl.value ? 'url(' + mediaUrl.value + ')' : 'linear-gradient(to bottom, #1e3a8a, #3b82f6)' }))
onMounted(async () => { banners.value = await bannerApi.getAll(1, 0) })
</script>
