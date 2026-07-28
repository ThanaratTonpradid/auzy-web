<script setup>
import { computed, onMounted, ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useRolesStore } from '../stores/roles';
import { useStaffStore } from '../stores/staff';

const rolesStore = useRolesStore();
const staffStore = useStaffStore();
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
  <v-container>
    <div class="mb-4">
      <h1 class="text-h4">{{ $t('roles.title') }}</h1>
      <p class="text-body-2 text-medium-emphasis">{{ $t('roles.subtitle') }}</p>
    </div>

    <v-row>
      <v-col v-for="role in items" :key="role.id" cols="12" md="4">
        <v-sheet border rounded class="pa-4 h-100">
          <div class="d-flex align-center justify-space-between mb-3">
            <div>
              <div class="text-h6">{{ role.label }}</div>
              <div class="text-caption text-medium-emphasis">ID: {{ role.id }}</div>
            </div>
            <v-btn
              v-if="canUpdate"
              size="small"
              variant="tonal"
              prepend-icon="mdi-shield-edit"
              @click="openEdit(role)"
            >
              {{ $t('roles.editPermissions') }}
            </v-btn>
          </div>
          <div class="d-flex flex-wrap ga-2">
            <v-chip
              v-for="code in role.permissions"
              :key="code"
              size="small"
              variant="outlined"
            >
              {{ code }}
            </v-chip>
            <span v-if="!(role.permissions || []).length" class="text-body-2 text-medium-emphasis">
              {{ $t('roles.noPermissions') }}
            </span>
          </div>
        </v-sheet>
      </v-col>
    </v-row>

    <v-dialog v-model="dialog" max-width="560">
      <v-card>
        <v-card-title>
          {{ $t('roles.editPermissions') }} — {{ editingRole?.label }}
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
            density="compact"
          />
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn variant="text" @click="dialog = false">{{ $t('common.cancel') }}</v-btn>
          <v-btn color="primary" @click="savePermissions">{{ $t('common.save') }}</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>
