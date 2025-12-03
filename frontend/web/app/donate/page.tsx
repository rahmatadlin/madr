"use client";

import { motion } from "framer-motion";
import { useQuery } from "@tanstack/react-query";
import { donationApi } from "@/lib/api/donations";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import CountUp from "react-countup";
import { Footer } from "@/components/layouts/footer";

export default function DonatePage() {
  const { data: summary, isLoading } = useQuery({
    queryKey: ["donation-summary"],
    queryFn: () => donationApi.getSummary(),
  });

  return (
    <main className="min-h-screen">
      <div className="container mx-auto px-4 py-20">
        <motion.div
          initial={{ opacity: 0, y: 30 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8 }}
          className="text-center mb-12"
        >
          <h1 className="text-4xl md:text-5xl font-bold text-gray-900 mb-4">
            Donasi untuk Masjid Al-Madr
          </h1>
          <p className="text-lg text-gray-600 max-w-2xl mx-auto">
            Bantu kami membangun dan mengembangkan masjid untuk kemaslahatan
            umat. Setiap donasi Anda sangat berarti.
          </p>
        </motion.div>

        {isLoading ? (
          <div className="text-center py-12">
            <p className="text-gray-600">Memuat data donasi...</p>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12">
            {summary?.per_category.map((category, index) => (
              <motion.div
                key={category.category_id}
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.5, delay: index * 0.1 }}
              >
                <Card className="h-full">
                  <CardHeader>
                    <CardTitle>{category.category}</CardTitle>
                    <CardDescription>
                      Total donasi terkumpul untuk kategori ini
                    </CardDescription>
                  </CardHeader>
                  <CardContent>
                    <div className="text-3xl font-bold text-blue-600 mb-4">
                      <CountUp
                        end={category.amount}
                        duration={2}
                        separator="."
                        prefix="Rp "
                        decimals={0}
                      />
                    </div>
                    <Button className="w-full">Donasi Sekarang</Button>
                  </CardContent>
                </Card>
              </motion.div>
            ))}
          </div>
        )}

        <motion.div
          initial={{ opacity: 0, y: 30 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8, delay: 0.4 }}
          className="max-w-2xl mx-auto bg-blue-50 rounded-lg p-8 text-center"
        >
          <h2 className="text-2xl font-bold text-gray-900 mb-4">
            Cara Donasi
          </h2>
          <p className="text-gray-600 mb-6">
            Untuk melakukan donasi, silakan hubungi kami melalui WhatsApp atau
            email. Kami akan memberikan informasi rekening dan cara transfer.
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Button asChild size="lg" className="bg-green-500 hover:bg-green-600">
              <a href="https://wa.me/62123456789" target="_blank" rel="noopener noreferrer">
                Hubungi via WhatsApp
              </a>
            </Button>
            <Button asChild size="lg" variant="outline">
              <a href="mailto:donasi@masjidalmadr.com">Kirim Email</a>
            </Button>
          </div>
        </motion.div>
      </div>
      <Footer />
    </main>
  );
}

