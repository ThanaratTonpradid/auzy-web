export const locationRoute = {
  path: '',
  name: 'location',
  component: () => import('../pages/LocationPage.vue'),
  meta: { requiresAuth: false },
};
