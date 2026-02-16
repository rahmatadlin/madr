<template>
  <LabelRoot :for="forId" class="block text-sm font-medium text-gray-700">
    <slot />
  </LabelRoot>
</template>

<script setup lang="ts">
import { Label as LabelRoot } from 'radix-vue'

withDefaults(defineProps<{ forId?: string }>(), { forId: undefined })
</script>
