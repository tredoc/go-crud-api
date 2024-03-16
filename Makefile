include .env

migrate/up:
	@echo "Migrating up..."
	@migrate -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@0.0.0.0:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE}" -path ./db/migrations up

migrate/down:
	@echo "Migrating down..."
	@migrate -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@0.0.0.0:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE}" -path ./db/migrations down

build:
	@echo "Building..."
	@GOOS=linux GOARCH=amd64 go build -o ./bin/main ./cmd/app

run/dev:
	@echo "Running in development mode..."
	@GOOS=linux GOARCH=amd64 go build -o ./bin/main ./cmd/app
	@docker compose up

mock/service:
	@echo "Running mock service..."
	@mockery --dir=internal/service --output=mocks/service --outpkg=mockservice --all

mock/repository:
	@echo "Running mock service..."
	@mockery --dir=internal/repository --output=mocks/repository --outpkg=mockrepository --all

test:
	@echo "Running tests..."
	@go test -v -cover ./...

swag:
	@echo "Generating swagger..."
	@swag init -o ./docs/swagger -g ./cmd/app/main.go

.PHONY: migrate/up migrate/down run/dev mock/service mock/repository swag
.SILENT: migrate/up migrate/down run/dev mock/service mock/repository swag