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
      name: 'login',
    },
  ],
};

const goInternal = (name) => {
  router.push({ name });
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
    <div class="public-profile__orb public-profile__orb--a" aria-hidden="true" />
    <div class="public-profile__orb public-profile__orb--b" aria-hidden="true" />
    <main class="public-profile__panel">
      <p class="public-profile__brand">{{ appConfig.appName }}</p>
      <h1 class="public-profile__name">{{ profile.name }}</h1>
      <p class="public-profile__title">{{ profile.title }}</p>
      <p class="public-profile__bio">{{ profile.bio }}</p>

      <div class="public-profile__links">
        <template v-for="link in profile.links" :key="link.label">
          <button
            v-if="link.name"
            type="button"
            class="public-profile__link"
            @click="goInternal(link.name)"
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
  color: var(--ink);
  background:
    radial-gradient(circle at 12% 18%, rgba(31, 122, 102, 0.18), transparent 38%),
    radial-gradient(circle at 88% 12%, rgba(18, 32, 28, 0.08), transparent 32%),
    linear-gradient(165deg, #f4f7f5 0%, #e9efec 55%, #e4ebe7 100%);
}

.public-profile__orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(40px);
  pointer-events: none;
  animation: drift 12s ease-in-out infinite alternate;
}

.public-profile__orb--a {
  width: 22rem;
  height: 22rem;
  top: -4rem;
  left: -3rem;
  background: rgba(31, 122, 102, 0.22);
}

.public-profile__orb--b {
  width: 18rem;
  height: 18rem;
  right: -2rem;
  bottom: 10%;
  background: rgba(216, 239, 232, 0.9);
  animation-delay: -4s;
}

.public-profile__panel {
  position: relative;
  width: min(560px, 100%);
  text-align: left;
  animation: rise 700ms var(--motion) both;
}

.public-profile__brand {
  margin: 0 0 1.5rem;
  font-size: 0.78rem;
  letter-spacing: 0.18em;
  text-transform: uppercase;
  color: var(--ink-soft);
  animation: fade 900ms var(--motion) both;
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
  color: var(--ink-soft);
}

.public-profile__bio {
  margin: 1.25rem 0 0;
  max-width: 34rem;
  font-size: 1rem;
  line-height: 1.7;
  color: var(--ink-soft);
  animation: fade 1100ms var(--motion) both;
}

.public-profile__links {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  margin-top: 2rem;
}

.public-profile__link {
  appearance: none;
  border: 1px solid transparent;
  background: var(--accent);
  color: white;
  text-decoration: none;
  padding: 0.75rem 1.2rem;
  border-radius: 14px;
  font: inherit;
  font-weight: 600;
  cursor: pointer;
  transition:
    transform var(--motion),
    background var(--motion),
    box-shadow var(--motion);
}

.public-profile__link:hover {
  transform: translateY(-2px);
  background: #176955;
  box-shadow: 0 12px 28px rgba(31, 122, 102, 0.28);
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

@keyframes drift {
  from {
    transform: translate3d(0, 0, 0) scale(1);
  }
  to {
    transform: translate3d(18px, 24px, 0) scale(1.06);
  }
}
</style>
