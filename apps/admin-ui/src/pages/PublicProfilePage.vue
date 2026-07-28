<script setup>
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { recordVisitService } from '../services/visitor';
import { appConfig } from '../config/app.config';

const router = useRouter();

const profile = {
  name: import.meta.env.VITE_PROFILE_NAME || 'Auzy',
  title: import.meta.env.VITE_PROFILE_TITLE || 'Personal profile',
  bio:
    import.meta.env.VITE_PROFILE_BIO ||
    'A quiet corner of the internet. Thanks for stopping by.',
  links: [
    {
      label: 'Admin',
      href: '/pub/login',
      internal: true,
    },
  ],
};

const goInternal = (path) => {
  router.push(path);
};

onMounted(() => {
  recordVisitService({
    path: window.location.pathname,
    referer: document.referrer || '',
  }).catch(() => {
    // Visit logging should never block the profile page.
  });
});
</script>

<template>
  <div class="public-profile">
    <div class="public-profile__glow" aria-hidden="true"></div>
    <main class="public-profile__panel">
      <p class="public-profile__brand">{{ appConfig.appName }}</p>
      <h1 class="public-profile__name">{{ profile.name }}</h1>
      <p class="public-profile__title">{{ profile.title }}</p>
      <p class="public-profile__bio">{{ profile.bio }}</p>

      <div class="public-profile__links">
        <template v-for="link in profile.links" :key="link.label">
          <button
            v-if="link.internal"
            type="button"
            class="public-profile__link"
            @click="goInternal(link.href)"
          >
            {{ link.label }}
          </button>
          <a
            v-else
            class="public-profile__link"
            :href="link.href"
            target="_blank"
            rel="noopener noreferrer"
          >
            {{ link.label }}
          </a>
        </template>
      </div>
    </main>
  </div>
</template>

<style scoped>
.public-profile {
  position: relative;
  min-height: 100vh;
  display: grid;
  place-items: center;
  padding: 2rem 1.25rem;
  overflow: hidden;
  background:
    radial-gradient(circle at top left, rgba(34, 120, 110, 0.22), transparent 42%),
    radial-gradient(circle at bottom right, rgba(180, 120, 60, 0.18), transparent 40%),
    linear-gradient(160deg, #f7f3ea 0%, #ebe4d6 48%, #e3ddd0 100%);
  color: #1f2a24;
  font-family: 'Prompt', 'Sukhumvit Set', sans-serif;
}

.public-profile__glow {
  position: absolute;
  inset: 18% 12% auto;
  height: 40%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.55), transparent 70%);
  filter: blur(20px);
  pointer-events: none;
}

.public-profile__panel {
  position: relative;
  width: min(560px, 100%);
  text-align: left;
  animation: rise 700ms ease-out both;
}

.public-profile__brand {
  margin: 0 0 1.5rem;
  font-size: 0.85rem;
  letter-spacing: 0.18em;
  text-transform: uppercase;
  color: #5f6f66;
  animation: fade 900ms ease-out both;
}

.public-profile__name {
  margin: 0;
  font-size: clamp(2.6rem, 8vw, 4.4rem);
  line-height: 0.95;
  font-weight: 700;
  letter-spacing: -0.04em;
}

.public-profile__title {
  margin: 0.85rem 0 0;
  font-size: 1.05rem;
  color: #3f5248;
}

.public-profile__bio {
  margin: 1.25rem 0 0;
  max-width: 34rem;
  font-size: 1rem;
  line-height: 1.7;
  color: #4a5a52;
  animation: fade 1100ms ease-out both;
}

.public-profile__links {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  margin-top: 2rem;
}

.public-profile__link {
  appearance: none;
  border: 1px solid rgba(31, 42, 36, 0.18);
  background: rgba(255, 255, 255, 0.55);
  color: inherit;
  text-decoration: none;
  padding: 0.7rem 1.1rem;
  border-radius: 999px;
  font: inherit;
  cursor: pointer;
  transition:
    transform 180ms ease,
    background 180ms ease,
    border-color 180ms ease;
}

.public-profile__link:hover {
  transform: translateY(-2px);
  background: rgba(255, 255, 255, 0.85);
  border-color: rgba(31, 42, 36, 0.35);
}

@keyframes rise {
  from {
    opacity: 0;
    transform: translateY(18px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fade {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
</style>
