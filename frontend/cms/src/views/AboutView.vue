<template>
  <div class="space-y-6">
    <div>
      <h1 class="text-3xl font-bold">About</h1>
      <p class="text-gray-600">Kelola konten About untuk ditampilkan di website publik</p>
    </div>
    <div v-if="loading" class="space-y-3">
      <div class="h-10 w-full bg-gray-200 rounded animate-pulse"></div>
      <div class="h-24 w-full bg-gray-200 rounded animate-pulse"></div>
    </div>
    <form v-else class="space-y-6" @submit.prevent="save">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium mb-1">Judul *</label>
            <input v-model="form.title" class="w-full px-3 py-2 border rounded-lg" required />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">Subjudul</label>
            <input v-model="form.subtitle" class="w-full px-3 py-2 border rounded-lg" />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">Deskripsi</label>
            <textarea v-model="form.description" rows="4" class="w-full px-3 py-2 border rounded-lg"></textarea>
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">Deskripsi Tambahan</label>
            <textarea v-model="form.additional_description" rows="4" class="w-full px-3 py-2 border rounded-lg"></textarea>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium mb-1">Tahun Berdiri</label>
              <input v-model.number="form.years_active" type="number" min="0" class="w-full px-3 py-2 border rounded-lg" />
            </div>
            <div>
              <label class="block text-sm font-medium mb-1">Jamaah Aktif</label>
              <input v-model.number="form.active_members" type="number" min="0" class="w-full px-3 py-2 border rounded-lg" />
            </div>
          </div>
        </div>
        <div>
          <label class="block text-sm font-medium mb-1">Gambar (maks 3)</label>
          <input ref="fileInput" type="file" accept="image/*" multiple class="hidden" @change="onUpload" />
          <div class="grid grid-cols-3 gap-3">
            <template v-for="url in uploadedImages" :key="url">
              <div class="relative h-28 rounded-md border overflow-hidden">
                <img :src="url" alt="Preview" class="w-full h-full object-cover" />
                <button type="button" class="absolute top-1 right-1 bg-red-500 text-white rounded-full w-6 h-6 text-xs" @click="removeImage(url)">×</button>
              </div>
            </template>
            <template v-for="key in emptySlotKeys" :key="key">
              <button type="button" :data-slot="key" class="h-28 rounded-md border border-dashed flex items-center justify-center text-gray-500 hover:border-primary" @click="fileInput?.click()">+</button>
            </template>
          </div>
        </div>
      </div>
      <button type="submit" class="px-4 py-2 rounded-lg bg-primary text-primary-foreground hover:opacity-90 disabled:opacity-50" :disabled="saving">{{ saving ? 'Menyimpan...' : 'Simpan' }}</button>
    </form>
  </div>
</template>
<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { aboutApi } from '@/api/about'
import type { UpdateAboutRequest } from '@/api/about'
import { apiClient } from '@/api/client'

const loading = ref(true)
const saving = ref(false)
const fileInput = ref<HTMLInputElement | null>(null)
const uploadedImages = ref<string[]>([])
const emptySlotKeys = computed(() => Array.from({ length: Math.max(0, 3 - uploadedImages.value.length) }, (_, i) => 'empty-' + i))
const form = reactive<UpdateAboutRequest>({
  title: '',
  subtitle: '',
  description: '',
  additional_description: '',
  image_url: '',
  years_active: 0,
  active_members: 0,
})

async function onUpload(e: Event) {
  const files = (e.target as HTMLInputElement).files
  if (!files?.length) return
  const current = uploadedImages.value.length
  if (current >= 3) return
  const toAdd = Math.min(3 - current, files.length)
  for (let i = 0; i < toAdd; i++) {
    const formData = new FormData()
    formData.append('file', files[i])
    const res = await apiClient.post('/admin/upload', formData, { headers: { 'Content-Type': 'multipart/form-data' } })
    const url = (res.data?.data as { url?: string })?.url ?? (res.data?.data as { public_url?: string })?.public_url
    if (url) uploadedImages.value.push(url)
  }
  if (fileInput.value) fileInput.value.value = ''
}

function removeImage(url: string) {
  uploadedImages.value = uploadedImages.value.filter((u) => u !== url)
}

async function save() {
  saving.value = true
  try {
    await aboutApi.update({
      ...form,
      image_url: uploadedImages.value.length ? JSON.stringify(uploadedImages.value) : '',
    })
  } finally {
    saving.value = false
  }
}

onMounted(async () => {
  try {
    const data = await aboutApi.get()
    form.title = data.title ?? ''
    form.subtitle = data.subtitle ?? ''
    form.description = data.description ?? ''
    form.additional_description = data.additional_description ?? ''
    form.years_active = data.years_active ?? 0
    form.active_members = data.active_members ?? 0
    if (data.image_url) {
      try {
        const p = JSON.parse(data.image_url)
        uploadedImages.value = Array.isArray(p) ? p.filter((x) => typeof x === 'string') : typeof p === 'string' ? [p] : []
      } catch {
        if (data.image_url.trim()) uploadedImages.value = [data.image_url]
      }
    }
  } finally {
    loading.value = false
  }
})
</script>
