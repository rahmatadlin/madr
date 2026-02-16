<template>
  <header class="flex h-14 items-center border-b px-4 lg:px-6">
    <button
      class="lg:hidden p-2 rounded-md hover:bg-gray-100"
      aria-label="Menu"
      @click="$emit('menu-click')"
    >
      <Bars3Icon class="h-6 w-6" />
    </button>
    <span class="ml-2 text-sm text-gray-600">Masjid CMS</span>
  </header>
</template>
<script setup lang="ts">
import { Bars3Icon } from '@heroicons/vue/24/outline'
defineEmits<{ 'menu-click': [] }>()
</script>
