import { defineStore } from 'pinia';
import {
  createStaffService,
  deleteStaffService,
  listStaffService,
  updateStaffService,
} from '../services/staff';
import { useAppStore } from './app';
import { i18n } from '../plugins/i18n';

export const useStaffListStore = defineStore('staffList', {
  state: () => ({
    items: [],
    total: 0,
  }),
  actions: {
    async fetchStaffAction() {
      const appStore = useAppStore();
      try {
        appStore.setLoading(true);
        const res = await listStaffService();
        this.items = res.items || [];
        this.total = res.total || 0;
      } catch (error) {
        appStore.setError(error.message || i18n.global.t('error.unknownError'));
        throw error;
      } finally {
        appStore.setLoading(false);
      }
    },
    async createStaffAction(payload) {
      const appStore = useAppStore();
      try {
        appStore.setLoading(true);
        await createStaffService(payload);
        appStore.showNotification(i18n.global.t('staff.createSuccess'), 'success');
        await this.fetchStaffAction();
      } catch (error) {
        appStore.setError(error.message || i18n.global.t('error.unknownError'));
        throw error;
      } finally {
        appStore.setLoading(false);
      }
    },
    async updateStaffAction(id, payload) {
      const appStore = useAppStore();
      try {
        appStore.setLoading(true);
        await updateStaffService(id, payload);
        appStore.showNotification(i18n.global.t('staff.updateSuccess'), 'success');
        await this.fetchStaffAction();
      } catch (error) {
        appStore.setError(error.message || i18n.global.t('error.unknownError'));
        throw error;
      } finally {
        appStore.setLoading(false);
      }
    },
    async deleteStaffAction(id) {
      const appStore = useAppStore();
      try {
        appStore.setLoading(true);
        await deleteStaffService(id);
        appStore.showNotification(i18n.global.t('staff.deleteSuccess'), 'success');
        await this.fetchStaffAction();
      } catch (error) {
        appStore.setError(error.message || i18n.global.t('error.unknownError'));
        throw error;
      } finally {
        appStore.setLoading(false);
      }
    },
  },
});
