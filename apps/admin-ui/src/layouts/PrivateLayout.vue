<script setup>
import AppBar from '../components/AppBar.vue';
</script>

<template>
  <div class="private-layout">
    <AppBar />
    <v-main class="private-layout__main">
      <RouterView />
    </v-main>
  </div>
</template>

<style scoped>
.private-layout__main {
  min-height: 100vh;
}
</style>
