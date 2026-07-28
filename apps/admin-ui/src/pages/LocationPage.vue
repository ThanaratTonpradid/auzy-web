<script setup>
import { computed, onMounted, ref } from 'vue';
import { getPublicLocationService } from '../services/visitor';
import { appConfig } from '../config/app.config';

const loading = ref(true);
const error = ref('');
const data = ref(null);

const meta = computed(() => data.value?.metadata || {});
const summary = computed(() => {
  const parts = [meta.value.city, meta.value.region, meta.value.country].filter(Boolean);
  return parts.length ? parts.join(', ') : '';
});

const mapsUrl = computed(() => {
  if (meta.value.latitude == null || meta.value.longitude == null) {
    return '';
  }
  return `https://www.openstreetmap.org/?mlat=${meta.value.latitude}&mlon=${meta.value.longitude}#map=10/${meta.value.latitude}/${meta.value.longitude}`;
});

const formatCoord = (value) => {
  if (value === null || value === undefined || Number.isNaN(Number(value))) {
    return '-';
  }
  return Number(value).toFixed(5);
};

const load = async () => {
  loading.value = true;
  error.value = '';
  try {
    data.value = await getPublicLocationService();
  } catch (err) {
    error.value = err.message || 'Failed to load location';
    data.value = null;
  } finally {
    loading.value = false;
  }
};

onMounted(load);
</script>

<template>
  <section class="location-page">
    <div class="location-page__pattern" aria-hidden="true" />
    <div class="location-page__orb location-page__orb--a" aria-hidden="true" />
    <div class="location-page__orb location-page__orb--b" aria-hidden="true" />

    <main class="location-card">
      <header class="location-card__header">
        <RouterLink class="location-card__brand" :to="{ name: 'public-profile' }">
          {{ appConfig.appName }}
        </RouterLink>
        <p class="location-card__eyebrow">{{ $t('locationPage.eyebrow') }}</p>
        <h1 class="location-card__title">{{ $t('locationPage.title') }}</h1>
        <p class="location-card__subtitle">{{ $t('locationPage.subtitle') }}</p>
      </header>

      <div v-if="loading" class="location-card__state">
        <v-progress-circular indeterminate color="primary" size="36" width="3" />
        <p>{{ $t('locationPage.loading') }}</p>
      </div>

      <div v-else-if="error" class="location-card__state location-card__state--error">
        <p>{{ error }}</p>
        <v-btn color="primary" rounded="lg" prepend-icon="mdi-refresh" @click="load">
          {{ $t('locationPage.refresh') }}
        </v-btn>
      </div>

      <template v-else-if="data">
        <div class="location-card__ip">
          <span class="location-card__label">{{ $t('locationPage.ip') }}</span>
          <strong>{{ data.ip || '-' }}</strong>
        </div>

        <div v-if="data.available" class="location-card__summary">
          <span class="location-card__label">{{ $t('locationPage.summary') }}</span>
          <strong>{{ summary || '-' }}</strong>
        </div>
        <v-alert
          v-else
          type="info"
          variant="tonal"
          rounded="lg"
          class="mb-4"
        >
          {{ data.message || $t('locationPage.unavailable') }}
        </v-alert>

        <dl class="location-grid">
          <div>
            <dt>{{ $t('visitors.country') }}</dt>
            <dd>{{ meta.country || '-' }}</dd>
          </div>
          <div>
            <dt>{{ $t('visitors.region') }}</dt>
            <dd>{{ meta.region || '-' }}</dd>
          </div>
          <div>
            <dt>{{ $t('visitors.city') }}</dt>
            <dd>{{ meta.city || '-' }}</dd>
          </div>
          <div>
            <dt>{{ $t('visitors.source') }}</dt>
            <dd>{{ meta.source || '-' }}</dd>
          </div>
          <div>
            <dt>{{ $t('visitors.latitude') }}</dt>
            <dd>{{ formatCoord(meta.latitude) }}</dd>
          </div>
          <div>
            <dt>{{ $t('visitors.longitude') }}</dt>
            <dd>{{ formatCoord(meta.longitude) }}</dd>
          </div>
        </dl>

        <div class="location-card__actions">
          <v-btn
            v-if="mapsUrl"
            color="primary"
            rounded="lg"
            :href="mapsUrl"
            target="_blank"
            rel="noopener noreferrer"
            prepend-icon="mdi-map-marker-radius-outline"
          >
            {{ $t('locationPage.maps') }}
          </v-btn>
          <v-btn
            variant="tonal"
            color="primary"
            rounded="lg"
            prepend-icon="mdi-refresh"
            @click="load"
          >
            {{ $t('locationPage.refresh') }}
          </v-btn>
          <v-btn
            variant="text"
            rounded="lg"
            :to="{ name: 'public-profile' }"
          >
            {{ $t('locationPage.goHome') }}
          </v-btn>
        </div>
      </template>
    </main>
  </section>
