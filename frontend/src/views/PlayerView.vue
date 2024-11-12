<script lang="ts" setup>
import Player from '../components/Player.vue';
import UpNext from '../components/UpNext.vue';
import Context from '../components/Context.vue';

import { computed } from 'vue';
import { useSubscription } from '@vue/apollo-composable';
import { graphql } from '@/__generated__/gql';

const { result, loading, error } = useSubscription(graphql(/* GraphQL */ `
  subscription playerState {
    playerState {
      ...PlayerState
    }
  }
`));
const playerState = computed(() => result?.value?.playerState);
</script>

<template>
  <main>
    <div v-if="!loading && playerState">
      <Player :fragment="playerState" />

      <UpNext />
      <Context />
    </div>
    <div class="text-center" v-else-if="error">
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
