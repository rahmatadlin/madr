<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold">Kajian Minggu Lalu</h1>
        <p class="text-gray-600">Video kajian dari YouTube Channel (30 hari terakhir)</p>
      </div>
      <UiButton :disabled="syncing" @click="syncFromYouTube">
        {{ syncing ? 'Menyinkronkan...' : 'Sinkronkan dari YouTube' }}
      </UiButton>
    </div>

    <div v-if="loading" class="text-gray-600">Memuat video...</div>
    <div v-else-if="error" class="text-red-600">{{ error }}</div>
    <div v-else-if="kajianList.length === 0" class="text-gray-600">Tidak ada video ditemukan. Klik tombol sinkronkan untuk mengambil video dari YouTube.</div>
    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="kajian in kajianList"
        :key="kajian.id"
        class="rounded-lg border bg-white overflow-hidden shadow-sm hover:shadow-lg transition-shadow"
      >
        <a
          :href="kajian.youtube_url"
          target="_blank"
          rel="noopener noreferrer"
          class="block"
        >
          <div class="relative aspect-video bg-gray-200">
            <img
              :src="kajian.thumbnail_url || 'https://placehold.co/640x360/e5e7eb/6b7280?text=Video'"
              :alt="kajian.title"
              class="w-full h-full object-cover"
            />
            <div class="absolute inset-0 flex items-center justify-center bg-black/30 opacity-0 hover:opacity-100 transition-opacity">
              <span class="rounded-full bg-red-600 p-3 text-white">▶</span>
            </div>
          </div>
          <div class="p-4">
            <p class="text-sm text-gray-500 mb-1">{{ formatDate(kajian.published_at) }}</p>
            <h3 class="font-semibold text-gray-900 mb-2 line-clamp-2">{{ kajian.title }}</h3>
            <p class="text-xs text-gray-600 line-clamp-2">{{ kajian.description }}</p>
          </div>
        </a>
        <div class="px-4 pb-4">
          <button
            class="text-red-600 hover:underline text-sm"
            @click.prevent="confirmDelete(kajian.id)"
          >
            Hapus
          </button>
        </div>
      </div>
    </div>

    <!-- Delete confirmation dialog -->
    <div v-if="deleteId" class="fixed inset-0 z-50 flex items-center justify-center bg-black/50" @click.self="deleteId = null">
      <div class="bg-white rounded-lg shadow-xl p-6 max-w-sm">
        <p class="mb-4">Hapus kajian ini?</p>
        <div class="flex gap-2 justify-end">
          <UiButton variant="outline" @click="deleteId = null">Batal</UiButton>
          <UiButton variant="destructive" @click="doDelete">Hapus</UiButton>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { kajianApi } from '@/api/kajian'
import type { Kajian } from '@/api/kajian'
import UiButton from '@/components/ui/Button.vue'

const kajianList = ref<Kajian[]>([])
const loading = ref(true)
const syncing = ref(false)
const error = ref<string | null>(null)
const deleteId = ref<number | null>(null)

function formatDate(d: string) {
  return new Date(d).toLocaleDateString('id-ID', { dateStyle: 'medium' })
}

async function syncFromYouTube() {
  syncing.value = true
  error.value = null
  try {
    const res = await kajianApi.syncFromYouTube(30)
    alert(`Berhasil menyinkronkan ${res.synced} video dari YouTube`)
    await load()
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Gagal menyinkronkan dari YouTube'
  } finally {
    syncing.value = false
  }
}

function confirmDelete(id: number) {
  deleteId.value = id
}

async function doDelete() {
  if (!deleteId.value) return
  try {
    await kajianApi.delete(deleteId.value)
    deleteId.value = null
    await load()
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Gagal menghapus kajian'
  }
}

async function load() {
  loading.value = true
  error.value = null
  try {
    const res = await kajianApi.getAll(50, 0)
    kajianList.value = res.data ?? []
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Gagal memuat kajian'
  } finally {
    loading.value = false
  }
}

onMounted(load)
</script>
