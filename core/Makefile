MIGRATIONS_DIR = ./cmd/migrate/migrations
DB_PATH = .insight/data.db

.PHONY: migrate-create
migrate-create:
	@goose --dir $(MIGRATIONS_DIR) -s create $(filter-out $@, $(MAKECMDGOALS)) sql

.PHONY: migrate-up
migrate-up:
	@goose --dir $(MIGRATIONS_DIR) sqlite3 $(DB_PATH) up

.PHONY: migrate-down
migrate-down:
	@goose --dir $(MIGRATIONS_DIR) sqlite3 $(DB_PATH) down

%:
	@: