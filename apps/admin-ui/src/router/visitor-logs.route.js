export const visitorLogsRoute = {
  path: 'visitor-logs',
  name: 'visitor-logs',
  component: () => import('../pages/VisitorLogsPage.vue'),
  meta: { requiresAuth: true, permission: 'VISITORS_READ' },
};
