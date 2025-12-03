import { apiClient, ApiResponse, PaginatedResponse } from "./client";

export interface Donation {
  id: number;
  category_id: number;
  donor_name: string | null;
  amount: number;
  message: string;
  payment_status: "pending" | "success" | "failed";
  created_at: string;
  category?: {
    id: number;
    name: string;
  };
}

export interface DonationCategory {
  id: number;
  name: string;
  description: string;
  created_at: string;
}

export interface CreateDonationRequest {
  category_id: number;
  donor_name?: string;
  amount: number;
  message?: string;
  payment_status?: "pending" | "success" | "failed";
}

export interface UpdateDonationRequest extends Partial<CreateDonationRequest> {}

export const donationApi = {
  getAll: async (
    limit = 10,
    offset = 0,
    status?: string
  ): Promise<PaginatedResponse<Donation>> => {
    const params: Record<string, unknown> = { limit, offset };
    if (status) params.status = status;

    const response = await apiClient.get<PaginatedResponse<Donation>>(
      "/admin/donations",
      { params }
    );
    return response.data;
  },
  getById: async (id: number): Promise<Donation> => {
    const response = await apiClient.get<ApiResponse<Donation>>(
      `/admin/donations/${id}`
    );
    return response.data.data!;
  },
  create: async (data: CreateDonationRequest): Promise<Donation> => {
    const response = await apiClient.post<ApiResponse<Donation>>(
      "/admin/donations",
      data
    );
    return response.data.data!;
  },
  update: async (id: number, data: UpdateDonationRequest): Promise<Donation> => {
    const response = await apiClient.put<ApiResponse<Donation>>(
      `/admin/donations/${id}`,
      data
    );
    return response.data.data!;
  },
  delete: async (id: number): Promise<void> => {
    await apiClient.delete(`/admin/donations/${id}`);
  },
  getCategories: async (): Promise<DonationCategory[]> => {
    const response = await apiClient.get<ApiResponse<DonationCategory[]>>(
      "/admin/donation-categories"
    );
    return response.data.data || [];
  },
};

