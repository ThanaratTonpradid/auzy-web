import { axiosInstance } from '../plugins/axios';

export async function loginService({ username, password }) {
  const res = await axiosInstance.post('/api/auth/login', { username, password });
  return res.data;
}

export async function logoutService() {
  const res = await axiosInstance.post('/api/auth/logout');
  return res.data;
}
