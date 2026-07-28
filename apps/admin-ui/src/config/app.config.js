// Application configuration
export const appConfig = {
  apiUrl: import.meta.env.VITE_API_URL || 'http://localhost:1323',
  appName: import.meta.env.VITE_APP_NAME || 'Auzy Admin',
  appVersion: import.meta.env.VITE_APP_VERSION || '1.0.0',
  isDevelopment: import.meta.env.DEV === true || import.meta.env.DEV === 'true',
  isProduction: import.meta.env.PROD === true || import.meta.env.PROD === 'true',
};

// Validate required configuration
const validateConfig = () => {
  const requiredVars = ['apiUrl'];
  const missing = requiredVars.filter(key => !appConfig[key]);
  
  if (missing.length > 0) {
    console.warn('Missing required environment variables:', missing);
  }
};

validateConfig();

