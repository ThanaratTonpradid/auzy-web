<template>
  <v-overlay
    v-model="loading"
    class="align-center justify-center"
    persistent
    :opacity="0.8"
  >
    <div class="text-center">
      <v-progress-circular
        :size="70"
        :width="7"
        color="primary"
        indeterminate
      ></v-progress-circular>
      <div class="mt-4 text-h6">{{ $t('common.loading') }}</div>
    </div>
  </v-overlay>
</template>

<script setup>
import { computed } from 'vue';
import { useAppStore } from '../stores/app';

const appStore = useAppStore();
const loading = computed(() => appStore.loading);
</script>

