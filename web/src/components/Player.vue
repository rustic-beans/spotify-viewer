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
const gmodhaha = Math.floor(Math.random() * 5);
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
      class="object-cover h-full w-full"
      v-if="gmodhaha == 1"
    >
      <img src="/gmoderror.jpg" alt="Gmod Error" class="object-cover h-full w-full" />
    </div>
    <div
      class="object-cover h-full w-full blur-sm "
      :style="{ 'background-color': color }"
    >
      <div class="object-cover h-full w-full bg-linear-to-tr from-neutral-600 to-neutral-900 opacity-25 blur-lg">
      </div>
    </div>
    <div class="absolute bottom-0 left-0 right-0 w-full h-96">
      <div
        class="absolute h-full w-full bg-linear-to-b from-transparent via-neutral-800/00 to-neutral-800/100 backdrop-blur-sm bottom-blur-gradient"
        style=""
      />
      <div class="absolute z-10 bottom-0 left-0 right-0 p-6">
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
  </div>
</template>

<style type="scss" scoped>
.bottom-blur-gradient {
  mask-image: linear-gradient(to bottom, transparent, black 40%, black);
}
</style>
