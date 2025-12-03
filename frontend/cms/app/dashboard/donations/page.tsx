"use client";

import { useState } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { donationApi } from "@/lib/api/donations";
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
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

export default function DonationsPage() {
  const [page, setPage] = useState(0);
  const [search, setSearch] = useState("");
  const [statusFilter, setStatusFilter] = useState<string>("");
  const [deleteConfirm, setDeleteConfirm] = useState<number | null>(null);
  const queryClient = useQueryClient();

  const limit = 10;
  const { data, isLoading } = useQuery({
    queryKey: ["donations", limit, page * limit, statusFilter],
    queryFn: () => donationApi.getAll(limit, page * limit, statusFilter || undefined),
  });

  const deleteMutation = useMutation({
    mutationFn: (id: number) => donationApi.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["donations"] });
      queryClient.invalidateQueries({ queryKey: ["stats"] });
      toast.success("Donation deleted successfully");
      setDeleteConfirm(null);
    },
    onError: () => {
      toast.error("Failed to delete donation");
    },
  });

  const filteredDonations = data?.data.filter((donation) =>
    donation.donor_name?.toLowerCase().includes(search.toLowerCase()) ||
    donation.message?.toLowerCase().includes(search.toLowerCase())
  ) || [];

  const formatCurrency = (amount: number) => {
    return new Intl.NumberFormat("id-ID", {
      style: "currency",
      currency: "IDR",
      minimumFractionDigits: 0,
    }).format(amount);
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold">Donations</h1>
          <p className="text-muted-foreground">Manage donations and transactions</p>
        </div>
        <Button>
          <Plus className="mr-2 h-4 w-4" />
          Add Donation
        </Button>
      </div>

      <div className="flex items-center gap-4">
        <div className="relative flex-1 max-w-sm">
          <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
          <Input
            placeholder="Search donations..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="pl-9"
          />
        </div>
        <Select value={statusFilter} onValueChange={setStatusFilter}>
          <SelectTrigger className="w-[180px]">
            <SelectValue placeholder="Filter by status" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="">All Status</SelectItem>
            <SelectItem value="pending">Pending</SelectItem>
            <SelectItem value="success">Success</SelectItem>
            <SelectItem value="failed">Failed</SelectItem>
          </SelectContent>
        </Select>
      </div>

      {isLoading ? (
        <div className="space-y-2">
          {[...Array(5)].map((_, i) => (
            <Skeleton key={i} className="h-16 w-full" />
          ))}
        </div>
      ) : filteredDonations.length === 0 ? (
        <div className="flex h-64 items-center justify-center text-muted-foreground">
          No donations found
        </div>
      ) : (
        <div className="rounded-md border">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Donor</TableHead>
                <TableHead>Category</TableHead>
                <TableHead>Amount</TableHead>
                <TableHead>Status</TableHead>
                <TableHead>Created</TableHead>
                <TableHead className="text-right">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {filteredDonations.map((donation) => (
                <TableRow key={donation.id}>
                  <TableCell className="font-medium">
                    {donation.donor_name || "Anonymous"}
                  </TableCell>
                  <TableCell>{donation.category?.name || "N/A"}</TableCell>
                  <TableCell>{formatCurrency(donation.amount)}</TableCell>
                  <TableCell>
                    <Badge
                      variant={
                        donation.payment_status === "success"
                          ? "default"
                          : donation.payment_status === "pending"
                          ? "secondary"
                          : "destructive"
                      }
                    >
                      {donation.payment_status}
                    </Badge>
                  </TableCell>
                  <TableCell>
                    {new Date(donation.created_at).toLocaleDateString("id-ID")}
                  </TableCell>
                  <TableCell className="text-right">
                    <Button
                      variant="ghost"
                      size="icon"
                      onClick={() => setDeleteConfirm(donation.id)}
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

      {/* TODO: Add Donation Modal for create/edit */}

      <Dialog open={!!deleteConfirm} onOpenChange={() => setDeleteConfirm(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Delete Donation</DialogTitle>
            <DialogDescription>
              Are you sure you want to delete this donation? This action cannot be undone.
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

