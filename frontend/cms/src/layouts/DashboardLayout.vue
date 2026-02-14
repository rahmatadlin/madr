<template>
  <div class="flex h-screen overflow-hidden">
    <Sidebar :open="sidebarOpen" @close="sidebarOpen = false" />
    <div class="flex-1 flex flex-col overflow-hidden lg:pl-64">
      <Navbar @menu-click="sidebarOpen = !sidebarOpen" />
      <main class="flex-1 overflow-y-auto p-4 lg:p-6">
        <router-view />
      </main>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import Sidebar from '@/components/Sidebar.vue'
import Navbar from '@/components/Navbar.vue'
const sidebarOpen = ref(false)
</script>
