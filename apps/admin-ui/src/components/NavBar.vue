<script setup>
import { storeToRefs } from 'pinia';
import { useAppStore } from '../stores/app';

const appStore = useAppStore();
const { navIsOpen } = storeToRefs(appStore);
</script>

<template>
  <v-navigation-drawer v-model="navIsOpen" temporary>
    <v-list lines="one">
      <v-list-item 
        :title="$t('auth.login')" 
        :to="{ name: 'login' }"
        prepend-icon="mdi-login"
      ></v-list-item>
      <v-list-item 
        :title="$t('menu.profile')" 
        :to="{ name: 'profile' }"
        prepend-icon="mdi-account"
      ></v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>
