<script lang="ts" setup>
import { ref } from 'vue';

const props = defineProps<{
  currentTime: number;
  duration: number;
}>();

const formatTime = (seconds: number) => {
  const minutes = Math.floor(seconds / 60);
  const remainingSeconds = seconds % 60;
  return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`;
}
</script>

<template>
  <div class="mt-12">
    <div class="relative w-full h-1 bg-neutral-700/50 rounded-full overflow-hidden group cursor-pointer">
      <div
        class="absolute inset-y-0 left-0 bg-white group-hover:bg-green-500 transition-colors"
        :style="{ width: `${(currentTime / duration) * 100}%` }"
      ></div>
    </div>
    <div class="flex justify-between mt-2 text-sm text-gray-400">
      <span>{{ formatTime(currentTime) }}</span>
      <span>{{ formatTime(duration) }}</span>
    </div>
  </div>
</template>
