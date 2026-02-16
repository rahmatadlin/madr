<template>
  <section class="py-20 bg-white">
    <div class="container mx-auto px-4">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
        <div>
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-2">
            {{ title }}
          </h2>
          <p v-if="subtitle" class="text-lg text-blue-700 mb-4">
            {{ subtitle }}
          </p>
          <p class="text-lg text-gray-600 mb-4">{{ description }}</p>
          <p v-if="additionalDescription" class="text-lg text-gray-600 mb-4">
            {{ additionalDescription }}
          </p>
          <p v-if="error" class="text-sm text-red-600 mt-2">{{ error }}</p>
          <div class="grid grid-cols-2 gap-4 mt-8">
            <div class="bg-blue-50 rounded-lg p-4">
              <h3 class="text-2xl font-bold text-blue-600">
                {{ yearsActive }}+
              </h3>
              <p class="text-gray-600">Tahun Berdiri</p>
            </div>
            <div class="bg-green-50 rounded-lg p-4">
              <h3 class="text-2xl font-bold text-green-600">
                {{ activeMembers }}+
              </h3>
              <p class="text-gray-600">Jamaah Aktif</p>
            </div>
          </div>
        </div>
        <div class="relative h-96 rounded-lg overflow-hidden shadow-xl">
          <Transition name="slide-fade" mode="out-in">
            <img
              :key="currentImage"
              :src="currentImage"
              :alt="title"
              class="absolute inset-0 w-full h-full object-cover"
            />
          </Transition>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from "vue";
import { aboutApi } from "@/api/about";
import type { AboutContent } from "@/api/about";

const FALLBACK =
  "https://images.unsplash.com/photo-1500530855697-b586d89ba3ee?auto=format&fit=crop&w=1600&q=80";
const about = ref<AboutContent | null>(null);
const error = ref<string | null>(null);
const currentIndex = ref(0);

const images = computed(() => {
  if (!about.value?.image_url) return [FALLBACK];
  try {
    const p = JSON.parse(about.value.image_url);
    if (Array.isArray(p)) {
      const arr = p.filter((x) => typeof x === "string" && x.trim() !== "");
      if (arr.length) return arr;
    } else if (typeof p === "string" && p.trim() !== "") return [p];
  } catch {
    if (about.value.image_url.trim() !== "") return [about.value.image_url];
  }
  return [FALLBACK];
});

const currentImage = computed(
  () => images.value[currentIndex.value] || FALLBACK,
);
const title = computed(
  () => about.value?.title || "Tentang Masjid Agung Discovery Residence",
);
const subtitle = computed(() => about.value?.subtitle || "");
const description = computed(() => about.value?.description || "");
const additionalDescription = computed(
  () => about.value?.additional_description || "",
);
const yearsActive = computed(() => about.value?.years_active ?? 15);
const activeMembers = computed(() => about.value?.active_members ?? 500);

let t: ReturnType<typeof setInterval> | null = null;
onMounted(async () => {
  try {
    about.value = await aboutApi.get();
    error.value = null;
  } catch (e) {
    error.value = "Gagal memuat data profil masjid";
  }
  if (images.value.length > 1)
    t = setInterval(() => {
      currentIndex.value = (currentIndex.value + 1) % images.value.length;
    }, 4000);
});
onUnmounted(() => {
  if (t) clearInterval(t);
});
</script>

<style scoped>
.slide-fade-enter-active,
.slide-fade-leave-active {
  transition:
    opacity 1.2s ease,
    transform 1.2s ease;
}
.slide-fade-enter-from {
  opacity: 0;
  transform: scale(1.02);
}
.slide-fade-leave-to {
  opacity: 0;
  transform: scale(0.98);
}
</style>
