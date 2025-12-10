import { AxiosError } from "axios";
import { apiClient, ApiResponse } from "./client";

export interface AboutContent {
  id: number;
  title: string;
  subtitle?: string;
  description?: string;
  additional_description?: string;
  image_url?: string;
  years_active?: number;
  active_members?: number;
  created_at?: string;
  updated_at?: string;
}

export const aboutApi = {
  get: async (): Promise<AboutContent | null> => {
    try {
      const response = await apiClient.get<ApiResponse<AboutContent>>("/about");
      return response.data.data || null;
    } catch (error) {
      const err = error as AxiosError<ApiResponse<AboutContent>>;
      if (err.response?.status === 404) {
        // Belum ada data about, biarkan fallback konten default
        return null;
      }
      throw error;
    }
  },
};