</template>

<style scoped>
.location-page {
  position: relative;
  min-height: 100vh;
  display: grid;
  place-items: center;
  padding: 1.5rem;
  overflow: hidden;
  background: #eef2f0;
}

.location-page__pattern {
  position: absolute;
  inset: 0;
  opacity: 0.4;
  background-image:
    radial-gradient(circle at 1px 1px, rgba(18, 32, 28, 0.08) 1px, transparent 0);
  background-size: 22px 22px;
  mask-image: radial-gradient(circle at center, black 35%, transparent 85%);
  pointer-events: none;
}

.location-page__orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(36px);
  pointer-events: none;
  animation: drift 12s ease-in-out infinite alternate;
}

.location-page__orb--a {
  width: 18rem;
  height: 18rem;
  top: -4rem;
  left: -3rem;
  background: rgba(31, 122, 102, 0.2);
}

.location-page__orb--b {
  width: 16rem;
  height: 16rem;
  right: -2rem;
  bottom: 8%;
  background: rgba(216, 239, 232, 0.9);
  animation-delay: -4s;
}

.location-card {
  position: relative;
  z-index: 1;
  width: min(560px, 100%);
  padding: 2rem 1.75rem;
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.92);
  border: 1px solid var(--line);
  box-shadow: var(--shadow-soft);
  animation: rise 520ms var(--motion) both;
}

.location-card__brand {
  display: inline-block;
  margin-bottom: 1.25rem;
  color: var(--ink);
  text-decoration: none;
  font-weight: 700;
  letter-spacing: -0.02em;
}

.location-card__eyebrow {
  margin: 0;
  font-size: 0.75rem;
  letter-spacing: 0.16em;
  text-transform: uppercase;
  color: var(--ink-soft);
}

.location-card__title {
  margin: 0.45rem 0 0;
  font-size: clamp(1.8rem, 4vw, 2.3rem);
  letter-spacing: -0.04em;
  line-height: 1.1;
}

.location-card__subtitle {
  margin: 0.7rem 0 0;
  color: var(--ink-soft);
  line-height: 1.55;
}

.location-card__state {
  display: grid;
  justify-items: center;
  gap: 0.85rem;
  margin-top: 2rem;
  padding: 1.5rem 0;
  color: var(--ink-soft);
  text-align: center;
}

.location-card__state--error {
  color: var(--danger);
}

.location-card__ip,
.location-card__summary {
  display: grid;
  gap: 0.3rem;
  margin-top: 1.35rem;
  padding: 1rem 1.1rem;
  border-radius: 14px;
  border: 1px solid var(--line);
  background: rgba(247, 250, 248, 0.9);
}

.location-card__label {
  font-size: 0.72rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: var(--ink-soft);
}

.location-card__ip strong,
.location-card__summary strong {
  font-size: 1.15rem;
  letter-spacing: -0.02em;
  word-break: break-all;
}

.location-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0.85rem 1rem;
  margin: 1.25rem 0 0;
}

.location-grid dt {
  font-size: 0.72rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: var(--ink-soft);
}

.location-grid dd {
  margin: 0.25rem 0 0;
  font-weight: 600;
}

.location-card__actions {
  display: flex;
  flex-wrap: wrap;
  gap: 0.55rem;
  margin-top: 1.5rem;
}

@keyframes rise {
  from {
    opacity: 0;
    transform: translateY(14px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes drift {
  from {
    transform: translate3d(0, 0, 0) scale(1);
  }
  to {
    transform: translate3d(12px, 18px, 0) scale(1.05);
  }
}

@media (max-width: 720px) {
  .location-grid {
    grid-template-columns: 1fr;
  }

  .location-card__actions {
    flex-direction: column;
  }

  .location-card__actions :deep(.v-btn) {
    width: 100%;
  }
}
</style>
