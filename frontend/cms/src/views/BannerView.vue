<template>
  <div class="space-y-6">
    <div>
      <h1 class="text-3xl font-bold">Banner</h1>
      <p class="text-gray-600">Kelola satu banner utama untuk Hero (gambar atau video).</p>
    </div>
    <div v-if="loading" class="space-y-2">
      <div v-for="i in 3" :key="i" class="h-12 w-full bg-gray-200 rounded animate-pulse"></div>
      <div class="h-48 w-full bg-gray-200 rounded animate-pulse"></div>
    </div>
    <div v-else class="space-y-4 rounded-lg border p-4">
      <div>
        <UiLabel for-id="banner-title">Judul</UiLabel>
        <UiInput id="banner-title" v-model="title" placeholder="Judul banner" class="mt-1" />
      </div>
      <div>
        <UiLabel>Jenis Media</UiLabel>
        <UiSelect v-model="mediaType" placeholder="Pilih jenis" :options="mediaTypeOptions" class="mt-1" />
      </div>
      <div>
        <UiLabel>File</UiLabel>
        <input type="file" :accept="mediaType === 'video' ? 'video/mp4' : 'image/*'" class="mt-1 w-full text-sm" @change="onFileChange" />
        <div v-if="previewUrl && mediaType === 'image'" class="mt-2 h-48 rounded-lg border overflow-hidden">
          <img :src="previewUrl" alt="Preview" class="h-full w-full object-cover" />
        </div>
        <div v-if="previewUrl && mediaType === 'video'" class="mt-2 rounded-lg border p-3">
          <video :src="previewUrl" controls class="w-full rounded"></video>
        </div>
      </div>
      <UiButton :disabled="saving || !title.trim() || (!bannerId && !file)" @click="save">
        {{ saving ? 'Menyimpan...' : 'Simpan' }}
      </UiButton>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { bannerApi } from '@/api/banners'
import type { Banner } from '@/api/banners'
import { resolveMediaUrl } from '@/utils/media'
import UiLabel from '@/components/ui/Label.vue'
import UiInput from '@/components/ui/Input.vue'
import UiSelect from '@/components/ui/Select.vue'
import UiButton from '@/components/ui/Button.vue'

const mediaTypeOptions = [
  { value: 'image', label: 'Image' },
  { value: 'video', label: 'Video (mp4)' },
]

const banner = ref<Banner | null>(null)
const loading = ref(true)
const saving = ref(false)
const title = ref('')
const mediaType = ref<'image' | 'video'>('image')
const file = ref<File | null>(null)

const bannerId = computed(() => banner.value?.id)
const previewUrl = computed(() => {
  if (file.value) return URL.createObjectURL(file.value)
  return resolveMediaUrl(banner.value?.media_url)
})

function onFileChange(e: Event) {
  const target = e.target as HTMLInputElement
  file.value = target.files?.[0] ?? null
}

async function save() {
  saving.value = true
  try {
    if (bannerId.value) {
      await bannerApi.update(bannerId.value, { title: title.value, type: mediaType.value, file: file.value ?? undefined })
    } else {
      await bannerApi.create({ title: title.value, type: mediaType.value, file: file.value ?? undefined })
    }
    file.value = null
    await load()
  } finally {
    saving.value = false
  }
}

async function load() {
  loading.value = true
  try {
    const res = await bannerApi.getAll(1, 0)
    banner.value = res.data?.[0] ?? null
    title.value = banner.value?.title ?? ''
    mediaType.value = (banner.value?.type as 'image' | 'video') ?? 'image'
  } finally {
    loading.value = false
  }
}

onMounted(load)
</script>
