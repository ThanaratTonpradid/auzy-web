#!/usr/bin/env bash
set -euo pipefail

# Delete visitor_logs older than N days from the MySQL container.
APP_DIR="${APP_DIR:-/opt/auzy-web}"
COMPOSE_FILE="${APP_DIR}/deploy/docker-compose.prod.yml"
ENV_FILE="${APP_DIR}/deploy/.env"
DAYS="${VISITOR_LOG_RETENTION_DAYS:-90}"

if [ ! -f "${COMPOSE_FILE}" ]; then
  echo "Compose file not found: ${COMPOSE_FILE}" >&2
  exit 1
fi

if [ -f "${ENV_FILE}" ]; then
  set -a
  # shellcheck disable=SC1090
  source "${ENV_FILE}"
  set +a
fi

CUTOFF=$(($(date +%s) - DAYS * 86400))
echo "$(date -Is) deleting visitor_logs older than ${DAYS} days (created_at < ${CUTOFF})"

docker compose -f "${COMPOSE_FILE}" --env-file "${ENV_FILE}" exec -T mysql \
  mysql -u"${MYSQL_USER:-auzy}" -p"${MYSQL_PASSWORD}" "${MYSQL_DATABASE:-auzy}" \
  -e "DELETE FROM visitor_logs WHERE created_at < ${CUTOFF};"
