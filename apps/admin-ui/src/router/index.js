import { createRouter, createWebHistory } from 'vue-router';
import { profileRoute } from './profile.route';
import { loginRoute } from './auth.route';
import { dashboardRoute } from './dashboard.route';
import { staffRoute } from './staff.route';
import { rolesRoute } from './roles.route';
import { routerHook } from './router-hook';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: { name: 'dashboard' },
    },
    {
      path: '/pub',
      component: () => import('../layouts/PublicLayout.vue'),
      children: [loginRoute],
      meta: { requiresAuth: false },
    },
    {
      path: '/pri',
      component: () => import('../layouts/PrivateLayout.vue'),
      children: [dashboardRoute, profileRoute, staffRoute, rolesRoute],
      meta: { requiresAuth: true },
    },
  ],
});

router.beforeEach(routerHook);

export default router;
