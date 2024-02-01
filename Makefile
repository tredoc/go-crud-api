include .env

migrate/up:
	@echo "Migrating up..."
	@migrate -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@0.0.0.0:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE}" -path ./db/migrations up

migrate/down:
	@echo "Migrating down..."
	@migrate -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@0.0.0.0:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE}" -path ./db/migrations down

.PHONY: migrate/up migrate/down
.SILENT: migrate/up migrate/down