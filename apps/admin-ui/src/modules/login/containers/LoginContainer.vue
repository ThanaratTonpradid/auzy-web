<script setup>
import LoginForm from '../components/LoginForm.vue';
import { useAuthStore } from '../../../stores/auth';
import { appConfig } from '../../../config/app.config';
import LanguageSwitcher from '../../../components/LanguageSwitcher.vue';

const authStore = useAuthStore();

const login = async (values) => {
  await authStore.loginAction({
    username: values.username,
    password: values.password,
  });
};
</script>

<template>
  <section class="login-stage">
    <div class="login-stage__orb login-stage__orb--a" aria-hidden="true"></div>
    <div class="login-stage__orb login-stage__orb--b" aria-hidden="true"></div>

    <header class="login-stage__top">
      <RouterLink class="login-stage__brand" :to="{ name: 'public-profile' }">
        {{ appConfig.appName }}
      </RouterLink>
      <LanguageSwitcher class="login-stage__lang" />
    </header>

    <div class="login-stage__panel">
      <p class="login-stage__eyebrow">{{ $t('common.appName') }}</p>
      <h1 class="login-stage__title">{{ $t('auth.login') }}</h1>
      <p class="login-stage__copy">
        {{ $t('auth.loginSubtitle') }}
      </p>
      <LoginForm class="mt-8" @submit="login" />
    </div>
  </section>
</template>

<style scoped>
.login-stage {
  position: relative;
  min-height: 100vh;
  display: grid;
  place-items: center;
  padding: 1.5rem;
  overflow: hidden;
}

.login-stage__orb {
  position: absolute;
  border-radius: 999px;
  filter: blur(8px);
  pointer-events: none;
}

.login-stage__orb--a {
  width: 28rem;
  height: 28rem;
  top: -8rem;
  right: -6rem;
  background: radial-gradient(circle, rgba(31, 122, 102, 0.28), transparent 68%);
  animation: drift 10s ease-in-out infinite alternate;
}

.login-stage__orb--b {
  width: 22rem;
  height: 22rem;
  bottom: -7rem;
  left: -5rem;
  background: radial-gradient(circle, rgba(18, 32, 28, 0.16), transparent 70%);
  animation: drift 12s ease-in-out infinite alternate-reverse;
}

.login-stage__top {
  position: absolute;
  inset: 1.25rem 1.25rem auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
  z-index: 2;
}

.login-stage__brand {
  color: var(--ink);
  text-decoration: none;
  font-weight: 700;
  letter-spacing: -0.02em;
}

.login-stage__lang {
  flex-shrink: 0;
}

.login-stage__panel {
  position: relative;
  z-index: 1;
  width: min(420px, 100%);
  padding: 2rem 1.75rem;
  border-radius: 24px;
  border: 1px solid rgba(18, 32, 28, 0.1);
  background: rgba(255, 255, 255, 0.78);
  backdrop-filter: blur(14px);
  box-shadow: 0 24px 60px rgba(18, 32, 28, 0.1);
  animation: rise 560ms cubic-bezier(0.22, 1, 0.36, 1) both;
}

.login-stage__eyebrow {
  margin: 0;
  font-size: 0.75rem;
  letter-spacing: 0.16em;
  text-transform: uppercase;
  color: var(--ink-soft);
}

.login-stage__title {
  margin: 0.55rem 0 0;
  font-size: clamp(2rem, 5vw, 2.6rem);
  line-height: 1;
  letter-spacing: -0.04em;
}

.login-stage__copy {
  margin: 0.75rem 0 0;
  color: var(--ink-soft);
  line-height: 1.55;
}

@keyframes rise {
  from {
    opacity: 0;
    transform: translateY(16px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes drift {
  from {
    transform: translate3d(0, 0, 0);
  }
  to {
    transform: translate3d(12px, -18px, 0);
  }
}
</style>
