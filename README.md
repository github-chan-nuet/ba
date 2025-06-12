# Securaware

## Structure
- ./docker: application virtualization (docker) related configurations
- api: HTTP API specification that the backend exposes in [OpenAPI format](https://swagger.io/specification/) 
- PhishingBackend: source code of the Phishing Backend that responds to HTTP requests and forwards emails.
- PhishingEducator: source code of the Phishing Educator which is responsible for a single page web application.
- PhishingProxy: source code of the proxy that setups a traefik instance

## Development
### Docker compose
```bash
# force rebuilding the images
docker compose -f ./.docker/compose.dev.yaml up --build
# use existing images
docker compose -f ./.docker/compose.dev.yaml up
# stop
docker compose -f ./.docker/compose.dev.yaml down
```