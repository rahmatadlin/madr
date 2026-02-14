import { apiClient } from './client'

export interface DashboardStats {
  total_events: number
  total_banners: number
  total_gallery: number
  total_donations: number
}

export const statsApi = {
  getDashboardStats: async (): Promise<DashboardStats> => {
    const results = await Promise.allSettled([
      apiClient.get('/events?limit=1&offset=0').catch(() => null),
      apiClient.get('/banners?limit=1&offset=0').catch(() => null),
      apiClient.get('/gallery?limit=1&offset=0').catch(() => null),
      apiClient.get('/admin/donations?limit=1&offset=0').catch(() => null),
    ])
    const getTotal = (r: PromiseSettledResult<{ data?: { total?: number } } | null>): number => {
      if (r.status === 'fulfilled' && r.value?.data?.total) return r.value.data.total
      return 0
    }
    return {
      total_events: getTotal(results[0]),
      total_banners: getTotal(results[1]),
      total_gallery: getTotal(results[2]),
      total_donations: getTotal(results[3]),
    }
  },
}
