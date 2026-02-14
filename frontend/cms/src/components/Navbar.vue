<template>
  <header class="flex h-14 items-center border-b px-4 lg:px-6">
    <button
      class="lg:hidden p-2 rounded-md hover:bg-gray-100"
      aria-label="Menu"
      @click="$emit('menu-click')"
    >
      <span class="w-5 h-5 block border border-current rounded"></span>
    </button>
    <span class="ml-2 text-sm text-gray-600">Masjid CMS</span>
  </header>
</template>
<script setup lang="ts">
defineEmits<{ 'menu-click': [] }>()
</script>
