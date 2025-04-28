<script lang="ts" setup>
import { computed } from 'vue';
import { type FragmentType, useFragment, graphql } from '@/__generated__';

const ImagesFragment = graphql(/* GraphQL */ `
  fragment AlbumImages on Album {
    externalUrls
    imageUrl
  }
`);

const props = defineProps<{
  fragment: FragmentType<typeof ImagesFragment>,
}>();

const imagesObj = computed(() => useFragment(ImagesFragment, props.fragment));

const albumUrl = computed(() => {
  return imagesObj.value.imageUrl;
});

const spotifyUrl = computed(() => {
  return imagesObj.value.externalUrls.spotify;
});
</script>

<template>
  <a
    :href="spotifyUrl"
    target="_blank"
    class="hidden md:block"
  >
    <img
      :src="albumUrl"
      alt="Album Cover"
      class="w-[8rem] h-[8rem] rounded-sm shadow-2xl"
    >
  </a>
</template>
