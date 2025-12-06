import { apiClient } from "./client";

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
    // Using Promise.allSettled to handle errors gracefully
    // Note: Some endpoints use public routes as admin GET endpoints don't exist

    const results = await Promise.allSettled([
      // Events: Use public endpoint (admin GET doesn't exist)
      apiClient.get("/events?limit=1&offset=0").catch(() => null),
      // Banners: Use public endpoint (admin GET doesn't exist)
      apiClient.get("/banners?limit=1&offset=0").catch(() => null),
      // Gallery: Use public endpoint (admin GET doesn't exist)
      apiClient.get("/gallery?limit=1&offset=0").catch(() => null),
      // Donations: Use admin endpoint (requires auth)
      apiClient.get("/admin/donations?limit=1&offset=0").catch(() => null),
    ]);

    const getTotal = (
      result: PromiseSettledResult<{ data?: { total?: number } } | null>
    ): number => {
      if (result.status === "fulfilled" && result.value?.data?.total) {
        return result.value.data.total;
      }
      return 0;
    };

    return {
      total_events: getTotal(results[0]),
      total_banners: getTotal(results[1]),
      total_gallery: getTotal(results[2]),
      total_donations: getTotal(results[3]),
    };
  },
};
