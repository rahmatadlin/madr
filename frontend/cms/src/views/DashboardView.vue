<template>
  <div class="space-y-6">
    <div>
      <h1 class="text-3xl font-bold">Dashboard</h1>
      <p class="text-gray-600">Welcome to Masjid CMS</p>
    </div>
    <div class="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
      <div v-for="stat in statCards" :key="stat.title" class="rounded-lg border bg-white p-6">
        <div class="flex items-center justify-between pb-2">
          <h3 class="text-sm font-medium">{{ stat.title }}</h3>
        </div>
        <div v-if="loading" class="h-8 w-20 bg-gray-200 rounded animate-pulse"></div>
        <template v-else>
          <div class="text-2xl font-bold">{{ stat.value }}</div>
          <p class="text-xs text-gray-500">{{ stat.description }}</p>
        </template>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { statsApi } from '@/api/stats'
import type { DashboardStats } from '@/api/stats'

const stats = ref<DashboardStats | null>(null)
const loading = ref(true)

const statCards = computed(() => [
  { title: 'Total Events', value: stats.value?.total_events ?? 0, description: 'Active events' },
  { title: 'Total Banners', value: stats.value?.total_banners ?? 0, description: 'Banner items' },
  { title: 'Gallery Images', value: stats.value?.total_gallery ?? 0, description: 'Total images' },
  { title: 'Total Donations', value: stats.value?.total_donations ?? 0, description: 'All donations' },
])

onMounted(async () => {
  try { stats.value = await statsApi.getDashboardStats() } finally { loading.value = false }
})
</script>
