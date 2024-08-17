.PHONY: up build test clean

include ./configs/.env
export

DB_URL="postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" 

up:build
	@bin/api	
run:
	@go run cmd/api/main.go

build:
	@go build -o bin/api cmd/api/main.go

test:
	go test -v ./...

clean:
	rm -rf bin

db-init:
	@echo "Initializing database..."
	@psql -U postgres -c "SELECT 1 FROM pg_user WHERE usename = '$(DB_USER)'" | grep -q 1 || \
		psql -U postgres -c "CREATE USER $(DB_USER) WITH PASSWORD '$(DB_PASS)'"
	@psql -U postgres -c "SELECT 1 FROM pg_database WHERE datname = '$(DB_NAME)'" | grep -q 1 || \
		( psql -U postgres -c "CREATE DATABASE $(DB_NAME)" && \
		  psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE $(DB_NAME) TO $(DB_USER)" )
	@psql -U postgres -d $(DB_NAME) -c "GRANT ALL ON SCHEMA public TO $(DB_USER)"
	@echo "Database initialized successfully."

db-drop:
	@echo "Dropping database and user..."
	@psql -U postgres -c "DROP DATABASE IF EXISTS $(DB_NAME)"
	@psql -U postgres -c "DROP USER IF EXISTS $(DB_USER)"
	@echo "Database and user dropped successfully."


# migrations related commands
migrate-up:
	migrate -path ./db/migrations -database $(DB_URL) up

migrate-down:
	migrate -path ./db/migrations -database $(DB_URL) down

migrate-create:
	migrate create -ext sql -dir ./db/migrations -seq $(name)

migrate-force:
	migrate -path ./db/migrations -database $(DB_URL) force $(version)

migrate-drop:
	migrate -path ./db/migrations -database $(DB_URL) drop