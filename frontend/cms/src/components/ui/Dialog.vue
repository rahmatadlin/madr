<template>
  <DialogRoot v-model:open="open">
    <DialogPortal>
      <DialogOverlay class="fixed inset-0 z-50 bg-black/50 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0" />
      <DialogContent
        :class="['fixed left-1/2 top-1/2 z-50 w-full max-w-lg -translate-x-1/2 -translate-y-1/2 rounded-lg border bg-white p-6 shadow-lg duration-200 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95', $attrs.class]"
      >
        <DialogTitle v-if="$slots.title" class="text-lg font-semibold">
          <slot name="title" />
        </DialogTitle>
        <DialogDescription v-if="$slots.description" class="mt-1 text-sm text-gray-500">
          <slot name="description" />
        </DialogDescription>
        <div class="mt-4">
          <slot />
        </div>
        <slot name="footer" />
      </DialogContent>
    </DialogPortal>
  </DialogRoot>
</template>

<script setup lang="ts">
import {
  DialogContent,
  DialogDescription,
  DialogOverlay,
  DialogPortal,
  DialogRoot,
  DialogTitle,
} from 'radix-vue'

const open = defineModel<boolean>('open', { default: false })

</script>
