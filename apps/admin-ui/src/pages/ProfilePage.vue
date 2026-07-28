<script setup>
import { computed, onMounted } from 'vue';
import { useAuthStore } from '../stores/auth';
import { useStaffStore } from '../stores/staff';

const authStore = useAuthStore();
const staffStore = useStaffStore();

const user = computed(() => staffStore.profile);

onMounted(async () => {
  if (!user.value?.id) {
    try {
      const profile = await staffStore.getProfileAction();
      authStore.user = profile;
    } catch (error) {
      console.error(error);
    }
  }
});
</script>

<template>
  <v-container>
    <h1 class="text-h4 mb-4">{{ $t('profile.title') }}</h1>

    <v-sheet border rounded class="pa-6">
      <h2 class="text-h6 mb-4">{{ $t('profile.personalInfo') }}</h2>
      <v-row>
        <v-col cols="12" md="6">
          <v-text-field
            :model-value="user.username"
            :label="$t('auth.username')"
            readonly
            variant="outlined"
            density="compact"
          />
        </v-col>
        <v-col cols="12" md="6">
          <v-text-field
            :model-value="user.fullname || '-'"
            :label="$t('profile.fullname')"
            readonly
            variant="outlined"
            density="compact"
          />
        </v-col>
        <v-col cols="12" md="6">
          <v-text-field
            :model-value="user.roleLabel || '-'"
            :label="$t('menu.roles')"
            readonly
            variant="outlined"
            density="compact"
          />
        </v-col>
        <v-col cols="12" md="6">
          <v-text-field
            :model-value="user.isAdmin ? $t('common.yes') : $t('common.no')"
            :label="$t('staff.admin')"
            readonly
            variant="outlined"
            density="compact"
          />
        </v-col>
        <v-col cols="12">
          <div class="text-subtitle-2 mb-2">{{ $t('menu.permissions') }}</div>
          <div class="d-flex flex-wrap ga-2">
            <v-chip
              v-for="code in user.permissions || []"
              :key="code"
              size="small"
              variant="outlined"
            >
              {{ code }}
            </v-chip>
          </div>
        </v-col>
      </v-row>
    </v-sheet>
  </v-container>
</template>
