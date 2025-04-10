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
  <div class="hidden md:block">
    <div :class="{ boxContainer: true, animated: isPlaying }">
      <div class="box box1"></div>
      <div class="box box2"></div>
      <div class="box box3"></div>
      <div class="box box4"></div>
      <div class="box box5"></div>
    </div>
  </div>
</template>

<style scoped>
@keyframes quiet {
  25%{
    transform: scaleY(.6);
  }
  50%{
    transform: scaleY(.4);
  }
  75%{
    transform: scaleY(.8);
  }
}

@keyframes normal {
  25%{
    transform: scaleY(1);
  }
  50%{
    transform: scaleY(.4);
  }
  75%{
    transform: scaleY(.6);
  }
}

@keyframes loud {
  25%{
    transform: scaleY(1);
  }
  50%{
    transform: scaleY(.4);
  }
  75%{
    transform: scaleY(1.2);
  }
}

body{
  display: flex;
  justify-content: center;
  background: black;
  margin: 0;padding: 0;
  align-items: center;
  height: 100vh;
}

.boxContainer {
  margin-right: 2rem;
  display: flex;
  justify-content: space-between;
  height: 4rem;
  --boxSize: 0.5rem;
  --gutter: 0.25rem;
  width: calc((var(--boxSize) + var(--gutter)) * 5);
}

.box {
  transform: scaleY(.4);
  height: 100%;
  width: var(--boxSize);
  background: #fff;
  animation-duration: 1.2s;
  animation-timing-function: ease-in-out;
  animation-iteration-count: infinite;
  border-radius: 0.5rem;
}

.boxContainer:not(.animated) .box {
  animation-play-state: paused;
}

.animated .box1{ animation-name: quiet; }
.animated .box2{ animation-name: normal; }
.animated .box3{ animation-name: quiet; }
.animated .box4{ animation-name: loud; }
.animated .box5{ animation-name: quiet; }
</style>
