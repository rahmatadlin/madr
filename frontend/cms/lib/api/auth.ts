import { apiClient, ApiResponse } from "./client";

export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  access_token: string;
  refresh_token: string;
  user: {
    id: number;
    username: string;
    email: string;
    role: string;
  };
}

export const authApi = {
  login: async (credentials: LoginRequest): Promise<LoginResponse> => {
    const response = await apiClient.post<ApiResponse<LoginResponse>>(
      "/auth/login",
      credentials
    );
    return response.data.data!;
  },
};
