<template>
  <div
    class="flex min-h-screen items-center justify-center p-4 relative"
    style="background-size: cover; background-position: center; background-image: url('/images/login-background.jpg');"
  >
    <div class="absolute inset-0 bg-black/40"></div>
    <div class="w-full max-w-md relative z-10 bg-white/95 backdrop-blur rounded-lg shadow-lg p-6">
      <h2 class="text-2xl font-bold mb-1">Masjid CMS</h2>
      <p class="text-sm text-gray-600 mb-6">Sign in to your account</p>
      <form @submit.prevent="onSubmit" class="space-y-4">
        <div>
          <label for="username" class="block text-sm font-medium text-gray-700 mb-1">Username</label>
          <input
            id="username"
            v-model="username"
            type="text"
            placeholder="Enter your username"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            :disabled="loading"
          />
          <p v-if="errors.username" class="text-sm text-red-500 mt-1">{{ errors.username }}</p>
        </div>
        <div>
          <label for="password" class="block text-sm font-medium text-gray-700 mb-1">Password</label>
          <input
            id="password"
            v-model="password"
            type="password"
            placeholder="Enter your password"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            :disabled="loading"
          />
          <p v-if="errors.password" class="text-sm text-red-500 mt-1">{{ errors.password }}</p>
        </div>
        <p v-if="errorMsg" class="text-sm text-red-500">{{ errorMsg }}</p>
        <button
          type="submit"
          class="w-full py-2 rounded-lg bg-primary text-primary-foreground hover:opacity-90 disabled:opacity-50"
          :disabled="loading"
        >
          {{ loading ? 'Signing in...' : 'Sign In' }}
        </button>
      </form>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const auth = useAuthStore()
const username = ref('')
const password = ref('')
const loading = ref(false)
const errorMsg = ref('')
const errors = reactive<Record<string, string>>({})

async function onSubmit() {
  errorMsg.value = ''
  errors.username = ''
  errors.password = ''
  if (!username.value.trim()) { errors.username = 'Username is required'; return }
  if (!password.value) { errors.password = 'Password is required'; return }
  loading.value = true
  try {
    await auth.login(username.value, password.value)
    router.push('/')
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : 'Login failed'
  } finally {
    loading.value = false
  }
}
</script>
