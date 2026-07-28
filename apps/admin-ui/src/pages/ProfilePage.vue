<script setup>
import { computed } from 'vue';
import { useAuthStore } from '../stores/auth';

const authStore = useAuthStore();
const user = computed(() => authStore.user);
</script>

<template>
  <main>
    <v-container>
      <v-card>
        <v-card-title class="bg-primary">
          <h2 class="text-h4">{{ $t('profile.title') }}</h2>
        </v-card-title>
        
        <v-card-text class="pt-6">
          <v-row>
            <v-col cols="12">
              <h3 class="text-h6 mb-4">{{ $t('profile.personalInfo') }}</h3>
            </v-col>
          </v-row>
          
          <v-row v-if="user">
            <v-col cols="12" md="6">
              <v-text-field
                :model-value="user.username"
                :label="$t('auth.username')"
                readonly
                variant="outlined"
                density="compact"
              ></v-text-field>
            </v-col>
            
            <v-col cols="12" md="6">
              <v-text-field
                :model-value="user.fullname || '-'"
                :label="$t('profile.fullname')"
                readonly
                variant="outlined"
                density="compact"
              ></v-text-field>
            </v-col>
            
            <v-col cols="12" md="6">
              <v-text-field
                :model-value="user.email || '-'"
                :label="$t('profile.email')"
                readonly
                variant="outlined"
                density="compact"
              ></v-text-field>
            </v-col>
            
            <v-col cols="12" md="6">
              <v-text-field
                :model-value="user.phone || '-'"
                :label="$t('profile.phone')"
                readonly
                variant="outlined"
                density="compact"
              ></v-text-field>
            </v-col>
          </v-row>
        </v-card-text>
        
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" variant="outlined">
            {{ $t('profile.editProfile') }}
          </v-btn>
          <v-btn color="secondary" variant="outlined">
            {{ $t('profile.changePassword') }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-container>
  </main>
</template>
