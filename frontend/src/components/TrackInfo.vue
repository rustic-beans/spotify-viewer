<script lang="ts" setup>
import { type FragmentType, useFragment, graphql } from '@/__generated__';
import { computed } from 'vue';

const TrackFragment = graphql(/* GraphQL */ `
  fragment Track on FullTrack {
    name
    album {
      artists {
        name
      }
    }
  }
`);

const props = defineProps<{
  fragment: FragmentType<typeof TrackFragment>,
}>();

const trackObj = computed(() => useFragment(TrackFragment, props.fragment));
const artistName = computed(() => {
  const artists = trackObj.value.album?.artists
  if (artists && artists.length > 0 && artists[0]) {
    return artists[0].name;
  }

  return "Unknown Artist";
});
</script>

<template>
  <div>
    <h1 class="text-6xl font-bold text-white mb-2">{{ trackObj.name }}</h1>
    <p class="text-xl text-gray-400">{{ artistName }}</p>
  </div>
</template>
