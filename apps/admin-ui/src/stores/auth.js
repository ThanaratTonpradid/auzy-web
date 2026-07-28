import { defineStore } from 'pinia';
import { loginService, logoutService } from '../services/auth';
import { ConfigName } from '../constants';
import router from '../router';
import { useAppStore } from './app';
import { useStaffStore } from './staff';
import { i18n } from '../plugins/i18n';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    isLogin: !!localStorage.getItem(ConfigName.ACCESS_TOKEN),
    user: null,
    tokenExpiresAt: null,
  }),
  actions: {
    async loginAction(payload) {
      const appStore = useAppStore();
      const staffStore = useStaffStore();
      try {
        appStore.setLoading(true);
        appStore.clearError();

        const res = await loginService({
          username: payload.username,
          password: payload.password,
        });

        localStorage.setItem(ConfigName.ACCESS_TOKEN, res.token);
        localStorage.setItem(ConfigName.REFRESH_TOKEN, res.refreshToken);

        this.isLogin = true;

        if (res.expiresIn) {
          this.tokenExpiresAt = Date.now() + res.expiresIn * 1000;
        }

        const profile = await staffStore.getProfileAction();
        this.user = profile;

        appStore.showNotification(i18n.global.t('auth.loginSuccess'), 'success');
        router.replace({ name: 'dashboard' });
      } catch (error) {
        console.error('Login error:', error);
        appStore.setError(error.message || i18n.global.t('auth.loginFailed'));
        throw error;
      } finally {
        appStore.setLoading(false);
      }
    },
    async logoutAction() {
      const appStore = useAppStore();
      const staffStore = useStaffStore();
      try {
        appStore.setLoading(true);
        await logoutService();
      } catch (error) {
        console.error('Logout error:', error);
      } finally {
        localStorage.removeItem(ConfigName.ACCESS_TOKEN);
        localStorage.removeItem(ConfigName.REFRESH_TOKEN);
        this.isLogin = false;
        this.user = null;
        this.tokenExpiresAt = null;
        staffStore.clearProfile();
        appStore.setLoading(false);
        appStore.showNotification(i18n.global.t('auth.logoutSuccess'), 'success');
        router.replace({ name: 'login' });
      }
    },
    async bootstrapAuth() {
      if (!this.isLogin) return false;
      const staffStore = useStaffStore();
      try {
        const profile = await staffStore.getProfileAction();
        this.user = profile;
        return true;
      } catch (error) {
        localStorage.removeItem(ConfigName.ACCESS_TOKEN);
        localStorage.removeItem(ConfigName.REFRESH_TOKEN);
        this.isLogin = false;
        this.user = null;
        staffStore.clearProfile();
        return false;
      }
    },
    checkAuth() {
      const token = localStorage.getItem(ConfigName.ACCESS_TOKEN);
      this.isLogin = !!token;
      return this.isLogin;
    },
  },
});
