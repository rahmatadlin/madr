import { apiClient, type ApiResponse, type PaginatedResponse } from './client'

export interface Banner {
  id: number
  title: string
  media_url: string
  type: 'image' | 'video'
  created_at: string
  updated_at: string
}

export interface CreateBannerRequest {
  title: string
  type: 'image' | 'video'
  media_url?: string
  file?: File
}

export type UpdateBannerRequest = Partial<CreateBannerRequest>

export const bannerApi = {
  getAll: async (limit = 10, offset = 0): Promise<PaginatedResponse<Banner>> => {
    const { data } = await apiClient.get<PaginatedResponse<Banner>>('/banners', { params: { limit, offset } })
    return data
  },
  getById: async (id: number): Promise<Banner> => {
    const { data } = await apiClient.get<ApiResponse<Banner>>(`/banners/${id}`)
    return data.data!
  },
  create: async (payload: CreateBannerRequest): Promise<Banner> => {
    const form = new FormData()
    form.append('title', payload.title)
    form.append('type', payload.type)
    if (payload.file) form.append('file', payload.file)
    else if (payload.media_url) form.append('media_url', payload.media_url)
    const { data } = await apiClient.post<ApiResponse<Banner>>('/admin/banners', form, { headers: { 'Content-Type': 'multipart/form-data' } })
    return data.data!
  },
  update: async (id: number, payload: UpdateBannerRequest): Promise<Banner> => {
    const form = new FormData()
    if (payload.title) form.append('title', payload.title)
    if (payload.type) form.append('type', payload.type)
    if (payload.file) form.append('file', payload.file)
    else if (payload.media_url) form.append('media_url', payload.media_url)
    const { data } = await apiClient.put<ApiResponse<Banner>>(`/admin/banners/${id}`, form, { headers: { 'Content-Type': 'multipart/form-data' } })
    return data.data!
  },
  delete: async (id: number): Promise<void> => {
    await apiClient.delete(`/admin/banners/${id}`)
  },
}
