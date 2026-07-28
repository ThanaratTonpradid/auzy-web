import { defineStore } from 'pinia';
import { loginService, logoutService } from '../services/auth';
import { ConfigName } from '../constants';
import router from '../router';
import { useAppStore } from './app';
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
      try {
        appStore.setLoading(true);
        appStore.clearError();
        
        const res = await loginService({ 
          username: payload.username, 
          password: payload.password 
        });
        
        // Save both access token and refresh token
        localStorage.setItem(ConfigName.ACCESS_TOKEN, res.token);
        localStorage.setItem(ConfigName.REFRESH_TOKEN, res.refreshToken);
        
        this.isLogin = true;
        this.user = res.user || null;
        
        // Calculate token expiration time
        if (res.expiresIn) {
          this.tokenExpiresAt = Date.now() + (res.expiresIn * 1000);
        }
        
        appStore.showNotification(i18n.global.t('auth.loginSuccess'), 'success');
        router.replace({ name: 'profile' });
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
      try {
        appStore.setLoading(true);
        await logoutService();
        
        localStorage.removeItem(ConfigName.ACCESS_TOKEN);
        localStorage.removeItem(ConfigName.REFRESH_TOKEN);
        this.isLogin = false;
        this.user = null;
        this.tokenExpiresAt = null;
        
        appStore.showNotification(i18n.global.t('auth.logoutSuccess'), 'success');
        router.replace({ name: 'login' });
      } catch (error) {
        console.error('Logout error:', error);
        // Still logout locally even if server request fails
        localStorage.removeItem(ConfigName.ACCESS_TOKEN);
        localStorage.removeItem(ConfigName.REFRESH_TOKEN);
        this.isLogin = false;
        this.user = null;
        this.tokenExpiresAt = null;
        router.replace({ name: 'login' });
      } finally {
        appStore.setLoading(false);
      }
    },
    checkAuth() {
      const token = localStorage.getItem(ConfigName.ACCESS_TOKEN);
      this.isLogin = !!token;
      return this.isLogin;
    },
  },
});
