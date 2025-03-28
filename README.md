# Securaware

## Docker compose
```bash
# force rebuilding the images
docker compose -f ./.docker/compose.dev.yaml up --build
# use existing images
docker compose -f ./.docker/compose.dev.yaml up
# stop
docker compose -f ./.docker/compose.dev.yaml down
```