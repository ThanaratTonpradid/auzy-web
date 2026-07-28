<script setup>
import { RouterView, useRoute } from 'vue-router';
import { computed } from 'vue';
import NavBar from './components/NavBar.vue';
import GlobalNotification from './components/GlobalNotification.vue';
import GlobalLoading from './components/GlobalLoading.vue';

const route = useRoute();
const showAdminNav = computed(() => route.path.startsWith('/pri'));
</script>

<template>
  <v-app full-height class="auzy-app">
    <NavBar v-if="showAdminNav" />
    <RouterView />
    <GlobalNotification />
    <GlobalLoading />
  </v-app>
</template>
