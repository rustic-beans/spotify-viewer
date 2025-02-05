<script setup lang="ts">
import AlbumCover from '@/components/AlbumCover.vue';
import Background from '@/components/Background.vue';
import TrackInfo from '@/components/TrackInfo.vue';
import Controls from '@/components/Controls.vue';
import ProgressBar from '@/components/ProgressBar.vue';
import { type FragmentType, useFragment, graphql } from '@/__generated__';
import { computed } from 'vue';

const PlayerStateFragment = graphql(/* GraphQL */ `
  fragment PlayerState on PlayerState {
    ...Progress
    ...Control
    track {
      ... Track
      album {
        ...AlbumImages
        ...BackgroundImages
      }
    }
  }
`);

const props = defineProps<{
  fragment: FragmentType<typeof PlayerStateFragment>,
}>();

const playerStateObj = computed(() => useFragment(PlayerStateFragment, props.fragment));
const trackObj = computed(() => playerStateObj.value.track);
const imagesObj = computed(() => trackObj.value?.album);
const backgroundObj = computed(() => trackObj.value?.album);
//TODO: once we have the artist image from the spotify api call, we can use the obj dominant color for this
var color = "#821271";


</script>

<template>
  <div class="fixed inset-0">
    <Background
      class="object-cover h-full w-full "
      v-if="backgroundObj"
      :fragment="backgroundObj"
    />
    <div
      class="object-cover h-full w-full blur "
      :style="{ 'background-color': color }"
    >
      <div class="object-cover h-full w-full bg-gradient-to-tr from-neutral-600 to-neutral-900 opacity-25 blur-lg">
      </div>
    </div>
    <div class="absolute bottom-0 left-0 right-0 p-12">
      <div class="flex flex-row items-end gap-6">
        <AlbumCover
          v-if="imagesObj"
          :fragment="imagesObj"
        />
        <TrackInfo
          v-if="trackObj"
          :fragment="trackObj"
        />
        <Controls
          class="ml-auto"
          :fragment="playerStateObj"
        />
      </div>

      <ProgressBar :fragment="playerStateObj" />
    </div>
  </div>
</template>

<style scoped>
.from-neutral-800 {
  --tw-gradient-from: rgb(38 38 38);
  --tw-gradient-to: rgb(23 23 23);
  --tw-gradient-stops: var(--tw-gradient-from), var(--tw-gradient-to);
}
</style>
