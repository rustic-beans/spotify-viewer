<script lang="ts" setup>
import { computed } from 'vue';
import { type FragmentType, useFragment, graphql } from '@/__generated__';

const ImagesFragment = graphql(/* GraphQL */ `
  fragment Images on SimpleAlbum {
    external_urls
    images {
      url
    }
  }
`);

const props = defineProps<{
  fragment: FragmentType<typeof ImagesFragment>,
}>();

const imagesObj = computed(() => useFragment(ImagesFragment, props.fragment));

const albumUrl = computed(() => {
  const imageValue = imagesObj.value.images;
  if (imageValue.length > 0) {
    return imageValue[0].url;
  }

  return 'https://i.scdn.co/image/ab67616d0000b273a545e3a3e6cf0cc009297553';
});

const spotifyUrl = computed(() => {
  return imagesObj.value.external_urls.spotify;
});
</script>

<template>
  <a
    :href="spotifyUrl"
    target="_blank"
  >
    <img
      :src="albumUrl"
      alt="Album Cover"
      class="w-24 h-24 rounded shadow-2xl"
    >
  </a>
</template>
