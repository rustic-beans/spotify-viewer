<script lang="ts" setup>
import { type FragmentType, useFragment, graphql } from '@/__generated__';
import { computed } from 'vue';

const ControlFragment = graphql(/* GraphQL */ `
  fragment Control on PlayerState {
    isPlaying
  }
`);

const props = defineProps<{
  fragment: FragmentType<typeof ControlFragment>,
}>();

const controlObj = computed(() => useFragment(ControlFragment, props.fragment));
const isPlaying = computed(() => controlObj.value.isPlaying);
</script>

<template>
  <div>
    <button class="w-16 h-16">
      <svg
        v-if="isPlaying"
        xmlns="http://www.w3.org/2000/svg"
        class="h-16 w-16"
        viewBox="0 0 24 24"
        fill="currentColor"
      >
        <path d="M8 5v14l11-7z" />
      </svg>
      <svg
        v-else
        xmlns="http://www.w3.org/2000/svg"
        class="h-16 w-16"
        viewBox="0 0 24 24"
        fill="currentColor"
      >
        <path d="M6 19h4V5H6v14zm8-14v14h4V5h-4z" />
      </svg>
    </button>
  </div>
</template>
