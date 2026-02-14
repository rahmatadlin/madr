import { apiClient, type ApiResponse, type PaginatedResponse } from './client'

export interface Event {
  id: number
  title: string
  description: string
  date: string
  location: string
  created_at: string
  updated_at: string
}

export const eventApi = {
  getAll: async (limit = 10, offset = 0): Promise<PaginatedResponse<Event>> => {
    const { data } = await apiClient.get<PaginatedResponse<Event>>('/events', {
      params: { limit, offset },
    })
    return data
  },
  getById: async (id: number): Promise<Event> => {
    const { data } = await apiClient.get<ApiResponse<Event>>(`/events/${id}`)
    return data.data!
  },
}
