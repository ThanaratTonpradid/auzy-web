# auzy-web

Monorepo for Auzy admin web:

- `apps/admin-api` — Go admin API (`auzy-api`)
- `apps/admin-ui` — Vue 3 + Vite admin UI

## Prerequisites

- Go 1.25+
- Node.js 18+
- [pnpm](https://pnpm.io/)
- Docker (for local API infra)

## Setup

```sh
# Frontend deps
pnpm install

# Backend env
cp apps/admin-api/local.env.sample apps/admin-api/local.env
```

## Development

```sh
# Start local infra (MySQL 9.7 / Redis 8.8 via docker-compose)
pnpm infra:up

# API (from repo root)
pnpm dev:api

# UI
pnpm dev:ui
```

Or work inside each app:

```sh
cd apps/admin-api && make api
cd apps/admin-ui && pnpm dev
```

## Build

```sh
pnpm build:api
pnpm build:ui
```

## Production (DigitalOcean $6 droplet)

Target: Basic Shared CPU Regular, 1 vCPU / 1GB RAM / 25GB SSD, SGP1, Debian 13.

```sh
# On the droplet (as root)
bash deploy/setup-droplet.sh

# Copy repo to /opt/auzy-web, then:
cp deploy/.env.example deploy/.env
# edit secrets in deploy/.env
docker compose -f deploy/docker-compose.prod.yml --env-file deploy/.env up -d --build
```

Notes:
- Public profile is at `/` and records visits via `POST /api/public/visit`
- Admin UI is at `/pub/login` then `/pri/*`
- Visitor log retention cron is installed by `setup-droplet.sh` (default 90 days)
- If upgrading MySQL/Redis major versions locally, reset `apps/admin-api/.development/data/*` first
