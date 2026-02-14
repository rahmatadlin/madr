import { apiClient, type PaginatedResponse } from './client'

export interface Banner {
  id: number
  title: string
  media_url: string
  type: 'image' | 'video'
  created_at: string
  updated_at: string
}

export const bannerApi = {
  getAll: async (limit = 10, offset = 0): Promise<PaginatedResponse<Banner>> => {
    const { data } = await apiClient.get<PaginatedResponse<Banner>>('/banners', {
      params: { limit, offset },
    })
    return data
  },
}
