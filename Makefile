.PHONY: install dev-ui build-ui lint-ui test-ui api build-api infra-up infra-down

install:
	pnpm install

dev-ui:
	pnpm --filter app dev

build-ui:
	pnpm --filter app build

lint-ui:
	pnpm --filter app lint

test-ui:
	pnpm --filter app test:unit

api:
	$(MAKE) -C apps/admin-api api

build-api:
	$(MAKE) -C apps/admin-api build

infra-up:
	$(MAKE) -C apps/admin-api infra-up

infra-down:
	$(MAKE) -C apps/admin-api infra-down
