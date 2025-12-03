import { apiClient, ApiResponse, PaginatedResponse } from "./client";

export interface Banner {
  id: number;
  title: string;
  media_url: string;
  type: "image" | "video";
  created_at: string;
  updated_at: string;
}

export interface CreateBannerRequest {
  title: string;
  type: "image" | "video";
  media_url?: string;
  file?: File;
}

export interface UpdateBannerRequest extends Partial<CreateBannerRequest> {}

export const bannerApi = {
  getAll: async (limit = 10, offset = 0): Promise<PaginatedResponse<Banner>> => {
    const response = await apiClient.get<PaginatedResponse<Banner>>("/admin/banners", {
      params: { limit, offset },
    });
    return response.data;
  },
  getById: async (id: number): Promise<Banner> => {
    const response = await apiClient.get<ApiResponse<Banner>>(`/admin/banners/${id}`);
    return response.data.data!;
  },
  create: async (data: CreateBannerRequest): Promise<Banner> => {
    const formData = new FormData();
    formData.append("title", data.title);
    formData.append("type", data.type);
    if (data.file) {
      formData.append("file", data.file);
    } else if (data.media_url) {
      formData.append("media_url", data.media_url);
    }

    const response = await apiClient.post<ApiResponse<Banner>>("/admin/banners", formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
    return response.data.data!;
  },
  update: async (id: number, data: UpdateBannerRequest): Promise<Banner> => {
    const formData = new FormData();
    if (data.title) formData.append("title", data.title);
    if (data.type) formData.append("type", data.type);
    if (data.file) {
      formData.append("file", data.file);
    } else if (data.media_url) {
      formData.append("media_url", data.media_url);
    }

    const response = await apiClient.put<ApiResponse<Banner>>(
      `/admin/banners/${id}`,
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
    await apiClient.delete(`/admin/banners/${id}`);
  },
};

