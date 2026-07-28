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
  <section class="login-page">
    <div class="login-page__pattern" aria-hidden="true" />

    <div class="login-card">
      <aside class="login-card__banner" aria-label="Brand">
        <div class="login-card__grid" aria-hidden="true">
          <span /><span /><span /><span />
          <span /><span /><span /><span />
        </div>
        <div class="login-card__banner-fade" aria-hidden="true" />
        <div class="login-card__banner-body">
          <RouterLink class="login-card__brand" :to="{ name: 'public-profile' }">
            {{ appConfig.appName }}
          </RouterLink>
          <h1 class="login-card__banner-title">{{ $t('auth.bannerTitle') }}</h1>
          <p class="login-card__banner-copy">{{ $t('auth.bannerCopy') }}</p>
        </div>
      </aside>

      <div class="login-card__form">
        <div class="login-card__form-top">
          <LanguageSwitcher />
        </div>

        <div class="login-card__form-body">
          <h2 class="login-card__title">{{ $t('auth.login') }}</h2>
          <p class="login-card__subtitle">{{ $t('auth.loginSubtitle') }}</p>
          <LoginForm class="login-card__fields" @submit="login" />
        </div>

        <p class="login-card__build">
          {{ appConfig.appName }} · v{{ appConfig.appVersion }}
        </p>
      </div>
    </div>
  </section>
</template>

<style scoped>
.login-page {
  position: relative;
  min-height: 100vh;
  display: grid;
  place-items: center;
  padding: 1.5rem;
  overflow: hidden;
  background: #eef2f0;
}

.login-page__pattern {
  position: absolute;
  inset: 0;
  opacity: 0.45;
  background-image:
    radial-gradient(circle at 1px 1px, rgba(18, 32, 28, 0.08) 1px, transparent 0);
  background-size: 22px 22px;
  mask-image: radial-gradient(circle at center, black 35%, transparent 85%);
  pointer-events: none;
}

.login-card {
  position: relative;
  z-index: 1;
  display: grid;
  grid-template-columns: 1fr 1fr;
  width: min(920px, 100%);
  min-height: min(560px, calc(100vh - 3rem));
  border-radius: 20px;
  overflow: hidden;
  background: #fff;
  box-shadow:
    0 24px 60px rgba(18, 32, 28, 0.12),
    0 2px 0 rgba(255, 255, 255, 0.6) inset;
  animation: rise 560ms var(--motion) both;
}

.login-card__banner {
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  padding: 2rem;
  color: #f4faf7;
  background: #145547;
  overflow: hidden;
}

.login-card__grid {
  position: absolute;
  inset: 0;
  display: grid;
  grid-template-columns: 1.1fr 0.9fr;
  grid-template-rows: 1.1fr 0.7fr 1fr;
  gap: 0;
}

.login-card__grid span:nth-child(1) {
  background: #1f7a66;
}
.login-card__grid span:nth-child(2) {
  background: #0f3f35;
}
.login-card__grid span:nth-child(3) {
  background: #176955;
  grid-column: 1;
}
.login-card__grid span:nth-child(4) {
  background: #2a9b82;
}
.login-card__grid span:nth-child(5) {
  background: #0c332b;
  grid-column: 1 / -1;
}
.login-card__grid span:nth-child(6),
.login-card__grid span:nth-child(7),
.login-card__grid span:nth-child(8) {
  display: none;
}

.login-card__banner-fade {
  position: absolute;
  inset: 0;
  background:
    linear-gradient(180deg, rgba(8, 28, 24, 0.15) 0%, rgba(8, 28, 24, 0.55) 100%),
    radial-gradient(circle at 80% 18%, rgba(255, 255, 255, 0.16), transparent 36%);
  pointer-events: none;
}

.login-card__banner-body {
  position: relative;
  z-index: 1;
}

.login-card__brand {
  display: inline-block;
  margin-bottom: 1.75rem;
  color: inherit;
  text-decoration: none;
  font-size: clamp(1.8rem, 3vw, 2.35rem);
  font-weight: 700;
  letter-spacing: -0.04em;
  line-height: 1;
}

.login-card__banner-title {
  margin: 0;
  font-size: clamp(1.35rem, 2.4vw, 1.75rem);
  font-weight: 650;
  letter-spacing: -0.03em;
  line-height: 1.2;
  max-width: 12ch;
}

.login-card__banner-copy {
  margin: 0.75rem 0 0;
  max-width: 28ch;
  color: rgba(244, 250, 247, 0.82);
  line-height: 1.6;
  font-size: 0.95rem;
}

.login-card__form {
  display: grid;
  grid-template-rows: auto 1fr auto;
  background: #fff;
  min-width: 0;
}

.login-card__form-top {
  display: flex;
  justify-content: flex-end;
  padding: 1.1rem 1.25rem 0;
}

.login-card__form-body {
  width: min(340px, calc(100% - 3rem));
  margin: 0 auto;
  padding: 1.25rem 0 1.5rem;
  align-self: center;
}

.login-card__title {
  margin: 0;
  font-size: clamp(1.7rem, 3vw, 2rem);
  font-weight: 700;
  letter-spacing: -0.03em;
  line-height: 1.15;
  color: var(--ink);
}

.login-card__subtitle {
  margin: 0.55rem 0 0;
  color: var(--ink-soft);
  line-height: 1.55;
  font-size: 0.95rem;
}

.login-card__fields {
  margin-top: 1.5rem;
}

.login-card__build {
  margin: 0;
  padding: 0 1.25rem 1rem;
  text-align: center;
  font-size: 0.72rem;
  color: rgba(61, 82, 74, 0.55);
}

@keyframes rise {
  from {
    opacity: 0;
    transform: translateY(16px) scale(0.985);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

@media (max-width: 860px) {
  .login-card {
    grid-template-columns: 1fr;
    min-height: auto;
  }

  .login-card__banner {
    min-height: 200px;
    padding: 1.5rem;
  }

  .login-card__brand {
    margin-bottom: 1rem;
    font-size: 1.7rem;
  }

  .login-card__banner-title {
    max-width: none;
    font-size: 1.25rem;
  }

  .login-card__form-body {
    padding-top: 0.5rem;
    padding-bottom: 1rem;
  }
}
</style>
