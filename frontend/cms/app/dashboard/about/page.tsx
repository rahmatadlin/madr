"use client";

import { useEffect, useRef, useState } from "react";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useForm, type Resolver, type SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { aboutApi } from "@/lib/api/about";
import { apiClient } from "@/lib/api/client";
import Image from "next/image";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Label } from "@/components/ui/label";
import { Plus } from "lucide-react";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { toast } from "sonner";

const aboutSchema = z.object({
  title: z.string().min(3, "Judul minimal 3 karakter"),
  subtitle: z.string().optional(),
  description: z.string().optional(),
  additional_description: z.string().optional(),
  image_url: z
    .string()
    .url("Harus berupa URL yang valid")
    .optional()
    .or(z.literal("")),
  years_active: z.coerce.number().min(0, "Tidak boleh negatif").optional(),
  active_members: z.coerce.number().min(0, "Tidak boleh negatif").optional(),
});

type AboutFormData = z.infer<typeof aboutSchema>;

export default function AboutPage() {
  const queryClient = useQueryClient();
  const [isUploading, setIsUploading] = useState(false);
  const [uploadedImages, setUploadedImages] = useState<string[]>([]);

  const { data, isLoading } = useQuery({
    queryKey: ["about"],
    queryFn: () => aboutApi.get(),
  });

  const {
    register,
    handleSubmit,
    reset,
    setValue,
    formState: { errors },
  } = useForm<AboutFormData>({
    resolver: zodResolver(aboutSchema) as Resolver<AboutFormData>,
    defaultValues: {
      title: "",
      subtitle: "",
      description: "",
      additional_description: "",
      image_url: "",
      years_active: 0,
      active_members: 0,
    },
  });

  useEffect(() => {
    if (data) {
      let parsedImages: string[] = [];
      if (data.image_url) {
        try {
          const parsed = JSON.parse(data.image_url);
          if (Array.isArray(parsed)) {
            parsedImages = parsed.filter(
              (item) => typeof item === "string" && item.trim() !== ""
            );
          } else if (typeof parsed === "string" && parsed.trim() !== "") {
            parsedImages = [parsed];
          }
        } catch {
          if (data.image_url.trim() !== "") {
            parsedImages = [data.image_url];
          }
        }
      }
      setUploadedImages(parsedImages);
      reset({
        title: data.title || "",
        subtitle: data.subtitle || "",
        description: data.description || "",
        additional_description: data.additional_description || "",
        image_url: data.image_url || "",
        years_active: data.years_active ?? 0,
        active_members: data.active_members ?? 0,
      });
    }
  }, [data, reset]);

  const mutation = useMutation({
    mutationFn: aboutApi.update,
    onSuccess: (updated) => {
      queryClient.setQueryData(["about"], updated);
      toast.success("About berhasil disimpan");
    },
    onError: () => {
      toast.error("Gagal menyimpan About");
    },
  });

  const fileInputRef = useRef<HTMLInputElement | null>(null);

  const handleUpload = async (files: FileList | null) => {
    if (!files || files.length === 0) return;

    const currentCount = uploadedImages.length;
    if (currentCount >= 3) {
      toast.error("Maksimal 3 gambar");
      return;
    }

    const filesArray = Array.from(files).slice(0, 3 - currentCount);
    setIsUploading(true);
    try {
      const uploaded: string[] = [];

      for (const file of filesArray) {
        const formData = new FormData();
        formData.append("file", file);
        const res = await apiClient.post("/admin/upload", formData, {
          headers: { "Content-Type": "multipart/form-data" },
        });
        const url = (res.data?.data?.url as string) || "";
        if (!url) throw new Error("URL upload kosong");
        uploaded.push(url);
      }

      const newList = [...uploadedImages, ...uploaded].slice(0, 3);
      setUploadedImages(newList);
      setValue("image_url", newList[0] || "", { shouldValidate: true });
      toast.success(`Berhasil upload ${uploaded.length} gambar`);
    } catch (error) {
      console.error(error);
      toast.error("Upload gagal");
    } finally {
      setIsUploading(false);
    }
  };

  const handleRemoveImage = (url: string) => {
    const filtered = uploadedImages.filter((item) => item !== url);
    setUploadedImages(filtered);
    setValue("image_url", filtered.length > 0 ? JSON.stringify(filtered) : "", {
      shouldValidate: true,
    });
  };

  const onSubmit: SubmitHandler<AboutFormData> = (formData) => {
    const payload = {
      ...formData,
      image_url:
        uploadedImages.length > 0 ? JSON.stringify(uploadedImages) : "",
    };
    mutation.mutate(payload);
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold">About</h1>
          <p className="text-muted-foreground">
            Kelola konten About untuk ditampilkan di website publik
          </p>
        </div>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Konten About</CardTitle>
          <CardDescription>
            Atur teks dan gambar yang akan tampil pada section About di web
            publik.
          </CardDescription>
        </CardHeader>
        <CardContent>
          {isLoading ? (
            <div className="space-y-3">
              <Skeleton className="h-10 w-full" />
              <Skeleton className="h-10 w-full" />
              <Skeleton className="h-24 w-full" />
              <Skeleton className="h-24 w-full" />
              <div className="grid grid-cols-2 gap-4">
                <Skeleton className="h-10 w-full" />
                <Skeleton className="h-10 w-full" />
              </div>
              <Skeleton className="h-10 w-32" />
            </div>
          ) : (
            <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
              <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
                <div className="space-y-4">
                  <div className="space-y-2">
                    <Label htmlFor="title">Judul *</Label>
                    <Input
                      id="title"
                      placeholder="Judul"
                      {...register("title")}
                    />
                    {errors.title && (
                      <p className="text-sm text-destructive">
                        {errors.title.message}
                      </p>
                    )}
                  </div>

                  <div className="space-y-2">
                    <Label htmlFor="subtitle">Subjudul</Label>
                    <Input
                      id="subtitle"
                      placeholder="Subjudul"
                      {...register("subtitle")}
                    />
                    {errors.subtitle && (
                      <p className="text-sm text-destructive">
                        {errors.subtitle.message}
                      </p>
                    )}
                  </div>

                  <div className="space-y-2">
                    <Label htmlFor="description">Deskripsi Utama</Label>
                    <Textarea
                      id="description"
                      placeholder="Deskripsi"
                      rows={4}
                      {...register("description")}
                    />
                    {errors.description && (
                      <p className="text-sm text-destructive">
                        {errors.description.message}
                      </p>
                    )}
                  </div>

                  <div className="space-y-2">
                    <Label htmlFor="additional_description">
                      Deskripsi Tambahan
                    </Label>
                    <Textarea
                      id="additional_description"
                      placeholder="Deskripsi tambahan"
                      rows={4}
                      {...register("additional_description")}
                    />
                    {errors.additional_description && (
                      <p className="text-sm text-destructive">
                        {errors.additional_description.message}
                      </p>
                    )}
                  </div>
                </div>

                <div className="space-y-4">
                  <div className="space-y-2">
                    <Label htmlFor="about-upload-input">Gambar (maks 3)</Label>
                    <input
                      ref={fileInputRef}
                      id="about-upload-input"
                      type="file"
                      accept="image/*"
                      multiple
                      className="hidden"
                      onChange={(e) => {
                        handleUpload(e.target.files);
                        if (fileInputRef.current) {
                          fileInputRef.current.value = "";
                        }
                      }}
                      disabled={isUploading}
                    />

                    <div className="grid grid-cols-3 gap-3">
                      {Array.from({ length: 3 }).map((_, idx) => {
                        const url = uploadedImages[idx];
                        if (url) {
                          return (
                            <div
                              key={url}
                              className="relative h-28 rounded-md border overflow-hidden"
                            >
                              <Image
                                src={url}
                                alt="About preview"
                                fill
                                sizes="150px"
                                className="object-cover"
                                unoptimized
                              />
                              <button
                                type="button"
                                onClick={() => handleRemoveImage(url)}
                                className="absolute top-1 right-1 bg-destructive text-destructive-foreground rounded-full px-2 py-1 text-xs"
                              >
                                ×
                              </button>
                            </div>
                          );
                        }
                        return (
                          <button
                            key={`empty-${idx}`}
                            type="button"
                            onClick={() => fileInputRef.current?.click()}
                            disabled={isUploading}
                            className="h-28 rounded-md border border-dashed border-muted-foreground/40 flex items-center justify-center text-muted-foreground hover:border-primary hover:text-primary transition"
                          >
                            <div className="flex flex-col items-center gap-1">
                              <Plus className="h-5 w-5" />
                              <span className="text-xs">Tambah</span>
                            </div>
                          </button>
                        );
                      })}
                    </div>

                    <p className="text-xs text-muted-foreground">
                      Upload hingga 3 gambar. URL akan terisi otomatis.
                    </p>
                    {errors.image_url && (
                      <p className="text-sm text-destructive">
                        {errors.image_url.message}
                      </p>
                    )}
                  </div>

                  <div className="grid grid-cols-2 gap-4">
                    <div className="space-y-2">
                      <Label htmlFor="years_active">Tahun Berdiri</Label>
                      <Input
                        id="years_active"
                        type="number"
                        min={0}
                        {...register("years_active")}
                      />
                      {errors.years_active && (
                        <p className="text-sm text-destructive">
                          {errors.years_active.message}
                        </p>
                      )}
                    </div>

                    <div className="space-y-2">
                      <Label htmlFor="active_members">Jamaah Aktif</Label>
                      <Input
                        id="active_members"
                        type="number"
                        min={0}
                        {...register("active_members")}
                      />
                      {errors.active_members && (
                        <p className="text-sm text-destructive">
                          {errors.active_members.message}
                        </p>
                      )}
                    </div>
                  </div>
                </div>
              </div>

              <Button type="submit" disabled={mutation.isPending}>
                {mutation.isPending ? "Menyimpan..." : "Simpan"}
              </Button>
            </form>
          )}
        </CardContent>
      </Card>
    </div>
  );
}
