import { apiClient, type ApiResponse } from './client'

export interface UploadResponse {
  url?: string
  filename?: string
  public_url?: string
}

export const uploadApi = {
  uploadFile: async (file: File): Promise<UploadResponse> => {
    const form = new FormData()
    form.append('file', file)
    const { data } = await apiClient.post<ApiResponse<UploadResponse>>('/admin/upload', form, { headers: { 'Content-Type': 'multipart/form-data' } })
    const d = data.data
    const url = (d && (d.url ?? d.public_url)) || ''
    return { ...d, url }
  },
}
