export const loginRoute = {
  path: '/login',
  name: 'login',
  component: () => import('../pages/LoginPage.vue'),
  meta: { requiresAuth: false },
};
