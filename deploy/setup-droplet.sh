#!/usr/bin/env bash
set -euo pipefail

# DigitalOcean Droplet bootstrap for Auzy (Debian 13 / 1GB RAM)
# Run as root on a fresh droplet:
#   curl -fsSL ... | bash
# or:
#   bash deploy/setup-droplet.sh

SWAP_SIZE_GB="${SWAP_SIZE_GB:-2}"
APP_DIR="${APP_DIR:-/opt/auzy-web}"
RETENTION_DAYS="${VISITOR_LOG_RETENTION_DAYS:-90}"

echo "==> Creating ${SWAP_SIZE_GB}G swap (if missing)"
if ! swapon --show | grep -q '/swapfile'; then
  fallocate -l "${SWAP_SIZE_GB}G" /swapfile || dd if=/dev/zero of=/swapfile bs=1M count=$((SWAP_SIZE_GB * 1024))
  chmod 600 /swapfile
  mkswap /swapfile
  swapon /swapfile
  if ! grep -q '/swapfile' /etc/fstab; then
    echo '/swapfile none swap sw 0 0' >> /etc/fstab
  fi
else
  echo "Swap already configured"
fi

echo "==> Installing Docker"
if ! command -v docker >/dev/null 2>&1; then
  apt-get update
  apt-get install -y ca-certificates curl
  install -m 0755 -d /etc/apt/keyrings
  curl -fsSL https://download.docker.com/linux/debian/gpg -o /etc/apt/keyrings/docker.asc
  chmod a+r /etc/apt/keyrings/docker.asc
  . /etc/os-release
  echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/debian ${VERSION_CODENAME} stable" \
    > /etc/apt/sources.list.d/docker.list
  apt-get update
  apt-get install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin
  systemctl enable --now docker
else
  echo "Docker already installed"
fi

echo "==> App directory: ${APP_DIR}"
mkdir -p "${APP_DIR}"
if [ ! -f "${APP_DIR}/deploy/.env" ] && [ -f "${APP_DIR}/deploy/.env.example" ]; then
  cp "${APP_DIR}/deploy/.env.example" "${APP_DIR}/deploy/.env"
  echo "Created deploy/.env from example — edit secrets before starting."
fi

echo "==> Installing visitor log retention cron (${RETENTION_DAYS} days)"
RETENTION_SCRIPT="${APP_DIR}/deploy/retention-visitor-logs.sh"
chmod +x "${RETENTION_SCRIPT}" || true
CRON_LINE="15 3 * * * VISITOR_LOG_RETENTION_DAYS=${RETENTION_DAYS} APP_DIR=${APP_DIR} ${RETENTION_SCRIPT} >> /var/log/auzy-visitor-retention.log 2>&1"
(crontab -l 2>/dev/null | grep -v 'retention-visitor-logs.sh' || true; echo "${CRON_LINE}") | crontab -

echo "==> Done"
echo "Next:"
echo "  1. Clone/copy repo into ${APP_DIR}"
echo "  2. Edit ${APP_DIR}/deploy/.env"
echo "  3. docker compose -f ${APP_DIR}/deploy/docker-compose.prod.yml --env-file ${APP_DIR}/deploy/.env up -d --build"
echo "Tip: use CPU Option Regular on SGP1 (Premium AMD unavailable)."
