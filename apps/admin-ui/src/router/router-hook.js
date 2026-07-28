import { ConfigName } from '../constants';
import { useAuthStore } from '../stores/auth';

export const routerHook = async (to, from, next) => {
  const authStore = useAuthStore();
  
  // Check if user has token
  const token = localStorage.getItem(ConfigName.ACCESS_TOKEN);
  const isAuthenticated = !!token;
  
  // Update auth store state
  authStore.isLogin = isAuthenticated;

  // Route requires authentication
  if (to.meta.requiresAuth) {
    if (!isAuthenticated) {
      // Not authenticated, redirect to login
      console.log('Redirecting to login - no auth');
      next({ name: 'login' });
    } else {
      // Authenticated, allow access
      next();
    }
  } 
  // Public route (login page)
  else {
    if (isAuthenticated && to.name === 'login') {
      // Already authenticated, redirect to profile
      console.log('Redirecting to profile - already authenticated');
      next({ name: 'profile' });
    } else {
      // Allow access to public route
      next();
    }
  }
};
