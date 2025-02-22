include .env
export $(shell sed 's/=.*//' .env)

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL)
MIGRATE=migrate -database "$(DB_URL)" -path internal/infrastructure/storage/postgres/migrations/

.PHONY: migrate-up migrate-down migrate-create load-env

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

migrate-create:
	@read -p "Введите имя миграции: " name; \
	migrate create -ext sql -dir migrations -seq $$name