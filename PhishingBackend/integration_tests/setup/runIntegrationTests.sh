#!/bin/bash
echo "delete docker database volume..."
docker volume rm setup_it_db

echo "starting docker database container..."
docker compose -f ./compose.integration.yaml up -d

echo "loading environment variables from integrationTests.env"
# https://stackoverflow.com/questions/43267413/how-to-set-environment-variables-from-env-file/76407401
set -a
source integrationTests.env
set +a

echo "starting integration tests..."
go test -tags integration ./../reminder/... -v

echo "stopping docker database container"
docker compose -f ./compose.integration.yaml down