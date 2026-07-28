import axios from 'axios';
import { ConfigName } from '../constants';
import { appConfig } from '../config/app.config';
import router from '../router';
import { i18n } from './i18n';

// Create axios instance
export const axiosInstance = axios.create({
  baseURL: appConfig.apiUrl,
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor
axiosInstance.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem(ConfigName.ACCESS_TOKEN);
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Track if we're currently refreshing the token
let isRefreshing = false;
let failedQueue = [];

const processQueue = (error, token = null) => {
  failedQueue.forEach(prom => {
    if (error) {
      prom.reject(error);
    } else {
      prom.resolve(token);
    }
  });
  
  failedQueue = [];
};

// Response interceptor
axiosInstance.interceptors.response.use(
  (response) => {
    return response;
  },
  async (error) => {
    const originalRequest = error.config;
    
    if (error.response) {
      // Server responded with error status
      const { status, data } = error.response;
      
      // Handle token refresh for 401 errors
      if (status === 401 && !originalRequest._retry) {
        if (isRefreshing) {
          // If already refreshing, queue this request
          return new Promise((resolve, reject) => {
            failedQueue.push({ resolve, reject });
          }).then(token => {
            originalRequest.headers.Authorization = `Bearer ${token}`;
            return axiosInstance(originalRequest);
          }).catch(err => {
            return Promise.reject(err);
          });
        }

        originalRequest._retry = true;
        isRefreshing = true;

        const refreshToken = localStorage.getItem(ConfigName.REFRESH_TOKEN);
        
        if (!refreshToken) {
          // No refresh token, redirect to login
          localStorage.removeItem(ConfigName.ACCESS_TOKEN);
          localStorage.removeItem(ConfigName.REFRESH_TOKEN);
          if (router.currentRoute.value.name !== 'login') {
            router.push({ name: 'login' });
          }
          return Promise.reject(error);
        }

        try {
          // Try to refresh the token
          const response = await axios.post(
            `${appConfig.apiUrl}/api/auth/refresh`,
            { refreshToken }
          );
          
          const { token, refreshToken: newRefreshToken } = response.data;
          
          // Save new tokens
          localStorage.setItem(ConfigName.ACCESS_TOKEN, token);
          localStorage.setItem(ConfigName.REFRESH_TOKEN, newRefreshToken);
          
          // Update authorization header
          axiosInstance.defaults.headers.common['Authorization'] = `Bearer ${token}`;
          originalRequest.headers.Authorization = `Bearer ${token}`;
          
          // Process queued requests
          processQueue(null, token);
          
          isRefreshing = false;
          
          // Retry original request
          return axiosInstance(originalRequest);
        } catch (refreshError) {
          // Refresh failed, clear tokens and redirect to login
          processQueue(refreshError, null);
          isRefreshing = false;
          
          localStorage.removeItem(ConfigName.ACCESS_TOKEN);
          localStorage.removeItem(ConfigName.REFRESH_TOKEN);
          
          if (router.currentRoute.value.name !== 'login') {
            router.push({ name: 'login' });
          }
          
          return Promise.reject(refreshError);
        }
      }
      
      switch (status) {
        case 401:
          error.message = data?.message || i18n.global.t('error.unauthorized');
          break;
        case 403:
          error.message = data?.message || i18n.global.t('error.forbidden');
          break;
        case 404:
          error.message = data?.message || i18n.global.t('error.notFound');
          break;
        case 422:
          error.message = data?.message || i18n.global.t('error.validationError');
          error.validationErrors = data?.errors || {};
          break;
        case 500:
          error.message = data?.message || i18n.global.t('error.serverError');
          break;
        default:
          error.message = data?.message || i18n.global.t('error.unknownError');
      }
    } else if (error.request) {
      // Request was made but no response received
      error.message = i18n.global.t('error.networkError');
    } else {
      // Something else happened
      error.message = error.message || i18n.global.t('error.unknownError');
    }
    
    return Promise.reject(error);
  }
);

// Legacy support - will be deprecated
const token = localStorage.getItem(ConfigName.ACCESS_TOKEN);
export const axiosInstanceWithToken = axios.create({
  baseURL: appConfig.apiUrl,
  headers: {
    Authorization: `Bearer ${token}`,
  },
});
