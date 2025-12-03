import { useQuery } from "@tanstack/react-query";
import { statsApi } from "@/lib/api/stats";

export function useDashboardStats() {
  return useQuery({
    queryKey: ["stats"],
    queryFn: () => statsApi.getDashboardStats(),
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}

