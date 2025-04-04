<script lang="ts" setup>
import { type FragmentType, useFragment, graphql } from '@/__generated__';
import { computed } from 'vue';

const TrackFragment = graphql(/* GraphQL */ `
  fragment Track on Track {
    name
    externalUrls
    artists {
      name
      externalUrls
    }
  }
`);

const props = defineProps<{
  fragment: FragmentType<typeof TrackFragment>,
}>();

const trackObj = computed(() => useFragment(TrackFragment, props.fragment));
const artists = computed(() => {
  return trackObj.value.artists.map((artist) => {
    return {
      name: artist.name,
      link: artist.externalUrls.spotify,
    };
  });
});

const trackLink = computed(() => {
  return trackObj.value.externalUrls?.spotify;
});
</script>

<template>
  <div>
    <a
      :href="trackLink"
      target="_blank"
    >
      <h1 class="text-6xl font-bold text-white mb-2">{{ trackObj.name }}</h1>
    </a>

    <p class="text-[2rem] text-gray-400">
      <template
        v-for="(artist, index) in artists"
        :key="artist.name!"
      >
        <template v-if="index > 0">, </template>
        <a
          :href="artist.link"
          target="_blank"
        >
          {{ artist.name }}
        </a>
      </template>
    </p>
  </div>
</template>
