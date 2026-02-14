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

export interface CreateEventRequest {
  title: string
  description: string
  date: string
  location: string
}

export type UpdateEventRequest = Partial<CreateEventRequest>

export const eventApi = {
  getAll: async (limit = 10, offset = 0): Promise<PaginatedResponse<Event>> => {
    const { data } = await apiClient.get<PaginatedResponse<Event>>('/events', { params: { limit, offset } })
    return data
  },
  getById: async (id: number): Promise<Event> => {
    const { data } = await apiClient.get<ApiResponse<Event>>(`/events/${id}`)
    return data.data!
  },
  create: async (payload: CreateEventRequest): Promise<Event> => {
    const { data } = await apiClient.post<ApiResponse<Event>>('/admin/events', payload)
    return data.data!
  },
  update: async (id: number, payload: UpdateEventRequest): Promise<Event> => {
    const { data } = await apiClient.put<ApiResponse<Event>>(`/admin/events/${id}`, payload)
    return data.data!
  },
  delete: async (id: number): Promise<void> => {
    await apiClient.delete(`/admin/events/${id}`)
  },
}
