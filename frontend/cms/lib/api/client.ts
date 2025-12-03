import axios, { AxiosError, AxiosInstance } from "axios";
import { logger } from "../logger";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api/v1";

export const apiClient: AxiosInstance = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// Request interceptor
apiClient.interceptors.request.use(
  (config) => {
    logger.api(config.method?.toUpperCase() || "REQUEST", config.url || "", config.data);
    return config;
  },
  (error) => {
    logger.error("API Request Error", error);
    return Promise.reject(error);
  }
);

// Response interceptor
apiClient.interceptors.response.use(
  (response) => {
    logger.debug(`API Response ${response.status}`, response.data);
    return response;
  },
  (error: AxiosError) => {
    if (error.response) {
      logger.error(
        `API Error ${error.response.status}`,
        error.response.data || error.message
      );
    } else {
      logger.error("API Network Error", error.message);
    }
    return Promise.reject(error);
  }
);

// Helper to set auth token
export const setAuthToken = (token: string | null) => {
  if (token) {
    apiClient.defaults.headers.common["Authorization"] = `Bearer ${token}`;
  } else {
    delete apiClient.defaults.headers.common["Authorization"];
  }
};

// Response types
export interface ApiResponse<T> {
  data?: T;
  message?: string;
  error?: string;
}

export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  limit: number;
  offset: number;
  total_pages: number;
}

