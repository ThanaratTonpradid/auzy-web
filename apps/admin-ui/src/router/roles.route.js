export const rolesRoute = {
  path: '/roles',
  name: 'roles',
  component: () => import('../pages/RolesPage.vue'),
  meta: { requiresAuth: true, permission: 'ROLES_READ' },
};
