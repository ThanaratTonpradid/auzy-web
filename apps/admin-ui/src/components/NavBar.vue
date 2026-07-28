<script setup>
import { computed } from 'vue';
import { useDisplay } from 'vuetify';
import { storeToRefs } from 'pinia';
import { useAppStore } from '../stores/app';
import { useAuthStore } from '../stores/auth';
import { useStaffStore } from '../stores/staff';

const appStore = useAppStore();
const authStore = useAuthStore();
const staffStore = useStaffStore();
const { navIsOpen } = storeToRefs(appStore);
const { mdAndUp } = useDisplay();

const username = computed(
  () => authStore.user?.username || staffStore.profile?.username || 'User'
);

const drawer = computed({
  get: () => (mdAndUp.value ? true : navIsOpen.value),
  set: (value) => {
    if (!mdAndUp.value) {
      appStore.navIsOpen = value;
    }
  },
});

const logout = async () => {
  await authStore.logoutAction();
};
</script>

<template>
  <v-navigation-drawer
    v-model="drawer"
    :permanent="mdAndUp"
    :temporary="!mdAndUp"
    class="admin-drawer"
    width="260"
  >
    <div class="admin-drawer__brand">
      <span class="admin-drawer__mark">A</span>
      <div>
        <div class="admin-drawer__name">{{ $t('common.appName') }}</div>
        <div class="admin-drawer__hint">{{ username }}</div>
      </div>
    </div>

    <v-list nav class="px-2 admin-drawer__nav" color="primary">
      <v-list-item
        :title="$t('menu.dashboard')"
        :to="{ name: 'dashboard' }"
        prepend-icon="mdi-view-dashboard-outline"
        rounded="lg"
      />
      <v-list-item
        :title="$t('menu.profile')"
        :to="{ name: 'profile' }"
        prepend-icon="mdi-account-outline"
        rounded="lg"
      />
      <v-list-item
        v-if="staffStore.canReadStaff"
        :title="$t('menu.staff')"
        :to="{ name: 'staff' }"
        prepend-icon="mdi-account-group-outline"
        rounded="lg"
      />
      <v-list-item
        v-if="staffStore.canReadRoles"
        :title="$t('menu.roles')"
        :to="{ name: 'roles' }"
        prepend-icon="mdi-shield-account-outline"
        rounded="lg"
      />
      <v-list-item
        v-if="staffStore.canReadVisitors"
        :title="$t('menu.visitors')"
        :to="{ name: 'visitor-logs' }"
        prepend-icon="mdi-map-marker-path"
        rounded="lg"
      />
    </v-list>

    <template #append>
      <div class="admin-drawer__footer">
        <v-btn
          block
          variant="text"
          color="error"
          rounded="lg"
          prepend-icon="mdi-logout"
          @click="logout"
        >
          {{ $t('auth.logout') }}
        </v-btn>
      </div>
    </template>
  </v-navigation-drawer>
</template>

<style scoped>
.admin-drawer {
  border-right: 1px solid var(--line) !important;
  background: rgba(247, 250, 248, 0.92) !important;
  backdrop-filter: blur(12px);
}

.admin-drawer__brand {
  display: flex;
  gap: 0.85rem;
  align-items: center;
  padding: 1.25rem 1rem 1rem;
}

.admin-drawer__mark {
  width: 2.35rem;
  height: 2.35rem;
  border-radius: 0.85rem;
  display: grid;
  place-items: center;
  background: var(--accent);
  color: white;
  font-weight: 700;
}

.admin-drawer__name {
  font-weight: 700;
  letter-spacing: -0.02em;
}

.admin-drawer__hint {
  font-size: 0.8rem;
  color: var(--ink-soft);
}

.admin-drawer__nav :deep(.v-list-item--active) {
  background: rgba(31, 122, 102, 0.12) !important;
  color: var(--accent) !important;
}

.admin-drawer__footer {
  padding: 0.75rem;
  border-top: 1px solid var(--line);
}
</style>
