<template>
  <div v-if="open" class="fixed inset-0 bg-black/50 z-40 lg:hidden" @click="$emit('close')"></div>
  <aside
      class="fixed top-0 left-0 z-50 h-full w-64 bg-white border-r transition-transform duration-300 lg:translate-x-0"
      :class="open ? 'translate-x-0' : '-translate-x-full'"
    >
      <div class="flex h-full flex-col">
        <div class="flex h-16 items-center border-b px-6">
          <h1 class="text-xl font-bold">Masjid CMS</h1>
        </div>
        <nav class="flex-1 space-y-1 p-4">
          <router-link
            v-for="item in menuItems"
            :key="item.href"
            :to="item.href"
            class="flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors"
            :class="isActive(item.href) ? 'bg-primary text-primary-foreground' : 'text-gray-600 hover:bg-gray-100'"
            @click="onLinkClick"
          >
            {{ item.title }}
          </router-link>
        </nav>
        <div class="border-t p-4">
          <button
            class="flex w-full items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium text-gray-600 hover:bg-gray-100"
            @click="logout"
          >
            Logout
          </button>
        </div>
      </div>
  </aside>
</template>
<script setup lang="ts">
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

defineProps<{ open: boolean }>()
const emit = defineEmits<{ close: [] }>()

const route = useRoute()
const auth = useAuthStore()

const menuItems = [
  { title: 'Dashboard', href: '/' },
  { title: 'Events', href: '/events' },
  { title: 'Gallery', href: '/gallery' },
  { title: 'Banner', href: '/banner' },
  { title: 'About', href: '/about' },
  { title: 'Donations', href: '/donations' },
]

function isActive(href: string) {
  if (href === '/') return route.path === '/'
  return route.path.startsWith(href)
}

function onLinkClick() {
  if (window.innerWidth < 1024) emit('close')
}

function logout() {
  auth.logout()
  window.location.href = '/login'
}
</script>
