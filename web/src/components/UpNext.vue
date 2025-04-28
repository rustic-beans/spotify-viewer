<script lang=ts setup>
import { defineProps, computed } from 'vue';

const props = defineProps({
  style: {
    type: String,
    default: ''
  }
});

const computedStyle = computed(() => props.style);
</script>

<template>
  <!-- Up Next -->
  <a
    href="https://www.youtube.com/watch?v=dQw4w9WgXcQ&pp=ygUIcmlja3JvbGw%3D"
    target="_blank"
    class="md:visible invisible absolute top-6 right-6 backdrop-blur-lg bg-neutral-800/40 rounded-sm p-3 flex items-center gap-3"
    :style="computedStyle"
  >
    <img
      src="https://picsum.photos/seed/next/60"
      alt="Next Track"
      class="w-12 h-12 rounded-sm object-cover"
    >
    <div class="md:pr-3 max-md:text-center md:text-left max-md:w-full max-md:pt-4 max-md:self-start">
      <div class="text-[0.875em] uppercase tracking-wider opacity-75">UP NEXT</div>
      <div class="font-medium">This section</div>
    </div>
  </a>
</template>

<style type="scss" scoped>
</style>
