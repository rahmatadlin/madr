import { apiClient, ApiResponse, PaginatedResponse } from "./client";

export interface GalleryItem {
  id: number;
  title: string;
  image_url: string;
  created_at: string;
}

export interface CreateGalleryRequest {
  title: string;
  image_url?: string;
  file?: File;
}

export const galleryApi = {
  getAll: async (limit = 10, offset = 0): Promise<PaginatedResponse<GalleryItem>> => {
    const response = await apiClient.get<PaginatedResponse<GalleryItem>>(
      "/admin/gallery",
      {
        params: { limit, offset },
      }
    );
    return response.data;
  },
  create: async (data: CreateGalleryRequest): Promise<GalleryItem> => {
    const formData = new FormData();
    formData.append("title", data.title);
    if (data.file) {
      formData.append("file", data.file);
    } else if (data.image_url) {
      formData.append("image_url", data.image_url);
    }

    const response = await apiClient.post<ApiResponse<GalleryItem>>(
      "/admin/gallery",
      formData,
      {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      }
    );
    return response.data.data!;
  },
  delete: async (id: number): Promise<void> => {
    await apiClient.delete(`/admin/gallery/${id}`);
  },
};

