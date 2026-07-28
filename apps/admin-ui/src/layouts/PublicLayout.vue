<script setup></script>

<template>
  <div class="public-layout">
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
