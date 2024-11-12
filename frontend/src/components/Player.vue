<script setup lang="ts">
import AlbumCover from '@/components/AlbumCover.vue';
import TrackInfo from '@/components/TrackInfo.vue';
import Controls from '@/components/Controls.vue';
import ProgressBar from '@/components/ProgressBar.vue';
import { type FragmentType, useFragment, graphql } from '@/__generated__';
import { computed } from 'vue';

const PlayerStateFragment = graphql(/* GraphQL */ `
  fragment PlayerState on PlayerState {
    ...Progress
    ...Control
    item {
      ... Track
      album {
        ...Images
      }
    }
  }
`);

const props = defineProps<{
  fragment: FragmentType<typeof PlayerStateFragment>,
}>();

const playerStateObj = computed(() => useFragment(PlayerStateFragment, props.fragment));
const trackObj = computed(() => playerStateObj.value.item);
const imagesObj = computed(() => trackObj.value?.album);
</script>

<template>
  <div class="fixed inset-0 bg-gradient-to-b from-neutral-800 to-neutral-900">
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
          v-if="trackObj"
          :fragment="playerStateObj"
        />
      </div>

      <ProgressBar
        v-if="playerStateObj"
        :fragment="playerStateObj"
      />
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
