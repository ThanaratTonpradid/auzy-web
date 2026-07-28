export const notFoundRoute = {
  path: '/404',
  name: 'not-found',
  component: () => import('../pages/NotFoundPage.vue'),
  meta: { requiresAuth: false },
};

export const catchAllRoute = {
  path: '/:pathMatch(.*)*',
  redirect: { name: 'not-found' },
};

export const pubCatchAllRoute = {
  path: ':pathMatch(.*)*',
  redirect: { name: 'not-found' },
};

export const priCatchAllRoute = {
  path: ':pathMatch(.*)*',
  redirect: { name: 'not-found' },
};
