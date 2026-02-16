import { apiClient } from './client'

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

export const kajianApi = {
  getAll: async (limit = 10, offset = 0): Promise<KajianListResponse> => {
    const { data } = await apiClient.get<KajianListResponse>('/kajian', {
      params: { limit, offset },
    })
    return data
  },

  getById: async (id: number): Promise<{ data: Kajian }> => {
    const { data } = await apiClient.get<{ data: Kajian }>(`/kajian/${id}`)
    return data
  },
}
