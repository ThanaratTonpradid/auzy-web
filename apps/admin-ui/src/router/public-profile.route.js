export const publicProfileRoute = {
  path: '',
  name: 'public-profile',
  component: () => import('../pages/PublicProfilePage.vue'),
  meta: { requiresAuth: false },
};
