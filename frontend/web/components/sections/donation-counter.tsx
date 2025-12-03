"use client";

import { motion } from "framer-motion";
import { useQuery } from "@tanstack/react-query";
import { donationApi } from "@/lib/api/donations";
import CountUp from "react-countup";
import { Button } from "@/components/ui/button";
import Link from "next/link";

export function DonationCounter() {
  const { data: summary, isLoading } = useQuery({
    queryKey: ["donation-summary"],
    queryFn: () => donationApi.getSummary(),
    refetchInterval: 60000, // Refetch every minute
  });

  if (isLoading) {
    return (
      <section className="py-20 bg-gradient-to-br from-blue-50 to-indigo-100">
        <div className="container mx-auto px-4">
          <div className="text-center">
            <p className="text-gray-600">Memuat data donasi...</p>
          </div>
        </div>
      </section>
    );
  }

  const formatCurrency = (amount: number) => {
    return new Intl.NumberFormat("id-ID", {
      style: "currency",
      currency: "IDR",
      minimumFractionDigits: 0,
      maximumFractionDigits: 0,
    }).format(amount);
  };

  return (
    <section className="py-20 bg-gradient-to-br from-blue-50 to-indigo-100">
      <div className="container mx-auto px-4">
        <motion.div
          initial={{ opacity: 0, y: 30 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ duration: 0.8 }}
          className="text-center mb-12"
        >
          <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
            Total Donasi Terkumpul
          </h2>
          <p className="text-lg text-gray-600">
            Mari bersama-sama membangun masjid yang lebih baik
          </p>
        </motion.div>

        {/* Total Amount */}
        <motion.div
          initial={{ opacity: 0, scale: 0.9 }}
          whileInView={{ opacity: 1, scale: 1 }}
          viewport={{ once: true }}
          transition={{ duration: 0.8 }}
          className="text-center mb-12"
        >
          <div className="inline-block bg-white rounded-2xl shadow-xl p-8 md:p-12">
            <p className="text-sm md:text-base text-gray-600 mb-2">
              Total Donasi
            </p>
            <div className="text-4xl md:text-6xl font-bold text-blue-600 mb-2">
              <CountUp
                end={summary?.total_amount || 0}
                duration={2}
                separator="."
                prefix="Rp "
                decimals={0}
              />
            </div>
            <p className="text-sm text-gray-500">
              dari {summary?.total_transactions || 0} transaksi
            </p>
          </div>
        </motion.div>

        {/* Per Category */}
        {summary?.per_category && summary.per_category.length > 0 && (
          <motion.div
            initial={{ opacity: 0, y: 30 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ duration: 0.8, delay: 0.2 }}
            className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12"
          >
            {summary.per_category.map((category, index) => (
              <motion.div
                key={category.category_id}
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                viewport={{ once: true }}
                transition={{ duration: 0.5, delay: index * 0.1 }}
                className="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow"
              >
                <h3 className="text-lg font-semibold text-gray-900 mb-2">
                  {category.category}
                </h3>
                <p className="text-2xl font-bold text-blue-600">
                  <CountUp
                    end={category.amount}
                    duration={2}
                    separator="."
                    prefix="Rp "
                    decimals={0}
                  />
                </p>
              </motion.div>
            ))}
          </motion.div>
        )}

        {/* CTA Button */}
        <motion.div
          initial={{ opacity: 0 }}
          whileInView={{ opacity: 1 }}
          viewport={{ once: true }}
          transition={{ duration: 0.8, delay: 0.4 }}
          className="text-center"
        >
          <Button
            asChild
            size="lg"
            className="bg-blue-600 hover:bg-blue-700 text-white px-8 py-6 text-lg"
          >
            <Link href="/donate">Donasi Sekarang</Link>
          </Button>
        </motion.div>
      </div>
    </section>
  );
}

