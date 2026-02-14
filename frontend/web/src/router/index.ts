import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import MainLayout from '@/layouts/MainLayout.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: MainLayout,
    children: [
      { path: '', name: 'home', component: () => import('@/views/HomeView.vue') },
      { path: 'donate', name: 'donate', component: () => import('@/views/DonateView.vue') },
      { path: 'contact', name: 'contact', component: () => import('@/views/ContactView.vue') },
      { path: 'events/:id', name: 'event', component: () => import('@/views/EventDetailView.vue') },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
