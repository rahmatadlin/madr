"use client";

import { useEffect, useMemo, useState } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { bannerApi } from "@/lib/api/banners";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { toast } from "sonner";
import { Skeleton } from "@/components/ui/skeleton";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

const apiBaseEnv = (process.env.NEXT_PUBLIC_API_URL || "").replace(/\/$/, "");
const apiBaseWithoutApi = apiBaseEnv.replace(/\/api\/v1\/?$/, "");
const defaultBase = apiBaseWithoutApi || "http://localhost:8080";
const resolveMediaUrl = (url?: string | null) => {
  if (!url) return null;
  return url.startsWith("http") ? url : `${defaultBase}/uploads/${url}`;
};

export default function BannerPage() {
  const queryClient = useQueryClient();

  const limit = 1;
  const { data, isLoading } = useQuery({
    queryKey: ["banners", limit, 0],
    queryFn: () => bannerApi.getAll(limit, 0),
  });

  const banner = data?.data?.[0];
  const initialTitle = banner?.title ?? "";
  const initialType = banner?.type ?? "image";
  const initialPreview = resolveMediaUrl(banner?.media_url);

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold">Banner</h1>
        <p className="text-muted-foreground">
          Kelola satu banner utama untuk Hero (gambar atau video).
        </p>
      </div>

      {isLoading ? (
        <div className="space-y-2">
          {[...Array(3)].map((_, i) => (
            <Skeleton key={i} className="h-12 w-full" />
          ))}
          <Skeleton className="h-48 w-full" />
        </div>
      ) : (
        <BannerForm
          key={banner?.id ?? "new"}
          bannerId={banner?.id}
          defaultTitle={initialTitle}
          defaultType={initialType}
          defaultPreview={initialPreview}
          onSaved={() => {
            queryClient.invalidateQueries({ queryKey: ["banners"] });
            queryClient.invalidateQueries({ queryKey: ["stats"] });
          }}
        />
      )}
    </div>
  );
}

type BannerFormProps = {
  bannerId?: number;
  defaultTitle: string;
  defaultType: "image" | "video";
  defaultPreview: string | null;
  onSaved: () => void;
};

function BannerForm({
  bannerId,
  defaultTitle,
  defaultType,
  defaultPreview,
  onSaved,
}: BannerFormProps) {
  const [title, setTitle] = useState(defaultTitle);
  const [type, setType] = useState<"image" | "video">(defaultType);
  const [file, setFile] = useState<File | null>(null);

  const accept = useMemo(
    () => (type === "video" ? "video/mp4" : "image/*"),
    [type]
  );

  const previewUrl = useMemo(() => {
    if (file) {
      return URL.createObjectURL(file);
    }
    return defaultPreview;
  }, [file, defaultPreview]);

  useEffect(() => {
    if (!file) return;
    return () => URL.revokeObjectURL(previewUrl || "");
  }, [file, previewUrl]);

  const saveMutation = useMutation({
    mutationFn: async () => {
      if (bannerId) {
        return bannerApi.update(bannerId, {
          title,
          type,
          file: file ?? undefined,
        });
      }
      return bannerApi.create({
        title,
        type,
        file: file ?? undefined,
      });
    },
    onSuccess: () => {
      toast.success("Banner berhasil disimpan");
      setFile(null);
      onSaved();
    },
    onError: (err: unknown) => {
      const maybeMsg = (err as { response?: { data?: { error?: string } } })
        ?.response?.data?.error;
      const msg = maybeMsg || "Gagal menyimpan banner";
      toast.error(msg);
    },
  });

  const isSaveDisabled =
    !title.trim() || (!bannerId && !file) || saveMutation.isPending;

  return (
    <div className="space-y-4 rounded-lg border p-4">
      <div className="space-y-2">
        <Label htmlFor="title">Judul</Label>
        <Input
          id="title"
          placeholder="Judul banner"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          disabled={saveMutation.isPending}
        />
      </div>
      <div className="space-y-2">
        <Label>Jenis Media</Label>
        <Select
          value={type}
          onValueChange={(v) => setType(v as "image" | "video")}
          disabled={saveMutation.isPending}
        >
          <SelectTrigger>
            <SelectValue placeholder="Pilih jenis" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="image">Image</SelectItem>
            <SelectItem value="video">Video (mp4)</SelectItem>
          </SelectContent>
        </Select>
      </div>
      <div className="space-y-2">
        <Label>File {type === "video" ? "Video" : "Gambar"}</Label>
        <Input
          type="file"
          accept={accept}
          onChange={(e) => setFile(e.target.files?.[0] || null)}
          disabled={saveMutation.isPending}
        />
        {type === "image" && previewUrl && (
          <div className="relative mt-2 h-48 w-full overflow-hidden rounded-lg border bg-muted/30">
            {/* eslint-disable-next-line @next/next/no-img-element */}
            <img
              src={previewUrl}
              alt="Preview"
              className="h-full w-full object-cover"
            />
          </div>
        )}
        {type === "video" && previewUrl && (
          <div className="mt-2 space-y-2 rounded-lg border bg-muted/30 p-3">
            <video
              src={previewUrl}
              controls
              className="w-full rounded"
              poster={previewUrl}
            />
            {file && (
              <p className="text-sm text-muted-foreground">
                File baru: {file.name}
              </p>
            )}
          </div>
        )}
        {type === "video" && !previewUrl && file && (
          <p className="text-sm text-muted-foreground">
            File baru: {file.name}
          </p>
        )}
      </div>
      <div className="flex justify-end">
        <Button onClick={() => saveMutation.mutate()} disabled={isSaveDisabled}>
          {saveMutation.isPending ? "Menyimpan..." : "Simpan"}
        </Button>
      </div>
      {!bannerId && (
        <p className="text-sm text-muted-foreground">
          Belum ada banner. Unggah file untuk membuat banner pertama.
        </p>
      )}
    </div>
  );
}
