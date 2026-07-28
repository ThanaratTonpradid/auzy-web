<script setup>
import { useField, useForm } from 'vee-validate';
import { useI18n } from 'vue-i18n';
import * as yup from 'yup';

const { t } = useI18n();

const schema = yup.object({
    username: yup.string().required(t('validation.required', { field: t('auth.username') })),
    password: yup.string().required(t('validation.required', { field: t('auth.password') })),
  })

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
  <form @submit.prevent="submit">
    <v-text-field
      v-model="username.value.value"
      :error-messages="username.errorMessage.value"
      :label="$t('auth.username')"
      density="compact"
      variant="outlined"
    ></v-text-field>

    <v-text-field
      v-model="password.value.value"
      :error-messages="password.errorMessage.value"
      :label="$t('auth.password')"
      type="password"
      density="compact"
      variant="outlined"
    ></v-text-field>

    <v-btn class="me-4" color="primary" block type="submit">
      {{ $t('auth.login') }}
    </v-btn>
  </form>
</template>
