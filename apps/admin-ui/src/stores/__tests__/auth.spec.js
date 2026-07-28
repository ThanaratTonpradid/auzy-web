import { describe, it, expect, beforeEach, vi } from 'vitest';
import { setActivePinia, createPinia } from 'pinia';
import { useAuthStore } from '../auth';
import * as authService from '../../services/auth';
import { ConfigName } from '../../constants';

// Mock router
vi.mock('../../router', () => ({
  default: {
    replace: vi.fn(),
  },
}));

// Mock auth services
vi.mock('../../services/auth', () => ({
  loginService: vi.fn(),
  logoutService: vi.fn(),
}));

// Mock localStorage
const localStorageMock = (() => {
  let store = {};
  return {
    getItem: vi.fn((key) => store[key] || null),
    setItem: vi.fn((key, value) => {
      store[key] = value.toString();
    }),
    removeItem: vi.fn((key) => {
      delete store[key];
    }),
    clear: vi.fn(() => {
      store = {};
    }),
  };
})();

global.localStorage = localStorageMock;

describe('Auth Store', () => {
  beforeEach(() => {
    setActivePinia(createPinia());
    localStorageMock.clear();
    vi.clearAllMocks();
  });

  it('initializes with logged out state', () => {
    const store = useAuthStore();
    expect(store.isLogin).toBe(false);
    expect(store.user).toBe(null);
  });

  it('logs in successfully', async () => {
    const mockResponse = {
      token: 'mock-token',
      refreshToken: 'mock-refresh-token',
      expiresIn: 3600,
      user: { id: 1, username: 'testuser' },
    };

    authService.loginService.mockResolvedValue(mockResponse);

    const store = useAuthStore();
    await store.loginAction({ username: 'testuser', password: 'password' });

    expect(store.isLogin).toBe(true);
    expect(store.user).toEqual(mockResponse.user);
    expect(localStorageMock.setItem).toHaveBeenCalledWith(
      ConfigName.ACCESS_TOKEN,
      mockResponse.token
    );
    expect(localStorageMock.setItem).toHaveBeenCalledWith(
      ConfigName.REFRESH_TOKEN,
      mockResponse.refreshToken
    );
  });

  it('handles login failure', async () => {
    authService.loginService.mockRejectedValue(new Error('Invalid credentials'));

    const store = useAuthStore();
    
    await expect(
      store.loginAction({ username: 'testuser', password: 'wrongpassword' })
    ).rejects.toThrow('Invalid credentials');

    expect(store.isLogin).toBe(false);
    expect(store.user).toBe(null);
  });

  it('logs out successfully', async () => {
    authService.logoutService.mockResolvedValue({ success: true });

    // Setup logged in state
    localStorageMock.setItem(ConfigName.ACCESS_TOKEN, 'mock-token');
    localStorageMock.setItem(ConfigName.REFRESH_TOKEN, 'mock-refresh-token');

    const store = useAuthStore();
    store.isLogin = true;
    store.user = { id: 1, username: 'testuser' };

    await store.logoutAction();

    expect(store.isLogin).toBe(false);
    expect(store.user).toBe(null);
    expect(localStorageMock.removeItem).toHaveBeenCalledWith(ConfigName.ACCESS_TOKEN);
    expect(localStorageMock.removeItem).toHaveBeenCalledWith(ConfigName.REFRESH_TOKEN);
  });

  it('checks auth status correctly', () => {
    const store = useAuthStore();

    // Not logged in
    expect(store.checkAuth()).toBe(false);

    // Logged in
    localStorageMock.setItem(ConfigName.ACCESS_TOKEN, 'mock-token');
    expect(store.checkAuth()).toBe(true);
  });
});

