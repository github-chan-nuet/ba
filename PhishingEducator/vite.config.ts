import { defineConfig } from "vite";
import { reactRouter } from "@react-router/dev/vite";
import { cjsInterop } from "vite-plugin-cjs-interop";
import path from "path";

export default defineConfig({
  plugins: [
    reactRouter(),

    // Add CJS interop plugin for Fluent UI packages until they are ESM compatible
    cjsInterop({
      dependencies: ["@fluentui/react-components"]
    })
  ],

  optimizeDeps: {
    exclude: ["@griffel/react"],
    include: ["@fluentui/react-components", "@fluentui/react-icons", "@hey-api/client-fetch", "framer-motion", 'react-chartjs-2', 'chart.js']
  },

  server: {
    host: "0.0.0.0", // Allow connections from outside the container
    port: 5173, // Match the port in docker-compose
    strictPort: true,
    watch: {
      usePolling: true, // Required for HMR in Docker on some OS
    },
  },

  resolve: {
    alias: {
      '@styles': path.resolve(__dirname, 'app/styles'),
      '@components': path.resolve(__dirname, 'app/components'),
      '@assets': path.resolve(__dirname, 'app/assets')
    }
  }
});