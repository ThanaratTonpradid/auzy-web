import { createRouter, createWebHistory } from 'vue-router';
import { profileRoute } from './profile.route';
import { loginRoute } from './auth.route';
import { dashboardRoute } from './dashboard.route';
import { staffRoute } from './staff.route';
import { rolesRoute } from './roles.route';
import { publicProfileRoute } from './public-profile.route';
import { visitorLogsRoute } from './visitor-logs.route';
import { catchAllRoute, notFoundRoute, priCatchAllRoute, pubCatchAllRoute } from './not-found.route';
import { routerHook } from './router-hook';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: () => import('../layouts/PublicLayout.vue'),
      children: [publicProfileRoute],
      meta: { requiresAuth: false },
    },
    {
      path: '/pub',
      component: () => import('../layouts/PublicLayout.vue'),
      children: [loginRoute, pubCatchAllRoute],
      meta: { requiresAuth: false },
    },
    {
      path: '/pri',
      component: () => import('../layouts/PrivateLayout.vue'),
      children: [
        dashboardRoute,
        profileRoute,
        staffRoute,
        rolesRoute,
        visitorLogsRoute,
        priCatchAllRoute,
      ],
      meta: { requiresAuth: true },
    },
    notFoundRoute,
    catchAllRoute,
  ],
});

router.beforeEach(routerHook);

export default router;
