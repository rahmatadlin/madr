import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { QueryProvider } from "@/lib/providers/query-provider";

const inter = Inter({
  subsets: ["latin"],
  variable: "--font-inter",
});

export const metadata: Metadata = {
  title: "Masjid Al-Madr - Pusat Kegiatan Keagamaan dan Sosial",
  description:
    "Masjid Al-Madr adalah pusat kegiatan keagamaan dan sosial yang berkomitmen untuk membangun komunitas yang harmonis dan berkualitas.",
  keywords: ["masjid", "islam", "donasi", "kegiatan keagamaan", "masyarakat"],
  authors: [{ name: "Masjid Al-Madr" }],
  openGraph: {
    title: "Masjid Al-Madr - Pusat Kegiatan Keagamaan dan Sosial",
    description:
      "Masjid Al-Madr adalah pusat kegiatan keagamaan dan sosial yang berkomitmen untuk membangun komunitas yang harmonis dan berkualitas.",
    type: "website",
    locale: "id_ID",
    siteName: "Masjid Al-Madr",
  },
  twitter: {
    card: "summary_large_image",
    title: "Masjid Al-Madr",
    description:
      "Pusat kegiatan keagamaan dan sosial masyarakat yang berkomitmen untuk membangun komunitas yang harmonis.",
  },
  robots: {
    index: true,
    follow: true,
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="id">
      <body className={`${inter.variable} font-sans antialiased`}>
        <QueryProvider>{children}</QueryProvider>
      </body>
    </html>
  );
}
