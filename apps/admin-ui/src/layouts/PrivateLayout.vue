<script setup>
import AppBar from '../components/AppBar.vue';
</script>

<template>
  <div class="private-layout">
    <AppBar />
    <v-main>
      <v-container class="pa-0" fluid>
        <v-row no-gutters>
          <v-col>
            <RouterView />
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </div>
</template>
