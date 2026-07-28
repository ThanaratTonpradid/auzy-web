import { axiosInstance } from '../plugins/axios';

export async function getProfileService() {
  const res = await axiosInstance.get('/api/staff/profile');
  return res.data;
}

export async function listStaffService() {
  const res = await axiosInstance.get('/api/staff');
  return res.data;
}

export async function getStaffService(id) {
  const res = await axiosInstance.get(`/api/staff/${id}`);
  return res.data;
}

export async function createStaffService(payload) {
  const res = await axiosInstance.post('/api/staff', payload);
  return res.data;
}

export async function updateStaffService(id, payload) {
  const res = await axiosInstance.put(`/api/staff/${id}`, payload);
  return res.data;
}

export async function deleteStaffService(id) {
  const res = await axiosInstance.delete(`/api/staff/${id}`);
  return res.data;
}
