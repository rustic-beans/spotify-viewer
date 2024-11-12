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
const artistName = computed(() => trackObj.value.album?.artists[0].name);
</script>

<template>
  <!-- Track Info -->
  <div>
    <h1 class="text-6xl font-bold text-white mb-2">{{ trackObj.name }}</h1>
    <p class="text-xl text-gray-400">{{ artistName }}</p>
  </div>
</template>
