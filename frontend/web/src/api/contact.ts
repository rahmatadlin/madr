import { apiClient, type ApiResponse } from './client'

export interface ContactFormData {
  name: string
  email: string
  subject: string
  message: string
  honeypot?: string
}

export const contactApi = {
  submit: async (data: ContactFormData): Promise<void> => {
    await apiClient.post<ApiResponse<void>>('/contact', data)
  },
}
