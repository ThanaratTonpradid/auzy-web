export const profileRoute = {
  path: '/profile',
  name: 'profile',
  component: () => import('../pages/ProfilePage.vue'),
  meta: { requiresAuth: true },
};
