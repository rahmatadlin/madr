import { apiClient, ApiResponse } from "./client";

export interface DashboardStats {
  total_events: number;
  total_banners: number;
  total_gallery: number;
  total_donations: number;
}

export const statsApi = {
  getDashboardStats: async (): Promise<DashboardStats> => {
    // Since backend doesn't have a dedicated stats endpoint,
    // we'll fetch counts from individual endpoints
    // In production, you might want to create a dedicated endpoint
    
    const [events, banners, gallery, donations] = await Promise.all([
      apiClient.get("/admin/events?limit=1&offset=0"),
      apiClient.get("/admin/banners?limit=1&offset=0"),
      apiClient.get("/admin/gallery?limit=1&offset=0"),
      apiClient.get("/admin/donations?limit=1&offset=0"),
    ]);

    return {
      total_events: events.data.total || 0,
      total_banners: banners.data.total || 0,
      total_gallery: gallery.data.total || 0,
      total_donations: donations.data.total || 0,
    };
  },
};

