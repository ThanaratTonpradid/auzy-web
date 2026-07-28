import { axiosInstance } from '../plugins/axios';

export async function recordVisitService({ path, referer } = {}) {
  await axiosInstance.post('/api/public/visit', {
    path: path || window.location.pathname,
    referer: referer || document.referrer || '',
  });
}

export async function listVisitorLogsService({ page = 1, limit = 20 } = {}) {
  const res = await axiosInstance.get('/api/visitor-logs', {
    params: { page, limit },
  });
  return res.data;
}
