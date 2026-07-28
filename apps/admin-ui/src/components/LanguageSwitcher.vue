<template>
  <v-menu offset-y>
    <template v-slot:activator="{ props }">
      <v-btn
        v-bind="props"
        icon
        variant="text"
        :title="$t('language.changeLanguage')"
      >
        <v-avatar size="24">
          <v-img :src="currentFlag" :alt="$t('language.changeLanguage')"></v-img>
        </v-avatar>
      </v-btn>
    </template>

    <v-list>
      <v-list-item
        v-for="lang in languages"
        :key="lang.code"
        :active="currentLanguage === lang.code"
        @click="changeLanguage(lang.code)"
      >
        <template v-slot:prepend>
          <v-avatar size="24">
            <v-img :src="lang.flag" :alt="lang.name"></v-img>
          </v-avatar>
        </template>
        <v-list-item-title>{{ lang.name }}</v-list-item-title>
      </v-list-item>
    </v-list>
  </v-menu>
</template>

<script setup>
import { computed } from 'vue';
import { useI18n } from 'vue-i18n';
import thFlag from '../assets/th_TH.png';
import enFlag from '../assets/en_UK.png';

const { locale } = useI18n();

const languages = [
  { code: 'th', name: 'ไทย', flag: thFlag },
  { code: 'en', name: 'English', flag: enFlag },
];

const currentLanguage = computed(() => locale.value);

const currentFlag = computed(() => {
  const lang = languages.find(l => l.code === currentLanguage.value);
  return lang ? lang.flag : thFlag;
});

const changeLanguage = (lang) => {
  locale.value = lang;
  localStorage.setItem('language', lang);
  // Set lang attribute to trigger CSS selector
  document.documentElement.setAttribute('lang', lang);
};
</script>

