import { ConfigName } from '../constants';
import { useAuthStore } from '../stores/auth';
import { useStaffStore } from '../stores/staff';

export const routerHook = async (to, from, next) => {
  const authStore = useAuthStore();
  const staffStore = useStaffStore();

  const token = localStorage.getItem(ConfigName.ACCESS_TOKEN);
  const isAuthenticated = !!token;
  authStore.isLogin = isAuthenticated;

  if (to.meta.requiresAuth) {
    if (!isAuthenticated) {
      next({ name: 'login' });
      return;
    }

    if (!staffStore.profile?.id) {
      const ok = await authStore.bootstrapAuth();
      if (!ok) {
        next({ name: 'login' });
        return;
      }
    }

    if (to.meta.permission && !staffStore.hasPermission(to.meta.permission)) {
      next({ name: 'dashboard' });
      return;
    }

    next();
    return;
  }

  if (isAuthenticated && to.name === 'login') {
    next({ name: 'dashboard' });
    return;
  }

  next();
};
