import { createRouter, createWebHistory } from 'vue-router';
import { profileRoute } from './profile.route';
import { loginRoute } from './auth.route';
import { routerHook } from './router-hook';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: { name: 'login' },
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
      children: [profileRoute],
      meta: { requiresAuth: true },
    },
  ],
});

// GOOD
router.beforeEach(routerHook);

export default router;
