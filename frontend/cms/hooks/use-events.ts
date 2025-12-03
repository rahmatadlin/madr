import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { eventApi, type Event, type CreateEventRequest, type UpdateEventRequest } from "@/lib/api/events";
import { toast } from "sonner";
import { logger } from "@/lib/logger";

export function useEvents(limit = 10, offset = 0) {
  return useQuery({
    queryKey: ["events", limit, offset],
    queryFn: () => eventApi.getAll(limit, offset),
  });
}

export function useEvent(id: number) {
  return useQuery({
    queryKey: ["event", id],
    queryFn: () => eventApi.getById(id),
    enabled: !!id,
  });
}

export function useCreateEvent() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateEventRequest) => eventApi.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["events"] });
      queryClient.invalidateQueries({ queryKey: ["stats"] });
      toast.success("Event created successfully");
      logger.info("Event created");
    },
    onError: (error) => {
      toast.error("Failed to create event");
      logger.error("Create event error", error);
    },
  });
}

export function useUpdateEvent() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: number; data: UpdateEventRequest }) =>
      eventApi.update(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: ["events"] });
      queryClient.invalidateQueries({ queryKey: ["event", variables.id] });
      toast.success("Event updated successfully");
      logger.info("Event updated", { id: variables.id });
    },
    onError: (error) => {
      toast.error("Failed to update event");
      logger.error("Update event error", error);
    },
  });
}

export function useDeleteEvent() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number) => eventApi.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["events"] });
      queryClient.invalidateQueries({ queryKey: ["stats"] });
      toast.success("Event deleted successfully");
      logger.info("Event deleted");
    },
    onError: (error) => {
      toast.error("Failed to delete event");
      logger.error("Delete event error", error);
    },
  });
}

