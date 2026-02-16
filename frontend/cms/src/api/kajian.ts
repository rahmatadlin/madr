import { apiClient, type ApiResponse } from './client'

export interface Kajian {
  id: number
  video_id: string
  title: string
  description: string
  published_at: string
  thumbnail_url: string
  youtube_url: string
  channel_title: string
  created_at: string
  updated_at: string
}

export interface KajianListResponse {
  data: Kajian[]
  total: number
  limit: number
  offset: number
  total_pages: number
}

export interface SyncResponse {
  message: string
  synced: number
}

export const kajianApi = {
  getAll: async (limit = 10, offset = 0): Promise<KajianListResponse> => {
    const { data } = await apiClient.get<KajianListResponse>('/kajian', {
      params: { limit, offset },
    })
    return data
  },

  getById: async (id: number): Promise<ApiResponse<Kajian>> => {
    const { data } = await apiClient.get<ApiResponse<Kajian>>(`/kajian/${id}`)
    return data
  },

  syncFromYouTube: async (days = 30): Promise<SyncResponse> => {
    const { data } = await apiClient.post<SyncResponse>('/admin/kajian/sync', null, {
      params: { days },
    })
    return data
  },

  delete: async (id: number): Promise<void> => {
    await apiClient.delete(`/admin/kajian/${id}`)
  },
}
