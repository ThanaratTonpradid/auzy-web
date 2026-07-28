# auzy-web

Monorepo for Auzy admin web:

- `apps/admin-api` — Go admin API (`mini-api`)
- `apps/admin-ui` — Vue 3 + Vite admin UI

## Prerequisites

- Go 1.19+
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
# Start local infra (MySQL/Redis via docker-compose)
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
