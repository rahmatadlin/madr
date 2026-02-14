<template>
  <section class="py-20 bg-gradient-to-br from-blue-50 to-indigo-100">
    <div class="container mx-auto px-4">
      <div class="text-center mb-12">
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">Total Donasi Terkumpul</h2>
        <p class="text-lg text-gray-600">Mari bersama-sama membangun masjid yang lebih baik</p>
      </div>
      <div v-if="isLoading" class="text-center text-gray-600">Memuat data donasi...</div>
      <template v-else>
        <div class="text-center mb-12">
          <div class="inline-block bg-white rounded-2xl shadow-xl p-8 md:p-12">
            <p class="text-sm md:text-base text-gray-600 mb-2">Total Donasi</p>
            <div class="text-4xl md:text-6xl font-bold text-blue-600 mb-2">{{ formatCurrency(summary ? summary.total_amount : 0) }}</div>
            <p class="text-sm text-gray-500">dari {{ summary ? summary.total_transactions : 0 }} transaksi</p>
          </div>
        </div>
        <div v-if="summary && summary.per_category && summary.per_category.length" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12">
          <div v-for="cat in summary.per_category" :key="cat.category_id" class="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow">
            <h3 class="text-lg font-semibold text-gray-900 mb-2">{{ cat.category }}</h3>
            <p class="text-2xl font-bold text-blue-600">{{ formatCurrency(cat.amount) }}</p>
          </div>
        </div>
        <div class="text-center">
          <router-link to="/donate" class="inline-flex items-center justify-center px-8 py-6 text-lg rounded-lg bg-blue-600 hover:bg-blue-700 text-white">Donasi Sekarang</router-link>
        </div>
      </template>
    </div>
  </section>
</template>
<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { donationApi } from '@/api/donations'
import type { DonationSummary } from '@/api/donations'
const summary = ref<DonationSummary | null>(null)
const isLoading = ref(true)
function formatCurrency(amount: number) {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0, maximumFractionDigits: 0 }).format(amount)
}
onMounted(async () => {
  try { summary.value = await donationApi.getSummary() } finally { isLoading.value = false }
})
</script>
