import { defineStore } from 'pinia';
import { getProfileService } from '../services/staff';
import { Permissions } from '../constants';

const defaultProfile = {
  id: 0,
  username: '',
  fullname: '',
  roleId: 0,
  roleLabel: '',
  isActive: false,
  isAdmin: false,
  permissions: [],
};

export const useStaffStore = defineStore('staff', {
  state: () => ({
    profile: { ...defaultProfile },
  }),
  getters: {
    permissions: (state) => state.profile.permissions || [],
    isAdmin: (state) => !!state.profile.isAdmin,
    hasPermission: (state) => (code) => {
      if (state.profile.isAdmin) return true;
      return (state.profile.permissions || []).includes(code);
    },
    canReadStaff() {
      return this.hasPermission(Permissions.STAFFS_READ);
    },
    canManageStaff() {
      return (
        this.hasPermission(Permissions.STAFFS_CREATE) ||
        this.hasPermission(Permissions.STAFFS_UPDATE) ||
        this.hasPermission(Permissions.STAFFS_DELETE)
      );
    },
    canReadRoles() {
      return this.hasPermission(Permissions.ROLES_READ);
    },
    canUpdateRoles() {
      return this.hasPermission(Permissions.ROLES_UPDATE);
    },
    canReadVisitors() {
      return this.hasPermission(Permissions.VISITORS_READ);
    },
  },
  actions: {
    async getProfileAction() {
      const res = await getProfileService();
      this.profile = {
        ...defaultProfile,
        ...res,
        permissions: res.permissions || [],
      };
      return this.profile;
    },
    clearProfile() {
      this.profile = { ...defaultProfile };
    },
  },
});
