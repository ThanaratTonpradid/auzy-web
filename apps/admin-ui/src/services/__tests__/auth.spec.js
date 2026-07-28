import { describe, it, expect, beforeEach, vi } from 'vitest';
import { loginService, logoutService } from '../auth';
import { axiosInstance } from '../../plugins/axios';

// Mock axios
vi.mock('../../plugins/axios', () => ({
  axiosInstance: {
    post: vi.fn(),
  },
}));

describe('Auth Service', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  describe('loginService', () => {
    it('sends correct login request', async () => {
      const mockResponse = {
        data: {
          token: 'mock-token',
          refreshToken: 'mock-refresh-token',
          expiresIn: 3600,
        },
      };

      axiosInstance.post.mockResolvedValue(mockResponse);

      const credentials = { username: 'testuser', password: 'password123' };
      const result = await loginService(credentials);

      expect(axiosInstance.post).toHaveBeenCalledWith('/api/auth/login', credentials);
      expect(result).toEqual(mockResponse.data);
    });

    it('throws error on login failure', async () => {
      const mockError = new Error('Invalid credentials');
      axiosInstance.post.mockRejectedValue(mockError);

      const credentials = { username: 'testuser', password: 'wrongpassword' };

      await expect(loginService(credentials)).rejects.toThrow('Invalid credentials');
    });
  });

  describe('logoutService', () => {
    it('sends logout request', async () => {
      const mockResponse = {
        data: {
          success: true,
          message: 'Logged out successfully',
        },
      };

      axiosInstance.post.mockResolvedValue(mockResponse);

      const result = await logoutService();

      expect(axiosInstance.post).toHaveBeenCalledWith(
        '/api/auth/logout',
        {},
        expect.any(Object)
      );
      expect(result).toEqual(mockResponse.data);
    });

    it('handles logout error gracefully', async () => {
      const mockError = new Error('Network error');
      axiosInstance.post.mockRejectedValue(mockError);

      await expect(logoutService()).rejects.toThrow('Network error');
    });
  });
});

