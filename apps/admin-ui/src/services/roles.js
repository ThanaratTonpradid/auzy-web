import { axiosInstance } from '../plugins/axios';

export async function listRolesService() {
  const res = await axiosInstance.get('/api/roles');
  return res.data;
}

export async function getRoleService(id) {
  const res = await axiosInstance.get(`/api/roles/${id}`);
  return res.data;
}

export async function updateRolePermissionsService(id, permissionIds) {
  const res = await axiosInstance.put(`/api/roles/${id}/permissions`, {
    permissionIds,
  });
  return res.data;
}

export async function listPermissionsService() {
  const res = await axiosInstance.get('/api/permissions');
  return res.data;
}
