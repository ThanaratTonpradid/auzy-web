<template>
  <v-snackbar
    v-model="notification.show"
    :timeout="notification.timeout"
    :color="notificationColor"
    location="top right"
    @update:model-value="onClose"
  >
    <div class="d-flex align-center">
      <v-icon :icon="notificationIcon" class="mr-2"></v-icon>
      <span>{{ notification.message }}</span>
    </div>
    <template v-slot:actions>
      <v-btn
        variant="text"
        icon="mdi-close"
        @click="appStore.hideNotification()"
      ></v-btn>
    </template>
  </v-snackbar>
</template>

<script setup>
import { computed } from 'vue';
import { useAppStore } from '../stores/app';

const appStore = useAppStore();

const notification = computed(() => appStore.notification);

const notificationColor = computed(() => {
  switch (notification.value.type) {
    case 'success':
      return 'success';
    case 'error':
      return 'error';
    case 'warning':
      return 'warning';
    case 'info':
    default:
      return 'info';
  }
});

const notificationIcon = computed(() => {
  switch (notification.value.type) {
    case 'success':
      return 'mdi-check-circle';
    case 'error':
      return 'mdi-alert-circle';
    case 'warning':
      return 'mdi-alert';
    case 'info':
    default:
      return 'mdi-information';
  }
});

const onClose = (value) => {
  if (!value) {
    appStore.hideNotification();
  }
};
</script>

