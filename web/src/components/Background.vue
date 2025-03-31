<script lang="ts" setup>
import { computed } from 'vue';
import { type FragmentType, useFragment, graphql } from '@/__generated__';

/* TODO: Change this to be the artist image from a different spotify api call later
 This requires changes to the current graphql types so for now we will just use a placeholder image
(the same as the album cover) ) */
const ImagesFragment = graphql(/* GraphQL */ `
  fragment BackgroundImages on Album {
    images {
      url
    }
  }
`);

const props = defineProps<{
  fragment?: FragmentType<typeof ImagesFragment>,
}>();

const backgroundObj = computed(() => useFragment(ImagesFragment, props.fragment));

const backgroundUrl = computed(() => {
  if (backgroundObj.value === undefined) {
    return Math.floor(Math.random() * 20) == 1 ? "/gmoderror.jpg" : undefined;
  }

  const backgroundValue = backgroundObj.value.images;
  if (backgroundValue.length > 0) {
    return backgroundValue[0].url;
  }

  return undefined;
});

</script>

<template>
  <img
    :src="backgroundUrl"
    v-if="backgroundUrl"
    alt="Background"
    class="w-24 h-24 rounded-sm shadow-2xl"
  >
</template>
