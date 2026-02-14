import { apiClient, type ApiResponse } from './client'

export interface LoginRequest {
  username: string
  password: string
}

export interface LoginResponse {
  access_token: string
  refresh_token: string
  user: { id: number; username: string; email: string; role: string }
}

export const authApi = {
  login: async (credentials: LoginRequest): Promise<LoginResponse> => {
    const { data } = await apiClient.post<ApiResponse<LoginResponse>>('/auth/login', credentials)
    if (!data.data) throw new Error(data.error || 'Login failed')
    return data.data
  },
}
