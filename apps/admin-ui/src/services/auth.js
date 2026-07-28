import { axiosInstance } from '../plugins/axios';
import { ConfigName } from '../constants';

const token = localStorage.getItem(ConfigName.ACCESS_TOKEN);

export async function loginService({ username, password }) {
  const res = await axiosInstance.post('/api/auth/login', { username, password });
  return res.data;
}

export async function logoutService() {
  const res = await axiosInstance.post('/api/auth/logout', {}, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return res.data;
}
