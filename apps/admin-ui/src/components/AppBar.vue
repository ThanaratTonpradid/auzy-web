<script setup>
import { computed } from 'vue';
import { useAppStore } from '../stores/app';
import { useAuthStore } from '../stores/auth';
import LanguageSwitcher from './LanguageSwitcher.vue';

const appStore = useAppStore();
const authStore = useAuthStore();

const username = computed(() => authStore.user?.username || 'User');

const toggleNavbar = () => {
  appStore.toggleNavAction();
};

const logout = async () => {
  await authStore.logoutAction();
};
</script>

<template>
  <v-app-bar>
    <v-app-bar-nav-icon @click="toggleNavbar"></v-app-bar-nav-icon>
    
    <v-toolbar-title>{{ $t('common.appName') }}</v-toolbar-title>

    <v-spacer></v-spacer>

    <!-- Language Switcher -->
    <LanguageSwitcher />

    <!-- User Menu -->
    <v-menu>
      <template v-slot:activator="{ props }">
        <v-btn 
          v-bind="props" 
          variant="text"
          prepend-icon="mdi-account-circle"
        >
          {{ username }}
        </v-btn>
      </template>
      <v-list>
        <v-list-item :to="{ name: 'profile' }">
          <template v-slot:prepend>
            <v-icon>mdi-account</v-icon>
          </template>
          <v-list-item-title>{{ $t('menu.profile') }}</v-list-item-title>
        </v-list-item>
        
        <v-divider></v-divider>
        
        <v-list-item @click="logout">
          <template v-slot:prepend>
            <v-icon color="error">mdi-logout</v-icon>
          </template>
          <v-list-item-title class="text-error">
            {{ $t('auth.logout') }}
          </v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </v-app-bar>
</template>
