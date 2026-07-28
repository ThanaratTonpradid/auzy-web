import { ConfigName } from '../constants';
import { useAuthStore } from '../stores/auth';
import { useStaffStore } from '../stores/staff';

export const routerHook = async (to) => {
  const authStore = useAuthStore();
  const staffStore = useStaffStore();

  const token = localStorage.getItem(ConfigName.ACCESS_TOKEN);
  const isAuthenticated = !!token;
  authStore.isLogin = isAuthenticated;

  if (to.meta.requiresAuth) {
    if (!isAuthenticated) {
      return { name: 'login' };
    }

    if (!staffStore.profile?.id) {
      const ok = await authStore.bootstrapAuth();
      if (!ok) {
        return { name: 'login' };
      }
    }

    if (to.meta.permission && !staffStore.hasPermission(to.meta.permission)) {
      return { name: 'dashboard' };
    }

    return true;
  }

  if (isAuthenticated && to.name === 'login') {
    return { name: 'dashboard' };
  }

  return true;
};
