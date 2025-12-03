"use client";

import { useState } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { bannerApi } from "@/lib/api/banners";
import { Button } from "@/components/ui/button";
import { Plus, Search } from "lucide-react";
import { Input } from "@/components/ui/input";
import { toast } from "sonner";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Skeleton } from "@/components/ui/skeleton";

export default function BannerPage() {
  const [page, setPage] = useState(0);
  const [search, setSearch] = useState("");
  const [deleteConfirm, setDeleteConfirm] = useState<number | null>(null);
  const queryClient = useQueryClient();

  const limit = 10;
  const { data, isLoading } = useQuery({
    queryKey: ["banners", limit, page * limit],
    queryFn: () => bannerApi.getAll(limit, page * limit),
  });

  const deleteMutation = useMutation({
    mutationFn: (id: number) => bannerApi.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["banners"] });
      queryClient.invalidateQueries({ queryKey: ["stats"] });
      toast.success("Banner deleted successfully");
      setDeleteConfirm(null);
    },
    onError: () => {
      toast.error("Failed to delete banner");
    },
  });

  const filteredBanners = data?.data.filter((banner) =>
    banner.title.toLowerCase().includes(search.toLowerCase())
  ) || [];

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold">Banners</h1>
          <p className="text-muted-foreground">Manage banners and media</p>
        </div>
        <Button>
          <Plus className="mr-2 h-4 w-4" />
          Add Banner
        </Button>
      </div>

      <div className="flex items-center gap-4">
        <div className="relative flex-1 max-w-sm">
          <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
          <Input
            placeholder="Search banners..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="pl-9"
          />
        </div>
      </div>

      {isLoading ? (
        <div className="space-y-2">
          {[...Array(5)].map((_, i) => (
            <Skeleton key={i} className="h-16 w-full" />
          ))}
        </div>
      ) : filteredBanners.length === 0 ? (
        <div className="flex h-64 items-center justify-center text-muted-foreground">
          No banners found
        </div>
      ) : (
        <div className="rounded-md border">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Title</TableHead>
                <TableHead>Type</TableHead>
                <TableHead>Media URL</TableHead>
                <TableHead>Created</TableHead>
                <TableHead className="text-right">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {filteredBanners.map((banner) => (
                <TableRow key={banner.id}>
                  <TableCell className="font-medium">{banner.title}</TableCell>
                  <TableCell>
                    <Badge variant={banner.type === "image" ? "default" : "secondary"}>
                      {banner.type}
                    </Badge>
                  </TableCell>
                  <TableCell className="max-w-xs truncate">{banner.media_url}</TableCell>
                  <TableCell>
                    {new Date(banner.created_at).toLocaleDateString("id-ID")}
                  </TableCell>
                  <TableCell className="text-right">
                    <Button
                      variant="ghost"
                      size="icon"
                      onClick={() => setDeleteConfirm(banner.id)}
                    >
                      <span className="text-destructive">Delete</span>
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </div>
      )}

      {/* TODO: Add Banner Modal for create/edit */}

      <Dialog open={!!deleteConfirm} onOpenChange={() => setDeleteConfirm(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Delete Banner</DialogTitle>
            <DialogDescription>
              Are you sure you want to delete this banner? This action cannot be undone.
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => setDeleteConfirm(null)}
              disabled={deleteMutation.isPending}
            >
              Cancel
            </Button>
            <Button
              variant="destructive"
              onClick={() => deleteMutation.mutate(deleteConfirm!)}
              disabled={deleteMutation.isPending}
            >
              {deleteMutation.isPending ? "Deleting..." : "Delete"}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}

