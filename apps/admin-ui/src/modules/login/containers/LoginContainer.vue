<script>
import LoginForm from '../components/LoginForm.vue';
import { useAuthStore } from '../../../stores/auth';

export default {
  components: { LoginForm },
  setup() {
    const authStore = useAuthStore();
    return { authStore };
  },
  methods: {
    async login(values) {
      await this.authStore.loginAction({ username: values.username, password: values.password });
    }
  },
};
</script>

<template>
  <v-card min-width="320px" max-width="400px">
    <v-card-title class="text-center py-4">
      <v-avatar color="surface-variant" size="120"></v-avatar>
    </v-card-title>
    <v-card-text>
      <LoginForm @submit="login" />
    </v-card-text>
  </v-card>
</template>
