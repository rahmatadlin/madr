<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold">Events</h1>
        <p class="text-gray-600">Manage events and activities</p>
      </div>
      <button class="px-4 py-2 rounded-lg bg-primary text-primary-foreground hover:opacity-90" @click="openModal()">Add Event</button>
    </div>
    <input v-model="search" placeholder="Search events..." class="w-full max-w-sm px-3 py-2 border rounded-lg" />
    <div v-if="loading" class="text-gray-600">Loading...</div>
    <div v-else class="rounded-lg border overflow-hidden">
      <table class="w-full text-sm">
        <thead class="bg-gray-50 border-b">
          <tr>
            <th class="text-left p-3">Title</th>
            <th class="text-left p-3">Date</th>
            <th class="text-left p-3">Location</th>
            <th class="p-3"></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="event in filteredEvents" :key="event.id" class="border-b last:border-0">
            <td class="p-3">{{ event.title }}</td>
            <td class="p-3">{{ formatDate(event.date) }}</td>
            <td class="p-3">{{ event.location || '-' }}</td>
            <td class="p-3">
              <button class="text-blue-600 hover:underline mr-2" @click="openModal(event)">Edit</button>
              <button class="text-red-600 hover:underline" @click="confirmDelete(event.id)">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div v-if="showModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black/50" @click.self="showModal = false">
      <div class="bg-white rounded-lg shadow-xl p-6 w-full max-w-md">
        <h3 class="text-lg font-bold mb-4">{{ editingId ? 'Edit Event' : 'Add Event' }}</h3>
        <form @submit.prevent="submitEvent" class="space-y-3">
          <div>
            <label class="block text-sm font-medium mb-1">Title</label>
            <input v-model="editForm.title" class="w-full px-3 py-2 border rounded-lg" required />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">Description</label>
            <textarea v-model="editForm.description" rows="3" class="w-full px-3 py-2 border rounded-lg"></textarea>
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">Date</label>
            <input v-model="editForm.date" type="datetime-local" class="w-full px-3 py-2 border rounded-lg" required />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">Location</label>
            <input v-model="editForm.location" class="w-full px-3 py-2 border rounded-lg" />
          </div>
          <div class="flex gap-2 justify-end pt-2">
            <button type="button" class="px-4 py-2 border rounded-lg" @click="showModal = false">Cancel</button>
            <button type="submit" class="px-4 py-2 rounded-lg bg-primary text-primary-foreground">{{ saving ? 'Saving...' : 'Save' }}</button>
          </div>
        </form>
      </div>
    </div>
    <div v-if="deleteId" class="fixed inset-0 z-50 flex items-center justify-center bg-black/50" @click.self="deleteId = null">
      <div class="bg-white rounded-lg shadow-xl p-6 max-w-sm">
        <p class="mb-4">Delete this event?</p>
        <div class="flex gap-2 justify-end">
          <button class="px-4 py-2 border rounded-lg" @click="deleteId = null">Cancel</button>
          <button class="px-4 py-2 rounded-lg bg-red-500 text-white" @click="doDelete">Delete</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, computed, reactive, onMounted } from 'vue'
import { eventApi } from '@/api/events'
import type { Event, CreateEventRequest } from '@/api/events'

const events = ref<Event[]>([])
const loading = ref(true)
const search = ref('')
const showModal = ref(false)
const editingId = ref<number | null>(null)
const saving = ref(false)
const deleteId = ref<number | null>(null)

const editForm = reactive<CreateEventRequest>({ title: '', description: '', date: '', location: '' })

const filteredEvents = computed(() => {
  const q = search.value.toLowerCase()
  return events.value.filter((e) => e.title.toLowerCase().includes(q))
})

function formatDate(d: string) {
  return new Date(d).toLocaleDateString('id-ID', { dateStyle: 'medium' })
}

function openModal(event?: Event) {
  editingId.value = event?.id ?? null
  editForm.title = event?.title ?? ''
  editForm.description = event?.description ?? ''
  editForm.date = event?.date ? new Date(event.date).toISOString().slice(0, 16) : ''
  editForm.location = event?.location ?? ''
  showModal.value = true
}

async function submitEvent() {
  saving.value = true
  try {
    const payload = { ...editForm, date: new Date(editForm.date).toISOString() }
    if (editingId.value) await eventApi.update(editingId.value, payload)
    else await eventApi.create(payload)
    showModal.value = false
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
  await eventApi.delete(deleteId.value)
  deleteId.value = null
  await load()
}

async function load() {
  loading.value = true
  try {
    const res = await eventApi.getAll(100, 0)
    events.value = res.data ?? []
  } finally {
    loading.value = false
  }
}

onMounted(load)
</script>
