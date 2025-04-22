import { createClient } from '@hey-api/openapi-ts';

try {
  await createClient({
    input: "../api/phishing-backend-open-api.yaml",
    output: "./src/api",
    plugins: ["@hey-api/client-fetch"]
  })
} catch (e) {
  console.error("Error generating the API client", e);
}