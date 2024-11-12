<script lang="ts" setup>
import { computed } from 'vue';
import { type FragmentType, useFragment, graphql } from '@/__generated__';

const ImagesFragment = graphql(/* GraphQL */ `
  fragment Images on SimpleAlbum {
    images {
      url
    }
  }
`);

const props = defineProps<{
  fragment: FragmentType<typeof ImagesFragment>,
}>();

const images = computed(() => useFragment(ImagesFragment, props.fragment));

const albumUrl = computed(() => {
  let imageValue = images.value.images;
  if (imageValue && imageValue.length > 0 && imageValue[0]) {
    const image = imageValue[0];
    if (image.url) {
      return image.url;
    }
  }

  return 'https://i.scdn.co/image/ab67616d0000b273a545e3a3e6cf0cc009297553';
});
</script>

<template>
  <img
    :src="albumUrl"
    alt="Album Cover"
    class="w-24 h-24 rounded shadow-2xl"
  >
</template>
