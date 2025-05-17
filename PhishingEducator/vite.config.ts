import { defineConfig } from "vite";
import { reactRouter } from "@react-router/dev/vite";
import { cjsInterop } from "vite-plugin-cjs-interop";

export default defineConfig({
  plugins: [
    reactRouter(),

    // Add CJS interop plugin for Fluent UI packages until they are ESM compatible
    cjsInterop({
      dependencies: ["@fluentui/react-components"]
    })
  ],

  // Required for Fluent UI icons in SSR
  ssr: {
    noExternal: ['@fluentui/react-icons'],
  },

  server: {
    host: "0.0.0.0", // Allow connections from outside the container
    port: 5173, // Match the port in docker-compose
    strictPort: true,
    watch: {
      usePolling: true, // Required for HMR in Docker on some OS
    },
  },
});