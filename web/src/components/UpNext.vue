<script lang=ts setup>
import { type FragmentType, useFragment, graphql } from '@/__generated__';
import { computed } from 'vue';

const UpNextTrackFragment = graphql(/* GraphQL */ `
  fragment UpNextTrack on Track {
    name
    artists {
      name
    }
    album {
      imageUrl
    }
  }
`);

const props = defineProps<{
  fragment: FragmentType<typeof UpNextTrackFragment>,
}>();

const track = computed(() => useFragment(UpNextTrackFragment, props.fragment));
const albumImage = computed(() => track.value.album?.imageUrl || 'https://picsum.photos/seed/next/60');
const trackName = computed(() => track.value.name);
const artistNames = computed(() => track.value.artists.map(a => a.name).join(', '));
</script>

<template>
  <!-- Up Next -->
  <div
    class="absolute md:top-6 md:right-6 md:backdrop-blur-lg md:bg-neutral-800/40 rounded-sm p-3 flex items-center gap-3 max-md:h-36 max-md:w-full max-md:bg-linear-to-t max-md:from-transparent max-md:via-neutral-800/00 max-md:to-neutral-800/100 top-blur-gradient"
  >
    <img
      :src="albumImage"
      alt="Next Track"
      class="hidden md:block w-12 h-12 rounded-sm object-cover"
    >
    <div class="md:pr-3 max-md:text-center md:text-left max-md:w-full max-md:pt-4 max-md:self-start">
      <div class="text-[0.875em] uppercase tracking-wider opacity-75">UP NEXT</div>
      <div class="font-medium">{{ trackName }}</div>
      <div class="text-sm opacity-75">{{ artistNames }}</div>
    </div>
  </div>
</template>

<style type="scss" scoped>
</style>
