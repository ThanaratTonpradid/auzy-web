BINARY_FILE = auzy-api
GOPATH_BIN := $(shell go env GOPATH)/bin
SWAG := $(shell command -v swag 2>/dev/null || echo "$(GOPATH_BIN)/swag")

export PATH := $(GOPATH_BIN):$(PATH)

clean:
	@rm -rf $(BINARY_FILE) coverage.out

build:
	@go build
	@ls -alh $(BINARY_FILE)

cover:
	@go test ./... -cover -coverprofile=coverage.out
	@go tool cover -html=coverage.out

docs:
	@if command -v swag >/dev/null 2>&1 || [ -x "$(GOPATH_BIN)/swag" ]; then \
		"$(SWAG)" init --instanceName API --ot go -d ./internal/api,./lib -g ./route/swagger.route.go -o ./internal/api/doc; \
		"$(SWAG)" fmt -g ./internal/api/route/swagger.route.go; \
	else \
		echo "swag not found — skipping docs generation"; \
		echo "install with: go install github.com/swaggo/swag/cmd/swag@latest"; \
	fi

models:
	go run cmd/generate/generate.go

tables:
	go run cmd/generate-table/generate-table.go

infra-up:
	docker compose -p mini -f .development/docker-compose.yml up --build -d

infra-down:
	docker compose -p mini -f .development/docker-compose.yml down

api: docs
	@go run main.go api
