/**
 * Test Setup
 * Global setup for all tests
 */

import { vi } from 'vitest';

// Mock console methods to reduce noise in tests
const originalConsoleError = console.error;
const originalConsoleWarn = console.warn;

// Setup DOM environment mocks
beforeEach(() => {
  // Reset all mocks before each test
  vi.clearAllMocks();
  
  // Restore console methods
  console.error = originalConsoleError;
  console.warn = originalConsoleWarn;
  
  // Setup localStorage mock
  const localStorageMock = {
    getItem: vi.fn(),
    setItem: vi.fn(),
    removeItem: vi.fn(),
    clear: vi.fn(),
    length: 0,
    key: vi.fn(),
  };
  
  // Mock window.localStorage
  Object.defineProperty(global, 'localStorage', {
    value: localStorageMock,
    writable: true,
  });
  
  // Ensure window exists for jsdom
  if (typeof window !== 'undefined') {
    Object.defineProperty(window, 'localStorage', {
      value: localStorageMock,
      writable: true,
    });
  }
});

// Mock performance if not available
if (!global.performance) {
  global.performance = {
    now: vi.fn(() => Date.now()),
  } as any;
}

// Mock requestAnimationFrame if not available
if (!global.requestAnimationFrame) {
  global.requestAnimationFrame = vi.fn((callback) => {
    setTimeout(callback, 16);
    return 1;
  });
}

if (!global.cancelAnimationFrame) {
  global.cancelAnimationFrame = vi.fn();
}

// Clean up after all tests
afterAll(() => {
  vi.restoreAllMocks();
});