import { HeroSection } from "@/components/sections/hero-section";
import { AboutSection } from "@/components/sections/about-section";
import { DonationCounter } from "@/components/sections/donation-counter";
import { EventsSection } from "@/components/sections/events-section";
import { GallerySection } from "@/components/sections/gallery-section";
import { ContactSection } from "@/components/sections/contact-section";
import { Footer } from "@/components/layouts/footer";

export default function Home() {
  return (
    <main className="min-h-screen">
      <HeroSection />
      <AboutSection />
      <DonationCounter />
      <EventsSection />
      <GallerySection />
      <ContactSection />
      <Footer />
    </main>
  );
}
