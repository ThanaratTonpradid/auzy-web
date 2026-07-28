<script setup>
import { appConfig } from '../config/app.config';
</script>

<template>
  <section class="not-found">
    <div class="not-found__pattern" aria-hidden="true" />
    <div class="not-found__card">
      <p class="not-found__code">404</p>
      <h1 class="not-found__title">{{ $t('notFound.title') }}</h1>
      <p class="not-found__copy">{{ $t('notFound.message') }}</p>
      <div class="not-found__actions">
        <v-btn color="primary" rounded="lg" :to="{ name: 'public-profile' }">
          {{ $t('notFound.goHome') }}
        </v-btn>
        <v-btn variant="text" rounded="lg" :to="{ name: 'login' }">
          {{ $t('auth.login') }}
        </v-btn>
      </div>
      <p class="not-found__brand">{{ appConfig.appName }}</p>
    </div>
  </section>
</template>

<style scoped>
.not-found {
  position: relative;
  min-height: 100vh;
  display: grid;
  place-items: center;
  padding: 1.5rem;
  overflow: hidden;
  background: #eef2f0;
}

.not-found__pattern {
  position: absolute;
  inset: 0;
  opacity: 0.45;
  background-image:
    radial-gradient(circle at 1px 1px, rgba(18, 32, 28, 0.08) 1px, transparent 0);
  background-size: 22px 22px;
  mask-image: radial-gradient(circle at center, black 35%, transparent 85%);
  pointer-events: none;
}

.not-found__card {
  position: relative;
  z-index: 1;
  width: min(420px, 100%);
  padding: 2.25rem 1.75rem;
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.92);
  border: 1px solid var(--line);
  box-shadow: var(--shadow-soft);
  text-align: center;
  animation: rise 480ms var(--motion) both;
}

.not-found__code {
  margin: 0;
  font-size: clamp(3.5rem, 10vw, 4.5rem);
  font-weight: 700;
  letter-spacing: -0.06em;
  line-height: 1;
  color: var(--accent);
}

.not-found__title {
  margin: 0.85rem 0 0;
  font-size: 1.45rem;
  font-weight: 700;
  letter-spacing: -0.03em;
}

.not-found__copy {
  margin: 0.65rem 0 0;
  color: var(--ink-soft);
  line-height: 1.6;
}

.not-found__actions {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 0.5rem;
  margin-top: 1.5rem;
}

.not-found__brand {
  margin: 1.5rem 0 0;
  font-size: 0.75rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
  color: rgba(61, 82, 74, 0.55);
}

@keyframes rise {
  from {
    opacity: 0;
    transform: translateY(12px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
