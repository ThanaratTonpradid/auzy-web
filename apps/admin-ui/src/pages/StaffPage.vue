<script setup>
import { computed, onMounted, reactive, ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useDisplay } from 'vuetify';
import { useStaffListStore } from '../stores/staffList';
import { useRolesStore } from '../stores/roles';
import { useStaffStore } from '../stores/staff';
import { Permissions } from '../constants';

const staffListStore = useStaffListStore();
const rolesStore = useRolesStore();
const staffStore = useStaffStore();
const { mdAndUp } = useDisplay();

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
  <div class="page-shell">
    <header class="page-header page-header--row">
      <div>
        <p class="page-header__eyebrow">{{ $t('common.appName') }}</p>
        <h1 class="page-header__title">{{ $t('staff.title') }}</h1>
        <p class="page-header__subtitle">{{ $t('staff.subtitle') }}</p>
      </div>
      <v-btn
        v-if="canCreate"
        color="primary"
        rounded="lg"
        prepend-icon="mdi-plus"
        class="page-header__action"
        @click="openCreate"
      >
        {{ $t('staff.create') }}
      </v-btn>
    </header>

    <section v-if="mdAndUp" class="surface-panel quiet-table overflow-hidden">
      <v-data-table :headers="headers" :items="items" item-value="id" class="bg-transparent">
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
    </section>

    <section v-else class="mobile-list">
      <article
        v-for="item in items"
        :key="item.id"
        class="surface-panel surface-panel--pad mobile-card"
      >
        <div class="mobile-card__head">
          <div>
            <h2 class="mobile-card__title">{{ item.fullname || item.username }}</h2>
            <p class="mobile-card__meta">@{{ item.username }} · #{{ item.id }}</p>
          </div>
          <div class="mobile-card__actions">
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
          </div>
        </div>
        <div class="mobile-card__chips">
          <span class="meta-chip">{{ item.roleLabel || '-' }}</span>
          <v-chip size="small" :color="item.isActive ? 'success' : 'default'" variant="tonal">
            {{ item.isActive ? $t('staff.active') : $t('common.no') }}
          </v-chip>
          <v-chip
            v-if="item.isAdmin"
            size="small"
            color="primary"
            variant="tonal"
          >
            {{ $t('staff.admin') }}
          </v-chip>
        </div>
      </article>
      <p v-if="!items.length" class="mobile-list__empty">{{ $t('common.noData') }}</p>
    </section>

    <v-dialog v-model="dialog" :fullscreen="!mdAndUp" max-width="520">
      <v-card rounded="xl" class="surface-dialog">
        <v-card-title class="text-h6 d-flex align-center justify-space-between">
          <span>{{ editingId ? $t('staff.edit') : $t('staff.create') }}</span>
          <v-btn v-if="!mdAndUp" icon="mdi-close" variant="text" @click="dialog = false" />
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
            density="comfortable"
            rounded="lg"
            class="mb-2"
          />
          <v-text-field
            v-model="form.password"
            :label="$t('auth.password')"
            type="password"
            :hint="editingId ? $t('staff.passwordHint') : ''"
            persistent-hint
            variant="outlined"
            density="comfortable"
            rounded="lg"
            class="mb-2"
          />
          <v-text-field
            v-model="form.fullname"
            :label="$t('profile.fullname')"
            variant="outlined"
            density="comfortable"
            rounded="lg"
            class="mb-2"
          />
          <v-select
            v-model="form.roleId"
            :items="roleOptions"
            :label="$t('menu.roles')"
            variant="outlined"
            density="comfortable"
            rounded="lg"
            class="mb-2"
          />
          <v-switch v-model="form.isActive" :label="$t('staff.active')" color="primary" hide-details />
          <v-switch v-model="form.isAdmin" :label="$t('staff.admin')" color="primary" hide-details />
        </v-card-text>
        <v-card-actions class="pa-4">
          <v-spacer />
          <v-btn variant="text" @click="dialog = false">{{ $t('common.cancel') }}</v-btn>
          <v-btn color="primary" rounded="lg" @click="submitForm">{{ $t('common.save') }}</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="deleteDialog" max-width="400">
      <v-card rounded="xl">
        <v-card-title>{{ $t('staff.deleteConfirmTitle') }}</v-card-title>
        <v-card-text>{{ $t('staff.deleteConfirmMessage') }}</v-card-text>
        <v-card-actions class="pa-4">
          <v-spacer />
          <v-btn variant="text" @click="deleteDialog = false">{{ $t('common.cancel') }}</v-btn>
          <v-btn color="error" rounded="lg" @click="confirmDelete">{{ $t('common.delete') }}</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<style scoped>
.page-header--row {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  gap: 1rem;
}

.mobile-list {
  display: grid;
  gap: 0.75rem;
}

.mobile-card__head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 0.5rem;
}

.mobile-card__title {
  margin: 0;
  font-size: 1.05rem;
  letter-spacing: -0.02em;
}

.mobile-card__meta {
  margin: 0.25rem 0 0;
  font-size: 0.82rem;
  color: var(--ink-soft);
}

.mobile-card__actions {
  display: flex;
  flex-shrink: 0;
}

.mobile-card__chips {
  display: flex;
  flex-wrap: wrap;
  gap: 0.45rem;
  margin-top: 0.9rem;
}

.mobile-list__empty {
  margin: 0;
  padding: 2rem 1rem;
  text-align: center;
  color: var(--ink-soft);
}

@media (max-width: 720px) {
  .page-header--row {
    flex-direction: column;
    align-items: stretch;
  }

  .page-header__action {
    width: 100%;
  }
}
</style>
