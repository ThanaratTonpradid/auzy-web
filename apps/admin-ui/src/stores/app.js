import { defineStore } from 'pinia';
import { appConfig } from '../config/app.config';

export const useAppStore = defineStore('app', {
  state: () => ({
    appName: appConfig.appName,
    appVersion: appConfig.appVersion,
    navIsOpen: false,
    loading: false,
    error: null,
    notification: {
      show: false,
      message: '',
      type: 'info', // info, success, warning, error
      timeout: 3000,
    },
  }),
  actions: {
    toggleNavAction() {
      this.navIsOpen = !this.navIsOpen;
    },
    setLoading(loading) {
      this.loading = loading;
    },
    setError(error) {
      this.error = error;
      if (error) {
        this.showNotification(error, 'error');
      }
    },
    clearError() {
      this.error = null;
    },
    showNotification(message, type = 'info', timeout = 3000) {
      this.notification = {
        show: true,
        message,
        type,
        timeout,
      };
    },
    hideNotification() {
      this.notification.show = false;
    },
  },
});
