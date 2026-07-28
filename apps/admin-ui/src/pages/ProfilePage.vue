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
  <div class="page-shell">
    <header class="page-header">
      <p class="page-header__eyebrow">{{ $t('common.appName') }}</p>
      <h1 class="page-header__title">{{ $t('profile.title') }}</h1>
      <p class="page-header__subtitle">{{ $t('profile.personalInfo') }}</p>
    </header>

    <section class="surface-panel surface-panel--pad profile-panel">
      <div class="profile-identity">
        <div class="profile-avatar">
          {{ (user.fullname || user.username || 'A').charAt(0).toUpperCase() }}
        </div>
        <div>
          <h2 class="profile-identity__name">{{ user.fullname || user.username || '-' }}</h2>
          <p class="profile-identity__meta">{{ user.roleLabel || '-' }}</p>
        </div>
      </div>

      <div class="profile-grid">
        <div class="profile-field">
          <span class="profile-field__label">{{ $t('auth.username') }}</span>
          <span class="profile-field__value">{{ user.username || '-' }}</span>
        </div>
        <div class="profile-field">
          <span class="profile-field__label">{{ $t('profile.fullname') }}</span>
          <span class="profile-field__value">{{ user.fullname || '-' }}</span>
        </div>
        <div class="profile-field">
          <span class="profile-field__label">{{ $t('menu.roles') }}</span>
          <span class="profile-field__value">{{ user.roleLabel || '-' }}</span>
        </div>
        <div class="profile-field">
          <span class="profile-field__label">{{ $t('staff.admin') }}</span>
          <span class="profile-field__value">
            {{ user.isAdmin ? $t('common.yes') : $t('common.no') }}
          </span>
        </div>
      </div>

      <div class="profile-perms">
        <div class="profile-field__label">{{ $t('menu.permissions') }}</div>
        <div class="profile-perms__list">
          <span
            v-for="code in user.permissions || []"
            :key="code"
            class="meta-chip"
          >
            {{ code }}
          </span>
          <span
            v-if="!(user.permissions || []).length"
            class="profile-field__value"
          >
            —
          </span>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.profile-panel {
  display: grid;
  gap: 1.5rem;
}

.profile-identity {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.profile-avatar {
  width: 3.25rem;
  height: 3.25rem;
  border-radius: 1rem;
  display: grid;
  place-items: center;
  background: var(--accent);
  color: white;
  font-weight: 700;
  font-size: 1.25rem;
}

.profile-identity__name {
  margin: 0;
  font-size: 1.35rem;
  letter-spacing: -0.02em;
}

.profile-identity__meta {
  margin: 0.2rem 0 0;
  color: var(--ink-soft);
}

.profile-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 1rem;
}

.profile-field {
  display: grid;
  gap: 0.3rem;
  padding: 0.9rem 1rem;
  border-radius: 14px;
  border: 1px solid var(--line);
  background: rgba(255, 255, 255, 0.5);
}

.profile-field__label {
  font-size: 0.72rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: var(--ink-soft);
}

.profile-field__value {
  font-weight: 600;
}

.profile-perms__list {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-top: 0.65rem;
}

@media (max-width: 720px) {
  .profile-grid {
    grid-template-columns: 1fr;
  }
}
</style>
