"use client";

import { useDashboardStats } from "@/hooks/use-stats";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Calendar, FileImage, Image, DollarSign } from "lucide-react";
import { Skeleton } from "@/components/ui/skeleton";
import { DonationsChart } from "@/components/dashboard/donations-chart";
import { useQuery } from "@tanstack/react-query";
import { donationApi } from "@/lib/api/donations";

export default function DashboardPage() {
  const { data: stats, isLoading: statsLoading } = useDashboardStats();
  const { data: donationsData } = useQuery({
    queryKey: ["donations", "chart"],
    queryFn: () => donationApi.getAll(100, 0, "success"),
  });

  const statCards = [
    {
      title: "Total Events",
      value: stats?.total_events || 0,
      icon: Calendar,
      description: "Active events",
    },
    {
      title: "Total Banners",
      value: stats?.total_banners || 0,
      icon: FileImage,
      description: "Banner items",
    },
    {
      title: "Gallery Images",
      value: stats?.total_gallery || 0,
      icon: Image,
      description: "Total images",
    },
    {
      title: "Total Donations",
      value: stats?.total_donations || 0,
      icon: DollarSign,
      description: "All donations",
    },
  ];

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold">Dashboard</h1>
        <p className="text-muted-foreground">Welcome to Masjid CMS</p>
      </div>

      {/* Stats Grid */}
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        {statCards.map((stat) => {
          const Icon = stat.icon;
          return (
            <Card key={stat.title}>
              <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                <CardTitle className="text-sm font-medium">
                  {stat.title}
                </CardTitle>
                <Icon className="h-4 w-4 text-muted-foreground" />
              </CardHeader>
              <CardContent>
                {statsLoading ? (
                  <Skeleton className="h-8 w-20" />
                ) : (
                  <>
                    <div className="text-2xl font-bold">{stat.value}</div>
                    <p className="text-xs text-muted-foreground">
                      {stat.description}
                    </p>
                  </>
                )}
              </CardContent>
            </Card>
          );
        })}
      </div>

      {/* Donations Chart */}
      <Card>
        <CardHeader>
          <CardTitle>Donations Overview</CardTitle>
          <CardDescription>Donation statistics by category</CardDescription>
        </CardHeader>
        <CardContent>
          <DonationsChart donations={donationsData?.data || []} />
        </CardContent>
      </Card>
    </div>
  );
}

