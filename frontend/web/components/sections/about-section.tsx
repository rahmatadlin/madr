"use client";

import { motion } from "framer-motion";
import Image from "next/image";

export function AboutSection() {
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
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-6">
              Tentang Masjid Al-Madr
            </h2>
            <p className="text-lg text-gray-600 mb-4">
              Masjid Al-Madr adalah pusat kegiatan keagamaan dan sosial yang
              berkomitmen untuk membangun komunitas yang harmonis dan
              berkualitas. Kami menyediakan berbagai program dan kegiatan untuk
              seluruh umat.
            </p>
            <p className="text-lg text-gray-600 mb-4">
              Dengan dukungan dari para jamaah dan donatur, kami terus
              mengembangkan fasilitas dan program yang dapat memberikan manfaat
              lebih luas bagi masyarakat.
            </p>
            <div className="grid grid-cols-2 gap-4 mt-8">
              <div className="bg-blue-50 rounded-lg p-4">
                <h3 className="text-2xl font-bold text-blue-600">15+</h3>
                <p className="text-gray-600">Tahun Berdiri</p>
              </div>
              <div className="bg-green-50 rounded-lg p-4">
                <h3 className="text-2xl font-bold text-green-600">500+</h3>
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
            <Image
              src="/placeholder-masjid.jpg"
              alt="Masjid Al-Madr"
              fill
              className="object-cover"
              placeholder="blur"
              blurDataURL="data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAYEBQYFBAYGBQYHBwYIChAKCgkJChQODwwQFxQYGBcUFhYaHSUfGhsjHBYWICwgIyYnKSopGR8tMC0oMCUoKSj/2wBDAQcHBwoIChMKChMoGhYaKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCj/wAARCAAIAAoDASIAAhEBAxEB/8QAFQABAQAAAAAAAAAAAAAAAAAAAAv/xAAhEAACAQMDBQAAAAAAAAAAAAABAgMABAUGIWGRkqGx0f/EABUBAQEAAAAAAAAAAAAAAAAAAAMF/8QAGhEAAgIDAAAAAAAAAAAAAAAAAAECEgMRkf/aAAwDAQACEQMRAD8AltJagyeH0AthI5xdrLcNM91BF5pX2HaH9bcfaSXWGaRmknyJckliyjqTzSlT54b6bk+h0R//2Q=="
            />
          </motion.div>
        </div>
      </div>
    </section>
  );
}

