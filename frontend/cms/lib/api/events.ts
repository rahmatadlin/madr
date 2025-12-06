import { apiClient, ApiResponse, PaginatedResponse } from "./client";

export interface Event {
  id: number;
  title: string;
  description: string;
  date: string;
  location: string;
  created_at: string;
  updated_at: string;
}

export interface CreateEventRequest {
  title: string;
  description: string;
  date: string;
  location: string;
}

export type UpdateEventRequest = Partial<CreateEventRequest>;

export const eventApi = {
  getAll: async (limit = 10, offset = 0): Promise<PaginatedResponse<Event>> => {
    // Use public endpoint for GET (admin GET doesn't exist)
    const response = await apiClient.get<PaginatedResponse<Event>>("/events", {
      params: { limit, offset },
    });
    return response.data;
  },
  getById: async (id: number): Promise<Event> => {
    // Use public endpoint for GET (admin GET doesn't exist)
    const response = await apiClient.get<ApiResponse<Event>>(`/events/${id}`);
    return response.data.data!;
  },
  create: async (data: CreateEventRequest): Promise<Event> => {
    const response = await apiClient.post<ApiResponse<Event>>(
      "/admin/events",
      data
    );
    return response.data.data!;
  },
  update: async (id: number, data: UpdateEventRequest): Promise<Event> => {
    const response = await apiClient.put<ApiResponse<Event>>(
      `/admin/events/${id}`,
      data
    );
    return response.data.data!;
  },
  delete: async (id: number): Promise<void> => {
    await apiClient.delete(`/admin/events/${id}`);
  },
};
