<script setup>
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import { useDisplay } from 'vuetify';
import { useAppStore } from '../stores/app';
import { useAuthStore } from '../stores/auth';
import { useStaffStore } from '../stores/staff';
import { useI18n } from 'vue-i18n';
import LanguageSwitcher from './LanguageSwitcher.vue';

const appStore = useAppStore();
const authStore = useAuthStore();
const staffStore = useStaffStore();
const route = useRoute();
const { mdAndUp } = useDisplay();
const { t } = useI18n();

const title = computed(() => {
  const map = {
    dashboard: 'menu.dashboard',
    profile: 'menu.profile',
    staff: 'menu.staff',
    roles: 'menu.roles',
    'visitor-logs': 'menu.visitors',
  };
  return t(map[route.name] || 'common.appName');
});

const username = computed(
  () => authStore.user?.username || staffStore.profile?.username || ''
);

const toggleNavbar = () => {
  appStore.toggleNavAction();
};
</script>

<template>
  <v-app-bar
    flat
    class="admin-topbar"
    height="64"
    :density="mdAndUp ? 'comfortable' : 'compact'"
  >
    <div class="admin-topbar__inner">
      <div class="admin-topbar__left">
        <v-btn
          v-if="!mdAndUp"
          icon
          variant="text"
          aria-label="Menu"
          @click="toggleNavbar"
        >
          <v-icon icon="mdi-menu" />
        </v-btn>
        <div class="admin-topbar__heading">
          <p v-if="mdAndUp" class="admin-topbar__eyebrow">{{ $t('common.appName') }}</p>
          <h1 class="admin-topbar__title">{{ title }}</h1>
        </div>
      </div>

      <div class="admin-topbar__actions">
        <LanguageSwitcher />
        <v-btn
          class="admin-topbar__site"
          variant="tonal"
          color="primary"
          size="small"
          rounded="lg"
          :to="{ name: 'public-profile' }"
          prepend-icon="mdi-open-in-new"
        >
          {{ $t('common.appName') }}
        </v-btn>
        <div v-if="username && mdAndUp" class="admin-topbar__user">
          <span class="admin-topbar__user-mark">
            {{ username.charAt(0).toUpperCase() }}
          </span>
          <span class="admin-topbar__user-name">{{ username }}</span>
        </div>
      </div>
    </div>
  </v-app-bar>
</template>

<style scoped>
.admin-topbar {
  background: rgba(247, 250, 248, 0.78) !important;
  backdrop-filter: blur(14px);
  border-bottom: 1px solid var(--line) !important;
}

.admin-topbar :deep(.v-toolbar__content) {
  padding: 0 1rem;
}

.admin-topbar__inner {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  width: 100%;
  min-width: 0;
}

.admin-topbar__left {
  display: flex;
  align-items: center;
  gap: 0.35rem;
  min-width: 0;
}

.admin-topbar__heading {
  min-width: 0;
}

.admin-topbar__eyebrow {
  margin: 0;
  font-size: 0.68rem;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: var(--ink-soft);
  line-height: 1.2;
}

.admin-topbar__title {
  margin: 0;
  font-size: 1.05rem;
  font-weight: 700;
  letter-spacing: -0.02em;
  line-height: 1.2;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.admin-topbar__actions {
  display: flex;
  align-items: center;
  gap: 0.65rem;
  flex-shrink: 0;
}

.admin-topbar__user {
  display: flex;
  align-items: center;
  gap: 0.55rem;
  padding: 0.28rem 0.7rem 0.28rem 0.28rem;
  border-radius: 999px;
  border: 1px solid var(--line);
  background: rgba(255, 255, 255, 0.65);
}

.admin-topbar__user-mark {
  width: 1.7rem;
  height: 1.7rem;
  border-radius: 50%;
  display: grid;
  place-items: center;
  background: var(--accent-soft);
  color: var(--accent);
  font-size: 0.75rem;
  font-weight: 700;
}

.admin-topbar__user-name {
  font-size: 0.85rem;
  font-weight: 600;
  max-width: 8rem;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

@media (max-width: 600px) {
  .admin-topbar__site {
    display: none;
  }

  .admin-topbar__actions {
    gap: 0.4rem;
  }

  .admin-topbar :deep(.v-toolbar__content) {
    padding: 0 0.65rem;
  }

  .admin-topbar__title {
    font-size: 0.98rem;
  }
}
</style>
