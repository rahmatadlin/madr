import { apiClient, ApiResponse } from "./client";

export interface UploadResponse {
  filename: string;
  public_url: string;
  message: string;
}

export const uploadApi = {
  uploadFile: async (file: File, type: "image" | "video" = "image"): Promise<UploadResponse> => {
    const formData = new FormData();
    formData.append("file", file);
    formData.append("type", type);

    const response = await apiClient.post<ApiResponse<UploadResponse>>(
      "/admin/upload",
      formData,
      {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      }
    );
    return response.data.data!;
  },
};

