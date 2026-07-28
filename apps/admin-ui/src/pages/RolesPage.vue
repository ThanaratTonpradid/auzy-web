<script setup>
import { computed, onMounted, ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useDisplay } from 'vuetify';
import { useRolesStore } from '../stores/roles';
import { useStaffStore } from '../stores/staff';

const rolesStore = useRolesStore();
const staffStore = useStaffStore();
const { mdAndUp } = useDisplay();
const { items, permissions } = storeToRefs(rolesStore);

const dialog = ref(false);
const editingRole = ref(null);
const selectedPermissionIds = ref([]);
const formError = ref('');

const canUpdate = computed(() => staffStore.canUpdateRoles);

const openEdit = async (role) => {
  formError.value = '';
  editingRole.value = role;
  if (!permissions.value.length) {
    await rolesStore.fetchPermissionsAction();
  }
  const selectedCodes = new Set(role.permissions || []);
  selectedPermissionIds.value = permissions.value
    .filter((p) => selectedCodes.has(p.codeName))
    .map((p) => p.id);
  dialog.value = true;
};

const savePermissions = async () => {
  try {
    await rolesStore.updateRolePermissionsAction(
      editingRole.value.id,
      selectedPermissionIds.value
    );
    dialog.value = false;
  } catch (error) {
    formError.value = error.message || 'Failed';
  }
};

onMounted(async () => {
  await rolesStore.fetchRolesAction();
  if (canUpdate.value) {
    await rolesStore.fetchPermissionsAction();
  }
});
</script>

<template>
  <div class="page-shell">
    <header class="page-header">
      <p class="page-header__eyebrow">{{ $t('common.appName') }}</p>
      <h1 class="page-header__title">{{ $t('roles.title') }}</h1>
      <p class="page-header__subtitle">{{ $t('roles.subtitle') }}</p>
    </header>

    <div class="roles-grid">
      <section v-for="role in items" :key="role.id" class="surface-panel surface-panel--pad role-card">
        <div class="role-card__head">
          <div>
            <h2 class="role-card__title">{{ role.label }}</h2>
            <p class="role-card__id">ID {{ role.id }}</p>
          </div>
          <v-btn
            v-if="canUpdate"
            size="small"
            variant="tonal"
            color="primary"
            rounded="lg"
            class="role-card__edit"
            prepend-icon="mdi-shield-edit-outline"
            @click="openEdit(role)"
          >
            {{ $t('roles.editPermissions') }}
          </v-btn>
        </div>
        <div class="role-card__perms">
          <span
            v-for="code in role.permissions"
            :key="code"
            class="meta-chip"
          >
            {{ code }}
          </span>
          <span
            v-if="!(role.permissions || []).length"
            class="role-card__empty"
          >
            {{ $t('roles.noPermissions') }}
          </span>
        </div>
      </section>
    </div>

    <v-dialog v-model="dialog" :fullscreen="!mdAndUp" max-width="560">
      <v-card rounded="xl">
        <v-card-title class="text-h6 d-flex align-center justify-space-between">
          <span>{{ $t('roles.editPermissions') }} — {{ editingRole?.label }}</span>
          <v-btn v-if="!mdAndUp" icon="mdi-close" variant="text" @click="dialog = false" />
        </v-card-title>
        <v-card-text>
          <v-alert v-if="formError" type="error" variant="tonal" class="mb-4">
            {{ formError }}
          </v-alert>
          <v-select
            v-model="selectedPermissionIds"
            :items="permissions"
            item-title="codeName"
            item-value="id"
            :label="$t('menu.permissions')"
            multiple
            chips
            closable-chips
            variant="outlined"
            density="comfortable"
            rounded="lg"
          />
        </v-card-text>
        <v-card-actions class="pa-4">
          <v-spacer />
          <v-btn variant="text" @click="dialog = false">{{ $t('common.cancel') }}</v-btn>
          <v-btn color="primary" rounded="lg" @click="savePermissions">
            {{ $t('common.save') }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<style scoped>
.roles-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 1rem;
}

.role-card {
  display: grid;
  gap: 1rem;
  min-height: 100%;
  transition:
    transform var(--motion),
    border-color var(--motion);
}

.role-card:hover {
  transform: translateY(-2px);
  border-color: rgba(31, 122, 102, 0.28);
}

.role-card__head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 0.75rem;
}

.role-card__title {
  margin: 0;
  font-size: 1.15rem;
  letter-spacing: -0.02em;
}

.role-card__id {
  margin: 0.25rem 0 0;
  font-size: 0.8rem;
  color: var(--ink-soft);
}

.role-card__perms {
  display: flex;
  flex-wrap: wrap;
  gap: 0.45rem;
}

.role-card__empty {
  color: var(--ink-soft);
  font-size: 0.9rem;
}

@media (max-width: 1100px) {
  .roles-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 720px) {
  .roles-grid {
    grid-template-columns: 1fr;
  }

  .role-card__head {
    flex-direction: column;
  }

  .role-card__edit {
    width: 100%;
  }
}
</style>
