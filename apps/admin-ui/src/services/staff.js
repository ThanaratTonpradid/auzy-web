import { axiosInstance } from '../plugins/axios';
import { ConfigName } from '../constants';

const token = localStorage.getItem(ConfigName.ACCESS_TOKEN);

export async function getSessionService() {
  const res = await axiosInstance.get('/api/staff/profile', {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return res.data;
}
export async function getProfileService() {
  const res = await axiosInstance.get('/api/staff/profile', {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return res.data;
}
