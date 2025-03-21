<script lang="ts" setup>
import Player from '../components/Player.vue';
import UpNext from '../components/UpNext.vue';
import Context from '../components/Context.vue';

import { computed, watch } from 'vue';
import { useSubscription } from '@vue/apollo-composable';
import { graphql } from '@/__generated__/gql';
import { useRoute } from 'vue-router';

const { result, loading, error } = useSubscription(graphql(/* GraphQL */ `
  subscription playerState {
    playerState {
      ...PlayerState
      context {
        ...Context
      }
    }
  }
`));

const playerState = computed(() => result?.value?.playerState);
const context = computed(() => playerState.value?.context);

const query = useRoute().query;
const bigMode = computed(() => (query.big ?? 'false') === 'true');
watch(() => bigMode.value, () => {
  const size = bigMode.value ? '32px' : '';
  document.getElementsByTagName("html")[0].style["font-size"] = size;
}, { immediate: true });
</script>

<template>
  <main>
    <div v-if="!loading && playerState && context">
      <Player :fragment="playerState" />

      <UpNext />
      <Context :fragment="context!" />
    </div>
    <div
      class="text-center"
      v-else-if="error"
    >
      {{ error.message }}
    </div>
    <div
      class="text-center"
      v-else
    >
      Loading....
    </div>
  </main>
</template>
