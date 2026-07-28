<script setup>
import { computed, onMounted } from 'vue';
import { useStaffStore } from '../stores/staff';
import { useAuthStore } from '../stores/auth';

const staffStore = useStaffStore();
const authStore = useAuthStore();

const profile = computed(() => staffStore.profile);
const permissions = computed(() => staffStore.permissions);

onMounted(async () => {
  if (!profile.value?.id) {
    try {
      const data = await staffStore.getProfileAction();
      authStore.user = data;
    } catch (error) {
      console.error(error);
    }
  }
});
</script>

<template>
  <v-container>
    <h1 class="text-h4 mb-2">{{ $t('dashboard.title') }}</h1>
    <p class="text-body-1 text-medium-emphasis mb-6">
      {{ $t('dashboard.welcome', { name: profile.fullname || profile.username || '-' }) }}
    </p>

    <v-row>
      <v-col cols="12" md="4">
        <v-sheet border rounded class="pa-4">
          <div class="text-overline mb-1">{{ $t('menu.profile') }}</div>
          <div class="text-h6">{{ profile.username || '-' }}</div>
          <div class="text-body-2 text-medium-emphasis">{{ profile.roleLabel || '-' }}</div>
        </v-sheet>
      </v-col>
      <v-col cols="12" md="4">
        <v-sheet border rounded class="pa-4">
          <div class="text-overline mb-1">{{ $t('dashboard.permissionsCount') }}</div>
          <div class="text-h6">{{ permissions.length }}</div>
        </v-sheet>
      </v-col>
      <v-col cols="12" md="4">
        <v-sheet border rounded class="pa-4">
          <div class="text-overline mb-1">{{ $t('dashboard.quickLinks') }}</div>
          <div class="d-flex flex-wrap ga-2 mt-2">
            <v-btn
              v-if="staffStore.canReadStaff"
              size="small"
              variant="tonal"
              :to="{ name: 'staff' }"
            >
              {{ $t('menu.staff') }}
            </v-btn>
            <v-btn
              v-if="staffStore.canReadRoles"
              size="small"
              variant="tonal"
              :to="{ name: 'roles' }"
            >
              {{ $t('menu.roles') }}
            </v-btn>
            <v-btn
              v-if="staffStore.canReadVisitors"
              size="small"
              variant="tonal"
              :to="{ name: 'visitor-logs' }"
            >
              {{ $t('menu.visitors') }}
            </v-btn>
            <v-btn size="small" variant="tonal" :to="{ name: 'profile' }">
              {{ $t('menu.profile') }}
            </v-btn>
          </div>
        </v-sheet>
      </v-col>
    </v-row>
  </v-container>
</template>
