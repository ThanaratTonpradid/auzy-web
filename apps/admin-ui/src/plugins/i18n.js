import { createI18n } from 'vue-i18n';
import th from '../locales/th.json';
import en from '../locales/en.json';

// Get saved language from localStorage or default to Thai
const savedLanguage = localStorage.getItem('language') || 'th';

// Set initial lang attribute for CSS selector
if (typeof document !== 'undefined') {
  document.documentElement.setAttribute('lang', savedLanguage);
}

export const i18n = createI18n({
  legacy: false, // Use Composition API mode
  locale: savedLanguage,
  fallbackLocale: 'en',
  messages: {
    th,
    en,
  },
  globalInjection: true,
});

