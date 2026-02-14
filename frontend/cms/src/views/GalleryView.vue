<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold">Gallery</h1>
        <p class="text-gray-600">Manage gallery images</p>
      </div>
      <button class="px-4 py-2 rounded-lg bg-primary text-primary-foreground hover:opacity-90" @click="showAdd = true">Add Image</button>
    </div>
    <div v-if="loading" class="text-gray-600">Loading...</div>
    <div v-else class="grid grid-cols-2 md:grid-cols-4 gap-4">
      <div v-for="item in items" :key="item.id" class="rounded-lg border overflow-hidden">
        <img :src="resolveUrl(item.image_url)" :alt="item.title" class="w-full aspect-square object-cover" />
        <div class="p-2 flex justify-between items-center">
          <span class="text-sm truncate">{{ item.title }}</span>
          <button class="text-red-600 text-sm hover:underline" @click="confirmDelete(item.id)">Delete</button>
        </div>
      </div>
    </div>
    <div v-if="showAdd" class="fixed inset-0 z-50 flex items-center justify-center bg-black/50" @click.self="showAdd = false">
      <div class="bg-white rounded-lg shadow-xl p-6 w-full max-w-md">
        <h3 class="text-lg font-bold mb-4">Add Image</h3>
        <form @submit.prevent="addImage" class="space-y-3">
          <div>
            <label class="block text-sm font-medium mb-1">Title</label>
            <input v-model="addTitle" class="w-full px-3 py-2 border rounded-lg" required />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">File</label>
            <input ref="addFileInput" type="file" accept="image/*" class="w-full" @change="onAddFile" />
          </div>
          <div class="flex gap-2 justify-end pt-2">
            <button type="button" class="px-4 py-2 border rounded-lg" @click="showAdd = false">Cancel</button>
            <button type="submit" class="px-4 py-2 rounded-lg bg-primary text-primary-foreground" :disabled="saving">{{ saving ? 'Saving...' : 'Save' }}</button>
          </div>
        </form>
      </div>
    </div>
    <div v-if="deleteId" class="fixed inset-0 z-50 flex items-center justify-center bg-black/50" @click.self="deleteId = null">
      <div class="bg-white rounded-lg shadow-xl p-6 max-w-sm">
        <p class="mb-4">Delete this image?</p>
        <div class="flex gap-2 justify-end">
          <button class="px-4 py-2 border rounded-lg" @click="deleteId = null">Cancel</button>
          <button class="px-4 py-2 rounded-lg bg-red-500 text-white" @click="doDelete">Delete</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { galleryApi } from '@/api/gallery'
import type { GalleryItem } from '@/api/gallery'
import { resolveMediaUrl } from '@/utils/media'

const items = ref<GalleryItem[]>([])
const loading = ref(true)
const showAdd = ref(false)
const addTitle = ref('')
const addFile = ref<File | null>(null)
const addFileInput = ref<HTMLInputElement | null>(null)
const saving = ref(false)
const deleteId = ref<number | null>(null)

function resolveUrl(url: string) {
  return url.startsWith('http') ? url : resolveMediaUrl(url) || url
}

function onAddFile(e: Event) {
  addFile.value = (e.target as HTMLInputElement).files?.[0] ?? null
}

async function addImage() {
  if (!addTitle.value.trim() || !addFile.value) return
  saving.value = true
  try {
    await galleryApi.create({ title: addTitle.value, file: addFile.value })
    showAdd.value = false
    addTitle.value = ''
    addFile.value = null
    if (addFileInput.value) addFileInput.value.value = ''
    await load()
  } finally {
    saving.value = false
  }
}

function confirmDelete(id: number) {
  deleteId.value = id
}

async function doDelete() {
  if (!deleteId.value) return
  await galleryApi.delete(deleteId.value)
  deleteId.value = null
  await load()
}

async function load() {
  loading.value = true
  try {
    const res = await galleryApi.getAll(100, 0)
    items.value = res.data ?? []
  } finally {
    loading.value = false
  }
}

onMounted(load)
</script>
