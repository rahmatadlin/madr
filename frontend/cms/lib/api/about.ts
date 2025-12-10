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

export interface UpdateAboutRequest {
  title?: string;
  subtitle?: string;
  description?: string;
  additional_description?: string;
  image_url?: string;
  years_active?: number;
  active_members?: number;
}

export const aboutApi = {
  get: async (): Promise<AboutContent> => {
    const response = await apiClient.get<ApiResponse<AboutContent>>("/admin/about");
    return response.data.data!;
  },
  update: async (data: UpdateAboutRequest): Promise<AboutContent> => {
    const response = await apiClient.put<ApiResponse<AboutContent>>("/admin/about", data);
    return response.data.data!;
  },
};

