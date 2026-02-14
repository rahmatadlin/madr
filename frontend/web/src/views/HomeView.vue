<template>
  <main class="min-h-screen">
    <HeroSection />
    <AboutSection />
    <DonationCounter />
    <EventsSection />
    <GallerySection />
    <ContactSection />
    <Footer />
  </main>
</template>

<script setup lang="ts">
import HeroSection from '@/components/sections/HeroSection.vue'
import AboutSection from '@/components/sections/AboutSection.vue'
import DonationCounter from '@/components/sections/DonationCounter.vue'
import EventsSection from '@/components/sections/EventsSection.vue'
import GallerySection from '@/components/sections/GallerySection.vue'
import ContactSection from '@/components/sections/ContactSection.vue'
import Footer from '@/components/layouts/Footer.vue'
</script>
