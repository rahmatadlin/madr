<template>
  <div class="space-y-6">
    <div>
      <h1 class="text-3xl font-bold">Donations</h1>
      <p class="text-gray-600">View donation records</p>
    </div>
    <div v-if="loading" class="text-gray-600">Loading...</div>
    <div v-else class="rounded-lg border overflow-hidden">
      <table class="w-full text-sm">
        <thead class="bg-gray-50 border-b">
          <tr>
            <th class="text-left p-3">Donor</th>
            <th class="text-left p-3">Amount</th>
            <th class="text-left p-3">Status</th>
            <th class="text-left p-3">Date</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="d in donations" :key="d.id" class="border-b last:border-0">
            <td class="p-3">{{ d.donor_name || '-' }}</td>
            <td class="p-3">{{ formatCurrency(d.amount) }}</td>
            <td class="p-3">{{ d.payment_status }}</td>
            <td class="p-3">{{ formatDate(d.created_at) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { donationApi } from '@/api/donations'
import type { Donation } from '@/api/donations'

const donations = ref<Donation[]>([])
const loading = ref(true)

function formatCurrency(amount: number) {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(amount)
}

function formatDate(d: string) {
  return new Date(d).toLocaleDateString('id-ID', { dateStyle: 'medium' })
}

onMounted(async () => {
  loading.value = true
  try {
    const res = await donationApi.getAll(100, 0, 'success')
    donations.value = res.data ?? []
  } finally {
    loading.value = false
  }
})
</script>
