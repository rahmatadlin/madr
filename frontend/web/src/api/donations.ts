import { apiClient, type ApiResponse } from './client'

export interface DonationSummary {
  total_amount: number
  total_transactions: number
  per_category: Array<{ category_id: number; category: string; amount: number }>
}

export const donationApi = {
  getSummary: async (): Promise<DonationSummary> => {
    const { data } = await apiClient.get<ApiResponse<DonationSummary>>('/donations/summary')
    return data.data!
  },
}
