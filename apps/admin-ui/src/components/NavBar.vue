<script setup>
import { storeToRefs } from 'pinia';
import { useAppStore } from '../stores/app';
import { useStaffStore } from '../stores/staff';

const appStore = useAppStore();
const staffStore = useStaffStore();
const { navIsOpen } = storeToRefs(appStore);
</script>

<template>
  <v-navigation-drawer v-model="navIsOpen" temporary>
    <v-list lines="one" nav>
      <v-list-item
        :title="$t('menu.dashboard')"
        :to="{ name: 'dashboard' }"
        prepend-icon="mdi-view-dashboard"
      />
      <v-list-item
        :title="$t('menu.profile')"
        :to="{ name: 'profile' }"
        prepend-icon="mdi-account"
      />
      <v-list-item
        v-if="staffStore.canReadStaff"
        :title="$t('menu.staff')"
        :to="{ name: 'staff' }"
        prepend-icon="mdi-account-group"
      />
      <v-list-item
        v-if="staffStore.canReadRoles"
        :title="$t('menu.roles')"
        :to="{ name: 'roles' }"
        prepend-icon="mdi-shield-account"
      />
    </v-list>
  </v-navigation-drawer>
</template>
