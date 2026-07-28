<script setup>
import { computed } from 'vue';
import { useI18n } from 'vue-i18n';

defineProps({
  block: {
    type: Boolean,
    default: false,
  },
});

const { locale, t } = useI18n();

const languages = [
  { code: 'th', label: 'TH', nameKey: 'language.thai' },
  { code: 'en', label: 'EN', nameKey: 'language.english' },
];

const currentLanguage = computed(() => locale.value);

const changeLanguage = (lang) => {
  if (locale.value === lang) return;
  locale.value = lang;
  localStorage.setItem('language', lang);
  document.documentElement.setAttribute('lang', lang);
};
</script>

<template>
  <div
    class="lang-switch"
    :class="{ 'lang-switch--block': block }"
    role="group"
    :aria-label="t('language.changeLanguage')"
  >
    <button
      v-for="lang in languages"
      :key="lang.code"
      type="button"
      class="lang-switch__btn"
      :class="{ 'lang-switch__btn--active': currentLanguage === lang.code }"
      :aria-pressed="currentLanguage === lang.code"
      :title="t(lang.nameKey)"
      @click="changeLanguage(lang.code)"
    >
      {{ lang.label }}
    </button>
  </div>
</template>

<style scoped>
.lang-switch {
  display: inline-grid;
  grid-auto-flow: column;
  grid-auto-columns: 1fr;
  gap: 2px;
  padding: 3px;
  border-radius: 12px;
  border: 1px solid var(--line);
  background: rgba(255, 255, 255, 0.72);
}

.lang-switch--block {
  display: grid;
  width: 100%;
}

.lang-switch__btn {
  appearance: none;
  border: 0;
  margin: 0;
  min-width: 2.5rem;
  height: 1.85rem;
  padding: 0 0.7rem;
  border-radius: 9px;
  background: transparent;
  color: var(--ink-soft);
  font: inherit;
  font-size: 0.75rem;
  font-weight: 650;
  letter-spacing: 0.04em;
  cursor: pointer;
  transition:
    background var(--motion),
    color var(--motion),
    box-shadow var(--motion);
}

.lang-switch__btn:hover {
  color: var(--ink);
}

.lang-switch__btn--active {
  background: var(--accent);
  color: #fff;
  box-shadow: 0 6px 14px rgba(31, 122, 102, 0.22);
}

.lang-switch__btn--active:hover {
  color: #fff;
}
</style>
