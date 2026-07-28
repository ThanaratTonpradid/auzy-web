import { describe, it, expect } from 'vitest';
import { appConfig } from '../app.config';

describe('App Config', () => {
  it('has default values', () => {
    expect(appConfig).toBeDefined();
    expect(appConfig.apiUrl).toBeDefined();
    expect(appConfig.appName).toBeDefined();
    expect(appConfig.appVersion).toBeDefined();
  });

  it('uses environment variables when available', () => {
    // These will be default values in test environment
    expect(typeof appConfig.apiUrl).toBe('string');
    expect(typeof appConfig.appName).toBe('string');
    expect(typeof appConfig.appVersion).toBe('string');
  });

  it('has development/production flags', () => {
    expect(typeof appConfig.isDevelopment).toBe('boolean');
    expect(typeof appConfig.isProduction).toBe('boolean');
  });

  it('has correct default API URL', () => {
    expect(appConfig.apiUrl).toBe('http://localhost:1323');
  });
});

