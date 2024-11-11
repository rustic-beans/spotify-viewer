<script lang="ts" setup>
import { computed } from "vue";
import { useGetPlayerStateSubscription } from "@/__generated__/queries";

const { result, loading, error } = useGetPlayerStateSubscription();
</script>

<template>
  <div v-if="loading">
    Loading...
  </div>
  <div v-else-if="error">
    Error: {{ error.message }}
  </div>
  <div v-else>
    {{ result }}
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
