"use client";

import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { useEvent, useCreateEvent, useUpdateEvent } from "@/hooks/use-events";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Label } from "@/components/ui/label";
import { Skeleton } from "@/components/ui/skeleton";

const eventSchema = z.object({
  title: z.string().min(3, "Title must be at least 3 characters"),
  description: z.string().optional(),
  date: z.string().min(1, "Date is required"),
  location: z.string().min(1, "Location is required"),
});

type EventFormData = z.infer<typeof eventSchema>;

interface EventModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  eventId: number | null;
  onSuccess: () => void;
}

export function EventModal({
  open,
  onOpenChange,
  eventId,
  onSuccess,
}: EventModalProps) {
  const { data: event, isLoading } = useEvent(eventId || 0);
  const createMutation = useCreateEvent();
  const updateMutation = useUpdateEvent();

  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<EventFormData>({
    resolver: zodResolver(eventSchema),
  });

  useEffect(() => {
    if (event && open) {
      const eventDate = new Date(event.date);
      const localDateTime = new Date(
        eventDate.getTime() - eventDate.getTimezoneOffset() * 60000
      )
        .toISOString()
        .slice(0, 16);
      reset({
        title: event.title,
        description: event.description,
        date: localDateTime,
        location: event.location,
      });
    } else if (!eventId && open) {
      reset({
        title: "",
        description: "",
        date: "",
        location: "",
      });
    }
  }, [event, eventId, open, reset]);

  const onSubmit = async (data: EventFormData) => {
    // Convert datetime-local format (YYYY-MM-DDTHH:mm) to RFC3339 (ISO 8601)
    // datetime-local doesn't include seconds and timezone, so we convert to ISO format
    const dateValue = data.date ? new Date(data.date).toISOString() : "";

    const submitData = {
      title: data.title,
      description: data.description || "",
      date: dateValue,
      location: data.location,
    };

    if (eventId) {
      await updateMutation.mutateAsync({ id: eventId, data: submitData });
    } else {
      await createMutation.mutateAsync(submitData);
    }
    onSuccess();
  };

  const isLoadingForm =
    isLoading || createMutation.isPending || updateMutation.isPending;

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-2xl">
        <DialogHeader>
          <DialogTitle>{eventId ? "Edit Event" : "Create Event"}</DialogTitle>
          <DialogDescription>
            {eventId
              ? "Update event information"
              : "Add a new event to the system"}
          </DialogDescription>
        </DialogHeader>

        {isLoading && eventId ? (
          <div className="space-y-4">
            <Skeleton className="h-10 w-full" />
            <Skeleton className="h-24 w-full" />
            <Skeleton className="h-10 w-full" />
            <Skeleton className="h-10 w-full" />
          </div>
        ) : (
          <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="title">Title *</Label>
              <Input
                id="title"
                {...register("title")}
                placeholder="Event title"
                disabled={isLoadingForm}
              />
              {errors.title && (
                <p className="text-sm text-destructive">
                  {errors.title.message}
                </p>
              )}
            </div>

            <div className="space-y-2">
              <Label htmlFor="description">Description</Label>
              <Textarea
                id="description"
                {...register("description")}
                placeholder="Event description"
                rows={4}
                disabled={isLoadingForm}
              />
            </div>

            <div className="grid grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label htmlFor="date">Date *</Label>
                <Input
                  id="date"
                  type="datetime-local"
                  {...register("date")}
                  disabled={isLoadingForm}
                />
                {errors.date && (
                  <p className="text-sm text-destructive">
                    {errors.date.message}
                  </p>
                )}
              </div>

              <div className="space-y-2">
                <Label htmlFor="location">Location *</Label>
                <Input
                  id="location"
                  {...register("location")}
                  placeholder="Event location"
                  disabled={isLoadingForm}
                />
                {errors.location && (
                  <p className="text-sm text-destructive">
                    {errors.location.message}
                  </p>
                )}
              </div>
            </div>

            <div className="flex justify-end gap-2 pt-4">
              <Button
                type="button"
                variant="outline"
                onClick={() => onOpenChange(false)}
                disabled={isLoadingForm}
              >
                Cancel
              </Button>
              <Button type="submit" disabled={isLoadingForm}>
                {isLoadingForm
                  ? "Saving..."
                  : eventId
                  ? "Update Event"
                  : "Create Event"}
              </Button>
            </div>
          </form>
        )}
      </DialogContent>
    </Dialog>
  );
}
