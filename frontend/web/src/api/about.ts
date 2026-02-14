import type { AxiosError } from 'axios'
import { apiClient, type ApiResponse } from './client'

export interface AboutContent {
  id: number
  title: string
  subtitle?: string
  description?: string
  additional_description?: string
  image_url?: string
  years_active?: number
  active_members?: number
  created_at?: string
  updated_at?: string
}

export const aboutApi = {
  get: async (): Promise<AboutContent | null> => {
    try {
      const { data } = await apiClient.get<ApiResponse<AboutContent>>('/about')
      return data.data ?? null
    } catch (err) {
      if ((err as AxiosError<ApiResponse<AboutContent>>)?.response?.status === 404) return null
      throw err
    }
  },
}
