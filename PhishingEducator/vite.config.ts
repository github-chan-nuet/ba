import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: "0.0.0.0", // Allow connections from outside the container
    port: 5173, // Match the port in docker-compose
    strictPort: true,
    watch: {
      usePolling: true, // Required for HMR in Docker on some OS
    },
  },
})
