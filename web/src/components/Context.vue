<script lang="ts" setup>
import { type FragmentType, useFragment, graphql } from '@/__generated__';
import { computed } from 'vue';

const ContextFragment = graphql(/* GraphQL */ `
  fragment Context on PlayerStateContext {
    name
    type
    href
    imageUrl
  }
`);

const props = defineProps<{
  fragment: FragmentType<typeof ContextFragment>,
}>();

const contextObj = computed(() => useFragment(ContextFragment, props.fragment));
const contextName = computed(() => contextObj.value.name || 'Name could not be found');
const contextType = computed(() => contextObj.value.type.toUpperCase() || 'PLAYING FROM SOMEWHERE');
const contextImage = computed(() => contextObj.value.imageUrl || '/placeholder.png');
</script>

<template>
  <!-- Header -->
  <a
    :href="contextObj.href"
    target="_blank"
  >
    <div class="absolute top-6 left-6 backdrop-blur-lg bg-neutral-800/40 rounded-sm p-3 flex items-center gap-3 ">
      <img
        :src="contextImage"
        alt="Context Image"
        class="w-12 h-12 rounded-sm object-cover"
      >
      <div class="pr-3">
        <div class="text-xs uppercase tracking-wider opacity-75">{{ contextType }}</div>
        <div class="font-medium">{{ contextName }}</div>
      </div>
    </div>
  </a>
</template>
