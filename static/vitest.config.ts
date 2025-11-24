import { defineConfig } from 'vitest/config';

export default defineConfig({
  test: {
    globals: true,
    environment: 'jsdom',
    include: ['js/datatable/tests/**/*.test.ts'],
    setupFiles: ['js/datatable/tests/setup.ts'],
    testTimeout: 10000,
  },
  resolve: {
    alias: {
      '@': new URL('./js/datatable/src', import.meta.url).pathname,
    },
  },
});