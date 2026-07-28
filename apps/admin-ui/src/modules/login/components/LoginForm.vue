<script setup>
import { useField, useForm } from 'vee-validate';
import { useI18n } from 'vue-i18n';
import * as yup from 'yup';

const { t } = useI18n();

const schema = yup.object({
  username: yup.string().required(t('validation.required', { field: t('auth.username') })),
  password: yup.string().required(t('validation.required', { field: t('auth.password') })),
});

const { handleSubmit } = useForm({
  validationSchema: schema,
});
const username = useField('username');
const password = useField('password');

const emit = defineEmits(['submit']);
const submit = handleSubmit((values) => {
  emit('submit', values);
});
</script>

<template>
  <form class="login-form" @submit.prevent="submit">
    <v-text-field
      v-model="username.value.value"
      :error-messages="username.errorMessage.value"
      :label="$t('auth.username')"
      autocomplete="username"
      prepend-inner-icon="mdi-account-outline"
    />

    <v-text-field
      v-model="password.value.value"
      :error-messages="password.errorMessage.value"
      :label="$t('auth.password')"
      type="password"
      autocomplete="current-password"
      prepend-inner-icon="mdi-lock-outline"
      class="mt-3"
    />

    <v-btn class="mt-5" color="primary" size="large" block type="submit">
      {{ $t('auth.login') }}
    </v-btn>
  </form>
</template>
