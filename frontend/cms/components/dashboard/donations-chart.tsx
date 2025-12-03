"use client";

import { useMemo } from "react";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts";
import { Donation } from "@/lib/api/donations";

interface DonationsChartProps {
  donations: Donation[];
}

export function DonationsChart({ donations }: DonationsChartProps) {
  const chartData = useMemo(() => {
    const categoryMap = new Map<string, number>();

    donations.forEach((donation) => {
      const categoryName = donation.category?.name || "Unknown";
      const current = categoryMap.get(categoryName) || 0;
      categoryMap.set(categoryName, current + donation.amount);
    });

    return Array.from(categoryMap.entries()).map(([name, amount]) => ({
      name,
      amount: amount / 1000, // Convert to thousands for better display
    }));
  }, [donations]);

  if (chartData.length === 0) {
    return (
      <div className="flex h-64 items-center justify-center text-muted-foreground">
        No donation data available
      </div>
    );
  }

  return (
    <ResponsiveContainer width="100%" height={300}>
      <BarChart data={chartData}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="name" />
        <YAxis />
        <Tooltip
          formatter={(value: number) => `Rp ${(value * 1000).toLocaleString("id-ID")}`}
        />
        <Legend />
        <Bar dataKey="amount" fill="hsl(var(--primary))" name="Amount (Rp 1000)" />
      </BarChart>
    </ResponsiveContainer>
  );
}

