import { defineStore } from 'pinia';
import {
  listPermissionsService,
  listRolesService,
  updateRolePermissionsService,
} from '../services/roles';
import { useAppStore } from './app';
import { i18n } from '../plugins/i18n';

export const useRolesStore = defineStore('roles', {
  state: () => ({
    items: [],
    permissions: [],
  }),
  actions: {
    async fetchRolesAction() {
      const appStore = useAppStore();
      try {
        appStore.setLoading(true);
        const res = await listRolesService();
        this.items = res.items || [];
      } catch (error) {
        appStore.setError(error.message || i18n.global.t('error.unknownError'));
        throw error;
      } finally {
        appStore.setLoading(false);
      }
    },
    async fetchPermissionsAction() {
      const res = await listPermissionsService();
      this.permissions = res.items || [];
      return this.permissions;
    },
    async updateRolePermissionsAction(roleId, permissionIds) {
      const appStore = useAppStore();
      try {
        appStore.setLoading(true);
        await updateRolePermissionsService(roleId, permissionIds);
        appStore.showNotification(i18n.global.t('roles.updateSuccess'), 'success');
        await this.fetchRolesAction();
      } catch (error) {
        appStore.setError(error.message || i18n.global.t('error.unknownError'));
        throw error;
      } finally {
        appStore.setLoading(false);
      }
    },
  },
});
