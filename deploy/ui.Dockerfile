FROM node:22-alpine AS build

WORKDIR /app

ARG VITE_API_URL=
ARG VITE_APP_NAME=Auzy
ARG VITE_PROFILE_NAME=Auzy
ARG VITE_PROFILE_TITLE=Personal profile
ARG VITE_PROFILE_BIO=A quiet corner of the internet.

ENV VITE_API_URL=$VITE_API_URL \
    VITE_APP_NAME=$VITE_APP_NAME \
    VITE_PROFILE_NAME=$VITE_PROFILE_NAME \
    VITE_PROFILE_TITLE=$VITE_PROFILE_TITLE \
    VITE_PROFILE_BIO=$VITE_PROFILE_BIO

COPY apps/admin-ui/package.json apps/admin-ui/pnpm-lock.yaml* ./
COPY package.json pnpm-workspace.yaml pnpm-lock.yaml .npmrc ./
COPY apps/admin-ui ./apps/admin-ui

RUN corepack enable && corepack prepare pnpm@11.17.0 --activate \
  && pnpm install --filter app... \
  && pnpm --filter app build

FROM nginx:1.27-alpine

COPY deploy/nginx/default.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/apps/admin-ui/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
