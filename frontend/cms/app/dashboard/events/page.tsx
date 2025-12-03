"use client";

import { useState } from "react";
import { useEvents, useDeleteEvent } from "@/hooks/use-events";
import { Button } from "@/components/ui/button";
import { Plus, Search } from "lucide-react";
import { Input } from "@/components/ui/input";
import { EventsTable } from "@/components/events/events-table";
import { EventModal } from "@/components/events/event-modal";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";

export default function EventsPage() {
  const [page, setPage] = useState(0);
  const [search, setSearch] = useState("");
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingEvent, setEditingEvent] = useState<number | null>(null);
  const [deleteConfirm, setDeleteConfirm] = useState<number | null>(null);

  const limit = 10;
  const { data, isLoading } = useEvents(limit, page * limit);
  const deleteMutation = useDeleteEvent();

  const filteredEvents = data?.data.filter((event) =>
    event.title.toLowerCase().includes(search.toLowerCase())
  ) || [];

  const handleEdit = (id: number) => {
    setEditingEvent(id);
    setIsModalOpen(true);
  };

  const handleDelete = (id: number) => {
    setDeleteConfirm(id);
  };

  const confirmDelete = () => {
    if (deleteConfirm) {
      deleteMutation.mutate(deleteConfirm);
      setDeleteConfirm(null);
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold">Events</h1>
          <p className="text-muted-foreground">Manage events and activities</p>
        </div>
        <Button onClick={() => setIsModalOpen(true)}>
          <Plus className="mr-2 h-4 w-4" />
          Add Event
        </Button>
      </div>

      <div className="flex items-center gap-4">
        <div className="relative flex-1 max-w-sm">
          <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
          <Input
            placeholder="Search events..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="pl-9"
          />
        </div>
      </div>

      <EventsTable
        events={filteredEvents}
        isLoading={isLoading}
        onEdit={handleEdit}
        onDelete={handleDelete}
        pagination={{
          page,
          totalPages: data?.total_pages || 0,
          onPageChange: setPage,
        }}
      />

      <EventModal
        open={isModalOpen}
        onOpenChange={setIsModalOpen}
        eventId={editingEvent}
        onSuccess={() => {
          setIsModalOpen(false);
          setEditingEvent(null);
        }}
      />

      {/* Delete Confirmation Dialog */}
      <Dialog open={!!deleteConfirm} onOpenChange={() => setDeleteConfirm(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Delete Event</DialogTitle>
            <DialogDescription>
              Are you sure you want to delete this event? This action cannot be undone.
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
              onClick={confirmDelete}
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

