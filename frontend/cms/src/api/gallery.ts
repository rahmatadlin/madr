import { apiClient, type ApiResponse, type PaginatedResponse } from './client'

export interface GalleryItem {
  id: number
  title: string
  image_url: string
  created_at: string
}

export interface CreateGalleryRequest {
  title: string
  image_url?: string
  file?: File
}

export const galleryApi = {
  getAll: async (limit = 10, offset = 0): Promise<PaginatedResponse<GalleryItem>> => {
    const { data } = await apiClient.get<PaginatedResponse<GalleryItem>>('/gallery', { params: { limit, offset } })
    return data
  },
  create: async (payload: CreateGalleryRequest): Promise<GalleryItem> => {
    const form = new FormData()
    form.append('title', payload.title)
    if (payload.file) form.append('file', payload.file)
    else if (payload.image_url) form.append('image_url', payload.image_url)
    const { data } = await apiClient.post<ApiResponse<GalleryItem>>('/admin/gallery', form, { headers: { 'Content-Type': 'multipart/form-data' } })
    return data.data!
  },
  delete: async (id: number): Promise<void> => {
    await apiClient.delete(`/admin/gallery/${id}`)
  },
}
