THIS_FILE := $(lastword $(MAKEFILE_LIST))
.PHONY: help build up start down destroy stop restart logs logs-api ps login-timescale login-api db-shell

build:
	docker compose build $(c)
down:
	docker compose down
up:
	docker compose up $(c)
create-migration:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up
migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

run-accepted-tests:
	go test -v ./test

generate:
	oapi-codegen -package=http -generate=chi-server,types,spec api/task_manager_api.yaml > internal/app/handlers/http/openapi.gen.go