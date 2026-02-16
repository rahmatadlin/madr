<template>
  <SelectRoot v-model="model" :disabled="disabled">
    <SelectTrigger
      :class="['inline-flex items-center justify-between gap-2 w-full px-3 py-2 rounded-lg border border-gray-300 bg-white hover:bg-gray-50 focus:ring-2 focus:ring-gray-500', $attrs.class]"
    >
      <SelectValue :placeholder="placeholder" />
      <ChevronDownIcon class="h-4 w-4 opacity-50" />
    </SelectTrigger>
    <SelectPortal>
      <SelectContent
        class="z-50 max-h-60 min-w-[8rem] overflow-hidden rounded-lg border bg-white shadow-md"
        position="popper"
        :side-offset="4"
      >
        <SelectViewport>
          <SelectItem
            v-for="opt in options"
            :key="opt.value"
            :value="opt.value"
            class="relative flex cursor-default select-none items-center px-3 py-2 text-sm outline-none focus:bg-gray-100 data-[highlighted]:bg-gray-100"
          >
            <SelectItemText>{{ opt.label }}</SelectItemText>
            <SelectItemIndicator class="absolute left-0 flex h-full w-6 items-center justify-center">
              <CheckIcon class="h-4 w-4" />
            </SelectItemIndicator>
          </SelectItem>
        </SelectViewport>
      </SelectContent>
    </SelectPortal>
  </SelectRoot>
</template>

<script setup lang="ts">
import {
  SelectContent,
  SelectItem,
  SelectItemIndicator,
  SelectItemText,
  SelectPortal,
  SelectRoot,
  SelectTrigger,
  SelectValue,
  SelectViewport,
} from 'radix-vue'
import { CheckIcon, ChevronDownIcon } from '@heroicons/vue/24/outline'

defineProps<{
  modelValue?: string
  placeholder?: string
  disabled?: boolean
  options: { value: string; label: string }[]
}>()

const model = defineModel<string>({ default: '' })
</script>
