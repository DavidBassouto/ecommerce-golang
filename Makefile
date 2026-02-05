.PHONY: help build run dev migrate-up migrate-down

help:
	@echo "Usage:"
	@echo "  make build    - Build the application"
	@echo "  make run      - Run the application"
	@echo "  make dev      - Run the application in development mode"
	@echo "  make lint     - Run linting on the code"
	@echo "  make migrate-up   - Run database migrations (up)"
	@echo "  make migrate-down - Run database migrations (down)"

build:
	go build -o bin/app ./cmd/api

run:
	go run ./cmd/api

dev:
	go run ./cmd/api

lint:
	golangci-lint run ./...

migrate-up:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/ecommerce_shop?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/ecommerce_shop?sslmode=disable" -verbose down

docker-up:
	docker compose -f docker/docker-compose.yml up -d

docker-down:
	docker compose -f docker/docker-compose.yml down
