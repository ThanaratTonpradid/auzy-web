<script setup>
import { computed, onMounted, reactive, ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useStaffListStore } from '../stores/staffList';
import { useRolesStore } from '../stores/roles';
import { useStaffStore } from '../stores/staff';
import { Permissions } from '../constants';

const staffListStore = useStaffListStore();
const rolesStore = useRolesStore();
const staffStore = useStaffStore();

const { items } = storeToRefs(staffListStore);
const dialog = ref(false);
const deleteDialog = ref(false);
const editingId = ref(null);
const deletingId = ref(null);
const formError = ref('');

const form = reactive({
  username: '',
  password: '',
  fullname: '',
  roleId: null,
  isActive: true,
  isAdmin: false,
});

const roleOptions = computed(() =>
  (rolesStore.items || []).map((role) => ({
    title: role.label,
    value: role.id,
  }))
);

const headers = computed(() => [
  { title: 'ID', key: 'id', width: 80 },
  { title: 'Username', key: 'username' },
  { title: 'Fullname', key: 'fullname' },
  { title: 'Role', key: 'roleLabel' },
  { title: 'Active', key: 'isActive' },
  { title: 'Admin', key: 'isAdmin' },
  { title: 'Actions', key: 'actions', sortable: false, width: 160 },
]);

const canCreate = computed(() => staffStore.hasPermission(Permissions.STAFFS_CREATE));
const canUpdate = computed(() => staffStore.hasPermission(Permissions.STAFFS_UPDATE));
const canDelete = computed(() => staffStore.hasPermission(Permissions.STAFFS_DELETE));

const resetForm = () => {
  editingId.value = null;
  formError.value = '';
  form.username = '';
  form.password = '';
  form.fullname = '';
  form.roleId = roleOptions.value[0]?.value || null;
  form.isActive = true;
  form.isAdmin = false;
};

const openCreate = () => {
  resetForm();
  dialog.value = true;
};

const openEdit = (item) => {
  editingId.value = item.id;
  formError.value = '';
  form.username = item.username;
  form.password = '';
  form.fullname = item.fullname || '';
  form.roleId = item.roleId;
  form.isActive = item.isActive;
  form.isAdmin = item.isAdmin;
  dialog.value = true;
};

const openDelete = (item) => {
  deletingId.value = item.id;
  deleteDialog.value = true;
};

const submitForm = async () => {
  formError.value = '';
  try {
    if (editingId.value) {
      const payload = {
        fullname: form.fullname,
        roleId: form.roleId,
        isActive: form.isActive,
        isAdmin: form.isAdmin,
      };
      if (form.password) {
        payload.password = form.password;
      }
      await staffListStore.updateStaffAction(editingId.value, payload);
    } else {
      await staffListStore.createStaffAction({
        username: form.username,
        password: form.password,
        fullname: form.fullname,
        roleId: form.roleId,
        isActive: form.isActive,
        isAdmin: form.isAdmin,
      });
    }
    dialog.value = false;
  } catch (error) {
    formError.value = error.message || 'Failed';
  }
};

const confirmDelete = async () => {
  try {
    await staffListStore.deleteStaffAction(deletingId.value);
    deleteDialog.value = false;
  } catch (error) {
    formError.value = error.message || 'Failed';
  }
};

onMounted(async () => {
  await Promise.all([staffListStore.fetchStaffAction(), rolesStore.fetchRolesAction()]);
  if (!form.roleId && roleOptions.value.length) {
    form.roleId = roleOptions.value[0].value;
  }
});
</script>

<template>
  <v-container>
    <div class="d-flex align-center justify-space-between mb-4">
      <div>
        <h1 class="text-h4">{{ $t('staff.title') }}</h1>
        <p class="text-body-2 text-medium-emphasis">{{ $t('staff.subtitle') }}</p>
      </div>
      <v-btn v-if="canCreate" color="primary" prepend-icon="mdi-plus" @click="openCreate">
        {{ $t('staff.create') }}
      </v-btn>
    </div>

    <v-data-table :headers="headers" :items="items" item-value="id" class="elevation-0" border>
      <template #[`item.isActive`]="{ item }">
        <v-chip size="small" :color="item.isActive ? 'success' : 'default'" variant="tonal">
          {{ item.isActive ? $t('common.yes') : $t('common.no') }}
        </v-chip>
      </template>
      <template #[`item.isAdmin`]="{ item }">
        <v-chip size="small" :color="item.isAdmin ? 'primary' : 'default'" variant="tonal">
          {{ item.isAdmin ? $t('common.yes') : $t('common.no') }}
        </v-chip>
      </template>
      <template #[`item.actions`]="{ item }">
        <v-btn
          v-if="canUpdate"
          icon="mdi-pencil"
          variant="text"
          size="small"
          @click="openEdit(item)"
        />
        <v-btn
          v-if="canDelete"
          icon="mdi-delete"
          variant="text"
          size="small"
          color="error"
          @click="openDelete(item)"
        />
      </template>
    </v-data-table>

    <v-dialog v-model="dialog" max-width="520">
      <v-card>
        <v-card-title>
          {{ editingId ? $t('staff.edit') : $t('staff.create') }}
        </v-card-title>
        <v-card-text>
          <v-alert v-if="formError" type="error" variant="tonal" class="mb-4">
            {{ formError }}
          </v-alert>
          <v-text-field
            v-model="form.username"
            :label="$t('auth.username')"
            :disabled="!!editingId"
            variant="outlined"
            density="compact"
            class="mb-2"
          />
          <v-text-field
            v-model="form.password"
            :label="$t('auth.password')"
            type="password"
            :hint="editingId ? $t('staff.passwordHint') : ''"
            persistent-hint
            variant="outlined"
            density="compact"
            class="mb-2"
          />
          <v-text-field
            v-model="form.fullname"
            :label="$t('profile.fullname')"
            variant="outlined"
            density="compact"
            class="mb-2"
          />
          <v-select
            v-model="form.roleId"
            :items="roleOptions"
            :label="$t('menu.roles')"
            variant="outlined"
            density="compact"
            class="mb-2"
          />
          <v-switch v-model="form.isActive" :label="$t('staff.active')" color="primary" hide-details />
          <v-switch v-model="form.isAdmin" :label="$t('staff.admin')" color="primary" hide-details />
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn variant="text" @click="dialog = false">{{ $t('common.cancel') }}</v-btn>
          <v-btn color="primary" @click="submitForm">{{ $t('common.save') }}</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="deleteDialog" max-width="400">
      <v-card>
        <v-card-title>{{ $t('staff.deleteConfirmTitle') }}</v-card-title>
        <v-card-text>{{ $t('staff.deleteConfirmMessage') }}</v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn variant="text" @click="deleteDialog = false">{{ $t('common.cancel') }}</v-btn>
          <v-btn color="error" @click="confirmDelete">{{ $t('common.delete') }}</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>
