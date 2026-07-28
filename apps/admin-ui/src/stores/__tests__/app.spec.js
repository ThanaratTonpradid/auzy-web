import { describe, it, expect, beforeEach } from 'vitest';
import { setActivePinia, createPinia } from 'pinia';
import { useAppStore } from '../app';

describe('App Store', () => {
  beforeEach(() => {
    setActivePinia(createPinia());
  });

  it('initializes with default state', () => {
    const store = useAppStore();
    expect(store.navIsOpen).toBe(false);
    expect(store.loading).toBe(false);
    expect(store.error).toBe(null);
    expect(store.notification.show).toBe(false);
  });

  it('toggles navigation', () => {
    const store = useAppStore();
    expect(store.navIsOpen).toBe(false);
    
    store.toggleNavAction();
    expect(store.navIsOpen).toBe(true);
    
    store.toggleNavAction();
    expect(store.navIsOpen).toBe(false);
  });

  it('sets loading state', () => {
    const store = useAppStore();
    
    store.setLoading(true);
    expect(store.loading).toBe(true);
    
    store.setLoading(false);
    expect(store.loading).toBe(false);
  });

  it('sets and clears error', () => {
    const store = useAppStore();
    const errorMessage = 'Test error message';
    
    store.setError(errorMessage);
    expect(store.error).toBe(errorMessage);
    expect(store.notification.show).toBe(true);
    expect(store.notification.type).toBe('error');
    
    store.clearError();
    expect(store.error).toBe(null);
  });

  it('shows notification with correct type', () => {
    const store = useAppStore();
    
    store.showNotification('Success message', 'success', 5000);
    expect(store.notification.show).toBe(true);
    expect(store.notification.message).toBe('Success message');
    expect(store.notification.type).toBe('success');
    expect(store.notification.timeout).toBe(5000);
  });

  it('hides notification', () => {
    const store = useAppStore();
    
    store.showNotification('Test message', 'info');
    expect(store.notification.show).toBe(true);
    
    store.hideNotification();
    expect(store.notification.show).toBe(false);
  });
});

