import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/items': 'http://localhost:8080',
      '/item': 'http://localhost:8080',
    },
    port: 3000,
  },
});
