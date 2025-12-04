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
    try {
      const response = await apiClient.post<ApiResponse<LoginResponse>>(
        "/auth/login",
        credentials
      );

      // Backend returns { message, data: LoginResponse }
      if (!response.data.data) {
        throw new Error(response.data.error || "Login failed");
      }

      return response.data.data;
    } catch (error) {
      // Better error handling with detailed logging
      const axiosError = error as {
        response?: {
          status?: number;
          data?: { error?: string; message?: string; details?: string };
        };
        message?: string;
        code?: string;
      };

      // Log detailed error for debugging
      logger.error("Auth API error", {
        status: axiosError.response?.status,
        error: axiosError.response?.data?.error,
        message: axiosError.response?.data?.message,
        details: axiosError.response?.data?.details,
        code: axiosError.code,
        originalMessage: axiosError.message,
      });

      // Extract error message
      let errorMessage = "Login failed";

      if (axiosError.response?.data?.error) {
        errorMessage = axiosError.response.data.error;
      } else if (axiosError.response?.data?.message) {
        errorMessage = axiosError.response.data.message;
      } else if (axiosError.message) {
        errorMessage = axiosError.message;
      }

      // Add status code info for network errors
      if (axiosError.response?.status) {
        if (axiosError.response.status === 401) {
          errorMessage = "Invalid username or password";
        } else if (axiosError.response.status === 500) {
          errorMessage = "Server error. Please try again later.";
        } else if (
          axiosError.response.status === 0 ||
          !axiosError.response.status
        ) {
          errorMessage = "Network error. Please check your connection.";
        }
      } else if (
        axiosError.code === "ERR_NETWORK" ||
        axiosError.code === "ECONNREFUSED"
      ) {
        errorMessage =
          "Cannot connect to server. Please check if backend is running.";
      }

      throw new Error(errorMessage);
    }
  },
};
