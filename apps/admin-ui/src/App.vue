<script setup>
import { RouterView } from 'vue-router';
import NavBar from './components/NavBar.vue';
import GlobalNotification from './components/GlobalNotification.vue';
import GlobalLoading from './components/GlobalLoading.vue';
</script>

<template>
  <v-app full-height>
    <NavBar />
    <RouterView />
    <GlobalNotification />
    <GlobalLoading />
  </v-app>
</template>
