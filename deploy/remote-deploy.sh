#!/usr/bin/env bash
# Run on the droplet after code is synced to APP_DIR.
set -euo pipefail

APP_DIR="${APP_DIR:-/opt/auzy-web}"
COMPOSE_FILE="${APP_DIR}/deploy/docker-compose.prod.yml"
ENV_FILE="${APP_DIR}/deploy/.env"

if [[ ! -f "${ENV_FILE}" ]]; then
  echo "Missing ${ENV_FILE} — copy from deploy/.env.example and set secrets first."
  exit 1
fi

cd "${APP_DIR}/deploy"
docker compose -f "${COMPOSE_FILE}" --env-file "${ENV_FILE}" up -d --build --remove-orphans
docker compose -f "${COMPOSE_FILE}" --env-file "${ENV_FILE}" ps
echo "Deploy finished."
