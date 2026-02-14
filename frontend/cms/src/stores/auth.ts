import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authApi } from '@/api/auth'
import type { LoginResponse } from '@/api/auth'
import { setAuthToken } from '@/api/client'

const TOKEN_KEY = 'madr_cms_token'
const USER_KEY = 'madr_cms_user'

function getStoredToken(): string | null {
  return localStorage.getItem(TOKEN_KEY)
}

function getStoredUser(): LoginResponse['user'] | null {
  try {
    const raw = localStorage.getItem(USER_KEY)
    return raw ? JSON.parse(raw) : null
  } catch {
    return null
  }
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(getStoredToken())
  const user = ref<LoginResponse['user'] | null>(getStoredUser())

  const isLoggedIn = computed(() => !!token.value)

  function setSession(t: string, u: LoginResponse['user']) {
    token.value = t
    user.value = u
    localStorage.setItem(TOKEN_KEY, t)
    localStorage.setItem(USER_KEY, JSON.stringify(u))
    setAuthToken(t)
  }

  function logout() {
    token.value = null
    user.value = null
    localStorage.removeItem(TOKEN_KEY)
    localStorage.removeItem(USER_KEY)
    setAuthToken(null)
  }

  async function login(username: string, password: string) {
    const res = await authApi.login({ username, password })
    setSession(res.access_token, res.user)
    return res
  }

  if (token.value) setAuthToken(token.value)

  return { token, user, isLoggedIn, login, logout, setSession }
})
