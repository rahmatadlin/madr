<template>
  <main class="min-h-screen">
    <div class="container mx-auto px-4 py-20">
      <div class="text-center mb-12 animate-fade-in">
        <h1 class="text-4xl md:text-5xl font-bold text-gray-900 mb-4">
          Donasi untuk Masjid Al-Madr
        </h1>
        <p class="text-lg text-gray-600 max-w-2xl mx-auto">
          Bantu kami membangun dan mengembangkan masjid untuk kemaslahatan umat.
        </p>
      </div>

      <div v-if="isLoading" class="text-center py-12 text-gray-600">
        Memuat data donasi...
      </div>
      <div
        v-else
        class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12"
      >
        <div
          v-for="category in summary?.per_category ?? []"
          :key="category.category_id"
          class="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow"
        >
          <h3 class="text-lg font-semibold text-gray-900 mb-2">
            {{ category.category }}
          </h3>
          <p class="text-3xl font-bold text-blue-600 mb-4">
            {{ formatCurrency(category.amount) }}
          </p>
          <a
            href="https://wa.me/62123456789"
            target="_blank"
            rel="noopener noreferrer"
            class="block w-full text-center py-2 rounded-lg bg-primary text-primary-foreground hover:opacity-90"
          >
            Donasi Sekarang
          </a>
        </div>
      </div>

      <div class="max-w-2xl mx-auto bg-blue-50 rounded-lg p-8 text-center">
        <h2 class="text-2xl font-bold text-gray-900 mb-4">Cara Donasi</h2>
        <p class="text-gray-600 mb-6">
          Untuk donasi, hubungi kami via WhatsApp atau email.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <a
            href="https://wa.me/62123456789"
            target="_blank"
            rel="noopener noreferrer"
            class="inline-flex justify-center items-center px-6 py-3 rounded-lg bg-green-500 hover:bg-green-600 text-white"
          >
            Hubungi via WhatsApp
          </a>
          <a
            href="mailto:donasi@masjidalmadr.com"
            class="inline-flex justify-center items-center px-6 py-3 rounded-lg border border-gray-300 hover:bg-gray-50"
          >
            Kirim Email
          </a>
        </div>
      </div>
    </div>
    <Footer />
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { donationApi } from '@/api/donations'
import type { DonationSummary } from '@/api/donations'
import Footer from '@/components/layouts/Footer.vue'

const summary = ref<DonationSummary | null>(null)
const isLoading = ref(true)

function formatCurrency(amount: number) {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  }).format(amount)
}

onMounted(async () => {
  try {
    summary.value = await donationApi.getSummary()
  } finally {
    isLoading.value = false
  }
})
</script>
