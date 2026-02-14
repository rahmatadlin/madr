import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const routes: RouteRecordRaw[] = [
  { path: '/login', name: 'login', component: () => import('@/views/LoginView.vue'), meta: { public: true } },
  {
    path: '/',
    component: () => import('@/layouts/DashboardLayout.vue'),
    meta: { requiresAuth: true },
    children: [
      { path: '', name: 'dashboard', component: () => import('@/views/DashboardView.vue') },
      { path: 'banner', name: 'banner', component: () => import('@/views/BannerView.vue') },
      { path: 'about', name: 'about', component: () => import('@/views/AboutView.vue') },
      { path: 'events', name: 'events', component: () => import('@/views/EventsView.vue') },
      { path: 'gallery', name: 'gallery', component: () => import('@/views/GalleryView.vue') },
      { path: 'donations', name: 'donations', component: () => import('@/views/DonationsView.vue') },
    ],
  },
]

const router = createRouter({ history: createWebHistory(), routes })

router.beforeEach((to, _from, next) => {
  const auth = useAuthStore()
  if (to.meta.requiresAuth && !auth.isLoggedIn) return next('/login')
  if (to.path === '/login' && auth.isLoggedIn) return next('/')
  next()
})

export default router
