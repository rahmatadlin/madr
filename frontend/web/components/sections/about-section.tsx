"use client";

import { useEffect, useMemo, useState } from "react";
import { AnimatePresence, motion } from "framer-motion";
import Image from "next/image";
import { aboutApi, type AboutContent } from "@/lib/api/about";

const FALLBACK_IMAGE =
  "https://images.unsplash.com/photo-1500530855697-b586d89ba3ee?auto=format&fit=crop&w=1600&q=80";

export function AboutSection() {
  const [about, setAbout] = useState<AboutContent | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [currentIndex, setCurrentIndex] = useState(0);

  useEffect(() => {
    const fetchAbout = async () => {
      try {
        const data = await aboutApi.get();
        setAbout(data);
        setError(null);
      } catch (err) {
        console.error("Failed to fetch about content", err);
        setError("Gagal memuat data profil masjid");
      }
    };

    fetchAbout();
  }, []);

  const images = useMemo(() => {
    if (about?.image_url) {
      try {
        const parsed = JSON.parse(about.image_url);
        if (Array.isArray(parsed)) {
          const arr = parsed.filter(
            (item) => typeof item === "string" && item.trim() !== ""
          );
          if (arr.length > 0) return arr;
        } else if (typeof parsed === "string" && parsed.trim() !== "") {
          return [parsed];
        }
      } catch {
        if (about.image_url.trim() !== "") return [about.image_url];
      }
    }
    return [FALLBACK_IMAGE];
  }, [about?.image_url]);

  // auto-slide
  useEffect(() => {
    if (images.length <= 1) return;
    const timer = setInterval(() => {
      setCurrentIndex((prev) => (prev + 1) % images.length);
    }, 4000);
    return () => clearInterval(timer);
  }, [images.length]);

  const title = about?.title || "Tentang Masjid Al-Madr";
  const subtitle =
    about?.subtitle ||
    "Masjid Al-Madr adalah pusat kegiatan keagamaan dan sosial yang berkomitmen untuk membangun komunitas yang harmonis dan berkualitas.";
  const description =
    about?.description ||
    "Masjid Al-Madr adalah pusat kegiatan keagamaan dan sosial yang berkomitmen untuk membangun komunitas yang harmonis dan berkualitas. Kami menyediakan berbagai program dan kegiatan untuk seluruh umat.";
  const additionalDescription =
    about?.additional_description ||
    "Dengan dukungan dari para jamaah dan donatur, kami terus mengembangkan fasilitas dan program yang dapat memberikan manfaat lebih luas bagi masyarakat.";

  const yearsActive = about?.years_active ?? 15;
  const activeMembers = about?.active_members ?? 500;

  const currentImage = images[currentIndex] || FALLBACK_IMAGE;
  const isRemoteImage = currentImage.startsWith("http");

  return (
    <section className="py-20 bg-white">
      <div className="container mx-auto px-4">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
          <motion.div
            initial={{ opacity: 0, x: -30 }}
            whileInView={{ opacity: 1, x: 0 }}
            viewport={{ once: true }}
            transition={{ duration: 0.8 }}
          >
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-2">
              {title}
            </h2>
            {subtitle && (
              <p className="text-lg text-blue-700 mb-4">{subtitle}</p>
            )}
            <p className="text-lg text-gray-600 mb-4">{description}</p>
            {additionalDescription && (
              <p className="text-lg text-gray-600 mb-4">
                {additionalDescription}
              </p>
            )}
            {error && <p className="text-sm text-red-600 mt-2">{error}</p>}
            <div className="grid grid-cols-2 gap-4 mt-8">
              <div className="bg-blue-50 rounded-lg p-4">
                <h3 className="text-2xl font-bold text-blue-600">
                  {yearsActive}+
                </h3>
                <p className="text-gray-600">Tahun Berdiri</p>
              </div>
              <div className="bg-green-50 rounded-lg p-4">
                <h3 className="text-2xl font-bold text-green-600">
                  {activeMembers}+
                </h3>
                <p className="text-gray-600">Jamaah Aktif</p>
              </div>
            </div>
          </motion.div>

          <motion.div
            initial={{ opacity: 0, x: 30 }}
            whileInView={{ opacity: 1, x: 0 }}
            viewport={{ once: true }}
            transition={{ duration: 0.8 }}
            className="relative h-96 rounded-lg overflow-hidden shadow-xl"
          >
            <AnimatePresence mode="wait">
              <motion.div
                key={currentImage}
                initial={{ opacity: 0, scale: 1.02 }}
                animate={{ opacity: 1, scale: 1 }}
                exit={{ opacity: 0, scale: 0.98 }}
                transition={{ duration: 1.2, ease: "easeInOut" }}
                className="absolute inset-0"
              >
              <Image
                  src={currentImage}
                  alt={title}
                  fill
                  className="object-cover"
                sizes="(max-width: 1024px) 100vw, 50vw"
                unoptimized={isRemoteImage}
                />
              </motion.div>
            </AnimatePresence>
          </motion.div>
        </div>
      </div>
    </section>
  );
}
