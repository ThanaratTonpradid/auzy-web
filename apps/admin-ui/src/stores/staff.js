import { defineStore } from 'pinia';
import { getProfileService } from '../services/staff';

const defaultValue = {
    id: 0,
    username: '',
    fullname: '',
    roleId: 0,
    isActive: false,
};

export const useStaffStore = defineStore('staff', {
  state: () => ({
    profile: defaultValue,
    permission: [],
  }),
  actions: {
    async getProfileAction() {
      try {
        const res = await getProfileService();
        this.profile = {
            ...this.profile,
            ...res,
        };
      } catch (error) {
        if (error.response.data.errorCode === "UNAUTHORIZED") {
            this.profile = defaultValue;
        }
      }
    },
  },
});
