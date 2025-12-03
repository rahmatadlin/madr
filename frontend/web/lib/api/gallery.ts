import { apiClient, ApiResponse, PaginatedResponse } from "./client";

export interface GalleryItem {
  id: number;
  title: string;
  image_url: string;
  created_at: string;
}

export const galleryApi = {
  getAll: async (limit = 10, offset = 0): Promise<PaginatedResponse<GalleryItem>> => {
    const response = await apiClient.get<PaginatedResponse<GalleryItem>>("/gallery", {
      params: { limit, offset },
    });
    return response.data;
  },
};

