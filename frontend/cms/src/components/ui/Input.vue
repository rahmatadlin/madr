<template>
  <input
    :id="id"
    :type="type"
    :value="modelValue"
    :placeholder="placeholder"
    :disabled="disabled"
    :class="inputClass"
    v-bind="$attrs"
    @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue'

defineEmits<{ 'update:modelValue': [value: string] }>()

const props = defineProps<{
  id?: string
  type?: string
  modelValue?: string | number
  placeholder?: string
  disabled?: boolean
  class?: string
}>()
const inputClass = computed(() => {
  const base = 'w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-gray-500 focus:border-gray-500 disabled:opacity-50 disabled:bg-gray-50'
  return [base, props.class].filter(Boolean).join(' ')
})
</script>
