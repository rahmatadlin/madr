"use client";

import { useState } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { galleryApi } from "@/lib/api/gallery";
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
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Skeleton } from "@/components/ui/skeleton";
import Image from "next/image";

export default function GalleryPage() {
  const [page, setPage] = useState(0);
  const [search, setSearch] = useState("");
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [deleteConfirm, setDeleteConfirm] = useState<number | null>(null);
  const queryClient = useQueryClient();

  const limit = 10;
  const { data, isLoading } = useQuery({
    queryKey: ["gallery", limit, page * limit],
    queryFn: () => galleryApi.getAll(limit, page * limit),
  });

  const deleteMutation = useMutation({
    mutationFn: (id: number) => galleryApi.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["gallery"] });
      queryClient.invalidateQueries({ queryKey: ["stats"] });
      toast.success("Gallery item deleted successfully");
      setDeleteConfirm(null);
    },
    onError: () => {
      toast.error("Failed to delete gallery item");
    },
  });

  const filteredItems = data?.data.filter((item) =>
    item.title.toLowerCase().includes(search.toLowerCase())
  ) || [];

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold">Gallery</h1>
          <p className="text-muted-foreground">Manage gallery images</p>
        </div>
        <Button onClick={() => setIsModalOpen(true)}>
          <Plus className="mr-2 h-4 w-4" />
          Add Image
        </Button>
      </div>

      <div className="flex items-center gap-4">
        <div className="relative flex-1 max-w-sm">
          <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
          <Input
            placeholder="Search gallery..."
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
      ) : filteredItems.length === 0 ? (
        <div className="flex h-64 items-center justify-center text-muted-foreground">
          No gallery items found
        </div>
      ) : (
        <div className="rounded-md border">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Image</TableHead>
                <TableHead>Title</TableHead>
                <TableHead>Created</TableHead>
                <TableHead className="text-right">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {filteredItems.map((item) => (
                <TableRow key={item.id}>
                  <TableCell>
                    <div className="relative h-16 w-16 overflow-hidden rounded">
                      <Image
                        src={item.image_url}
                        alt={item.title}
                        fill
                        className="object-cover"
                      />
                    </div>
                  </TableCell>
                  <TableCell className="font-medium">{item.title}</TableCell>
                  <TableCell>
                    {new Date(item.created_at).toLocaleDateString("id-ID")}
                  </TableCell>
                  <TableCell className="text-right">
                    <Button
                      variant="ghost"
                      size="icon"
                      onClick={() => setDeleteConfirm(item.id)}
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

      {/* TODO: Add Gallery Modal for create/edit */}

      <Dialog open={!!deleteConfirm} onOpenChange={() => setDeleteConfirm(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Delete Gallery Item</DialogTitle>
            <DialogDescription>
              Are you sure you want to delete this gallery item? This action cannot be undone.
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

