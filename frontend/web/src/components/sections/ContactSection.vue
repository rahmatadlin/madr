<template>
  <section class="py-20 bg-gray-50">
    <div class="container mx-auto px-4">
      <div class="text-center mb-12">
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">Hubungi Kami</h2>
        <p class="text-lg text-gray-600">Ada pertanyaan atau saran? Silakan hubungi kami</p>
      </div>
      <div class="max-w-2xl mx-auto">
        <form class="bg-white rounded-lg shadow-lg p-6 md:p-8 space-y-6" @submit.prevent="onSubmit">
          <input v-model="form.honeypot" type="text" class="hidden" tabindex="-1" autocomplete="off" />
          <div>
            <label for="name" class="block text-sm font-medium text-gray-700 mb-1">Nama</label>
            <input id="name" v-model="form.name" type="text" placeholder="Masukkan nama Anda" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500" />
            <p v-if="errors.name" class="text-sm text-red-500 mt-1">{{ errors.name }}</p>
          </div>
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700 mb-1">Email</label>
            <input id="email" v-model="form.email" type="email" placeholder="nama@example.com" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500" />
            <p v-if="errors.email" class="text-sm text-red-500 mt-1">{{ errors.email }}</p>
          </div>
          <div>
            <label for="subject" class="block text-sm font-medium text-gray-700 mb-1">Subjek</label>
            <input id="subject" v-model="form.subject" type="text" placeholder="Subjek pesan" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500" />
            <p v-if="errors.subject" class="text-sm text-red-500 mt-1">{{ errors.subject }}</p>
          </div>
          <div>
            <label for="message" class="block text-sm font-medium text-gray-700 mb-1">Pesan</label>
            <textarea id="message" v-model="form.message" rows="5" placeholder="Tuliskan pesan Anda..." class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500" />
            <p v-if="errors.message" class="text-sm text-red-500 mt-1">{{ errors.message }}</p>
          </div>
          <div v-if="submitStatus === 'success'" class="bg-green-50 border border-green-200 text-green-800 px-4 py-3 rounded">Pesan berhasil dikirim!</div>
          <div v-if="submitStatus === 'error'" class="bg-red-50 border border-red-200 text-red-800 px-4 py-3 rounded">Terjadi kesalahan. Silakan coba lagi.</div>
          <button type="submit" class="w-full py-2 rounded-lg bg-primary text-primary-foreground hover:opacity-90 disabled:opacity-50" :disabled="isSubmitting">{{ isSubmitting ? 'Mengirim...' : 'Kirim Pesan' }}</button>
        </form>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { contactApi } from '@/api/contact'

const form = reactive({ name: '', email: '', subject: '', message: '', honeypot: '' })
const errors = reactive<Record<string, string>>({})
const isSubmitting = ref(false)
const submitStatus = ref<'success' | 'error' | null>(null)

function validate() {
  Object.keys(errors).forEach((k) => delete (errors as Record<string, unknown>)[k])
  if (form.name.length < 3) errors.name = 'Nama minimal 3 karakter'
  if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email)) errors.email = 'Email tidak valid'
  if (form.subject.length < 5) errors.subject = 'Subjek minimal 5 karakter'
  if (form.message.length < 10) errors.message = 'Pesan minimal 10 karakter'
  return Object.keys(errors).length === 0
}

async function onSubmit() {
  if (form.honeypot) return
  if (!validate()) return
  isSubmitting.value = true
  submitStatus.value = null
  try {
    await contactApi.submit({ name: form.name, email: form.email, subject: form.subject, message: form.message })
    submitStatus.value = 'success'
    form.name = ''; form.email = ''; form.subject = ''; form.message = ''
  } catch {
    submitStatus.value = 'error'
  } finally {
    isSubmitting.value = false
  }
}
</script>
