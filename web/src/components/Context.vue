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
const contextType = computed(() => 'PLAYING FROM ' + contextObj.value.type.toUpperCase() || 'SOMEWHERE');
const contextImage = computed(() => contextObj.value.imageUrl || '/placeholder.png');

</script>

<template>
  <!-- Header -->
  <a
    :href="contextObj.href"
    target="_blank"
    class="absolute md:top-6 md:left-6 md:backdrop-blur-lg md:bg-neutral-800/40 rounded-sm p-3 flex items-center gap-3 max-md:h-36 max-md:w-full max-md:bg-linear-to-t max-md:from-transparent max-md:via-neutral-800/00 max-md:to-neutral-800/100 top-blur-gradient"
  >
    <img
      :src="contextImage"
      alt="Context Image"
      class="hidden md:block w-12 h-12 rounded-sm object-cover"
    >
    <div class="md:pr-3 max-md:text-center md:text-left max-md:w-full max-md:pt-4 max-md:self-start">
      <div class="text-[0.875em] uppercase tracking-wider opacity-75" >{{ contextType }}</div>
      <div class="font-medium">{{ contextName }}</div>
    </div>
  </a>
</template>

<style type="scss" scoped>
.top-blur-gradient {

  /* Apparently @media cannot use variables, so this is hardcoded */
  @media (width < 48rem) {
    mask-image: linear-gradient(to top, transparent, black 40%, black);
  }
}
</style>
