import { apiClient, type ApiResponse, type PaginatedResponse } from './client'

export interface Donation {
  id: number
  category_id: number
  donor_name: string | null
  amount: number
  message: string
  payment_status: 'pending' | 'success' | 'failed'
  created_at: string
  category?: { id: number; name: string }
}

export interface DonationCategory {
  id: number
  name: string
  description: string
  created_at: string
}

export const donationApi = {
  getAll: async (limit = 10, offset = 0, status?: string): Promise<PaginatedResponse<Donation>> => {
    const params: Record<string, string | number> = { limit, offset }
    if (status) params.status = status
    const { data } = await apiClient.get<PaginatedResponse<Donation>>('/admin/donations', { params })
    return data
  },
  getCategories: async (): Promise<DonationCategory[]> => {
    const { data } = await apiClient.get<ApiResponse<DonationCategory[]>>('/admin/donation-categories')
    return data.data || []
  },
}
