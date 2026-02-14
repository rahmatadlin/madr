<template>
  <footer class="bg-gray-900 text-gray-300">
    <div class="container mx-auto px-4 py-12">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div>
          <h3 class="text-white text-lg font-semibold mb-4">Masjid Al-Madr</h3>
          <p class="text-sm mb-4">
            Pusat kegiatan keagamaan dan sosial masyarakat yang berkomitmen
            untuk membangun komunitas yang harmonis dan berkualitas.
          </p>
        </div>
        <div>
          <h3 class="text-white text-lg font-semibold mb-4">Kontak</h3>
          <ul class="space-y-3 text-sm">
            <li class="flex items-start gap-2">
              <span class="inline-block w-5 h-5 rounded bg-gray-700 flex-shrink-0" />
              Jl. Masjid Raya No. 123, Jakarta Selatan
            </li>
            <li class="flex items-center gap-2">
              <span class="inline-block w-5 h-5 rounded bg-gray-700 flex-shrink-0" />
              <a href="tel:+62123456789" class="hover:text-white">+62 123 456 789</a>
            </li>
            <li class="flex items-center gap-2">
              <span class="inline-block w-5 h-5 rounded bg-gray-700 flex-shrink-0" />
              <a href="mailto:info@masjidalmadr.com" class="hover:text-white"
                >info@masjidalmadr.com</a
              >
            </li>
          </ul>
        </div>
        <div>
          <h3 class="text-white text-lg font-semibold mb-4">Tautan</h3>
          <ul class="space-y-2 text-sm">
            <li>
              <router-link to="/" class="hover:text-white transition-colors"
                >Beranda</router-link
              >
            </li>
            <li>
              <router-link to="/donate" class="hover:text-white transition-colors"
                >Donasi</router-link
              >
            </li>
            <li>
              <router-link to="/contact" class="hover:text-white transition-colors"
                >Kontak</router-link
              >
            </li>
          </ul>
          <div class="mt-6 flex gap-4">
            <a
              href="https://facebook.com"
              target="_blank"
              rel="noopener noreferrer"
              class="hover:text-white transition-colors"
              aria-label="Facebook"
              >Facebook</a
            >
            <a
              href="https://instagram.com"
              target="_blank"
              rel="noopener noreferrer"
              class="hover:text-white transition-colors"
              aria-label="Instagram"
              >Instagram</a
            >
            <a
              href="https://youtube.com"
              target="_blank"
              rel="noopener noreferrer"
              class="hover:text-white transition-colors"
              aria-label="YouTube"
              >YouTube</a
            >
          </div>
        </div>
      </div>
      <div class="mt-8 pt-8 border-t border-gray-800">
        <a
          href="https://wa.me/62123456789"
          target="_blank"
          rel="noopener noreferrer"
          class="inline-flex items-center gap-2 bg-green-500 hover:bg-green-600 text-white px-6 py-3 rounded-lg transition-colors"
        >
          Hubungi via WhatsApp
        </a>
      </div>
      <div class="mt-8 pt-8 border-t border-gray-800 text-center text-sm">
        <p>&copy; {{ new Date().getFullYear() }} Masjid Al-Madr. All rights reserved.</p>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
</script>
