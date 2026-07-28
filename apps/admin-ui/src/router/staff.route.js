export const staffRoute = {
  path: 'staff',
  name: 'staff',
  component: () => import('../pages/StaffPage.vue'),
  meta: { requiresAuth: true, permission: 'STAFFS_READ' },
};
