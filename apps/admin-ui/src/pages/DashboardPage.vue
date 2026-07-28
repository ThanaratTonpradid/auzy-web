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
  <div class="page-shell">
    <header class="page-header">
      <p class="page-header__eyebrow">{{ $t('common.appName') }}</p>
      <h1 class="page-header__title">{{ $t('dashboard.title') }}</h1>
      <p class="page-header__subtitle">
        {{ $t('dashboard.welcome', { name: profile.fullname || profile.username || '-' }) }}
      </p>
    </header>

    <div class="dashboard-grid">
      <section class="surface-panel surface-panel--pad dashboard-hero">
        <p class="page-header__eyebrow">{{ $t('menu.profile') }}</p>
        <h2 class="dashboard-hero__name">{{ profile.fullname || profile.username || '-' }}</h2>
        <p class="dashboard-hero__meta">{{ profile.roleLabel || '-' }}</p>
        <div class="dashboard-hero__stats">
          <div>
            <div class="dashboard-hero__stat-label">{{ $t('dashboard.permissionsCount') }}</div>
            <div class="dashboard-hero__stat-value">{{ permissions.length }}</div>
          </div>
          <div>
            <div class="dashboard-hero__stat-label">{{ $t('staff.admin') }}</div>
            <div class="dashboard-hero__stat-value">
              {{ profile.isAdmin ? $t('common.yes') : $t('common.no') }}
            </div>
          </div>
        </div>
      </section>

      <section class="surface-panel surface-panel--pad">
        <p class="page-header__eyebrow">{{ $t('dashboard.quickLinks') }}</p>
        <div class="dashboard-links">
          <RouterLink
            v-if="staffStore.canReadStaff"
            class="dashboard-link"
            :to="{ name: 'staff' }"
          >
            <v-icon icon="mdi-account-group-outline" />
            <span>{{ $t('menu.staff') }}</span>
          </RouterLink>
          <RouterLink
            v-if="staffStore.canReadRoles"
            class="dashboard-link"
            :to="{ name: 'roles' }"
          >
            <v-icon icon="mdi-shield-account-outline" />
            <span>{{ $t('menu.roles') }}</span>
          </RouterLink>
          <RouterLink
            v-if="staffStore.canReadVisitors"
            class="dashboard-link"
            :to="{ name: 'visitor-logs' }"
          >
            <v-icon icon="mdi-map-marker-path" />
            <span>{{ $t('menu.visitors') }}</span>
          </RouterLink>
          <RouterLink class="dashboard-link" :to="{ name: 'profile' }">
            <v-icon icon="mdi-account-outline" />
            <span>{{ $t('menu.profile') }}</span>
          </RouterLink>
        </div>
      </section>
    </div>
  </div>
</template>

<style scoped>
.dashboard-grid {
  display: grid;
  grid-template-columns: 1.4fr 1fr;
  gap: 1rem;
}

.dashboard-hero__name {
  margin: 0.35rem 0 0;
  font-size: clamp(1.6rem, 3vw, 2.1rem);
  letter-spacing: -0.03em;
  line-height: 1.15;
}

.dashboard-hero__meta {
  margin: 0.35rem 0 0;
  color: var(--ink-soft);
}

.dashboard-hero__stats {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
  margin-top: 1.5rem;
  padding-top: 1.25rem;
  border-top: 1px solid var(--line);
}

.dashboard-hero__stat-label {
  font-size: 0.75rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: var(--ink-soft);
}

.dashboard-hero__stat-value {
  margin-top: 0.25rem;
  font-size: 1.45rem;
  font-weight: 700;
}

.dashboard-links {
  display: grid;
  gap: 0.65rem;
  margin-top: 1rem;
}

.dashboard-link {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.9rem 1rem;
  border-radius: 14px;
  border: 1px solid var(--line);
  background: rgba(255, 255, 255, 0.55);
  color: inherit;
  text-decoration: none;
  transition:
    transform var(--motion),
    border-color var(--motion),
    background var(--motion);
}

.dashboard-link:hover {
  transform: translateY(-2px);
  border-color: rgba(31, 122, 102, 0.35);
  background: rgba(216, 239, 232, 0.55);
}

@media (max-width: 960px) {
  .dashboard-grid {
    grid-template-columns: 1fr;
  }
}
</style>
