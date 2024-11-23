<script lang="ts" setup>
import { type FragmentType, useFragment, graphql } from '@/__generated__';
import { computed } from 'vue';
import { useTimestamp } from '@vueuse/core';

const ProgressFragment = graphql(/* GraphQL */ `
  fragment Progress on PlayerState {
    progress_ms
    is_playing
    timestamp
    item {
      duration_ms
    }
  }
`);

const props = defineProps<{
  fragment: FragmentType<typeof ProgressFragment>,
}>();

const progressObj = computed(() => useFragment(ProgressFragment, props.fragment));
const timestamp = useTimestamp({ interval: 1000 });
const timeAgo = computed(() => timestamp.value - (progressObj.value.timestamp));
const progressMs = computed(() => progressObj.value.progress_ms);
const durationMs = computed(() => progressObj.value.item?.duration_ms || 0);
const actualProgressMs = computed(() => {
  let prog = progressMs.value;
  if (progressObj.value.is_playing) {
    prog += timeAgo.value;
  }

  return min(prog, durationMs.value)
});

const progressPercentage = computed(() => {
  return (actualProgressMs.value / durationMs.value) * 100;
});

const formatTime = (ms: number) => {
  const seconds = Math.floor(ms / 1000);
  const minutes = Math.floor(seconds / 60);
  const remainingSeconds = seconds % 60;

  return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`;
}

function min(a: number, b: number): number {
  if (a < b) {
    return a;
  } else {
    return b;
  }
}
</script>

<template>
  <div class="mt-12">
    <div class="relative w-full h-1 bg-neutral-700/50 rounded-full overflow-hidden group cursor-pointer">
      <div
        class="absolute inset-y-0 left-0 bg-white group-hover:bg-green-500 transition-colors"
        :style="{ width: `${progressPercentage}%` }"
      ></div>
    </div>
    <div class="flex justify-between mt-2 text-sm text-gray-400">
      <span>{{ formatTime(actualProgressMs) }}</span>
      <span>{{ formatTime(durationMs) }}</span>
    </div>
  </div>
</template>
