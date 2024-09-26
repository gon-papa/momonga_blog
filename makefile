.PHONY: help build-local up down logs openapi
.DEFAULT_GOAL := help
include .env

help: ## Show Commands
	@grep -E -h '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build-local: ## Build local environment
	docker compose build --no-cache

up: ## Start local environment
	docker compose up -d

down: ## Stop local environment
	docker compose down

logs: ## Show logs
	docker compose logs -f

openapi_lint: ## openapi lint
	docker run --rm -v ${PWD}:/spec redocly/cli lint openapi/openapi.yml

openapi_ignore: ## Generate openapi ignore
	docker run --rm -v ${PWD}:/spec redocly/cli lint --generate-ignore-file openapi/openapi.yml

openapi: ## Generate openapi
	make openapi_lint
	docker run --rm -v ${PWD}:/spec redocly/cli bundle openapi/openapi.yml -o ./openapi.yml
	cp ./openapi.yml ./src/openapi.yml

openapi_g: ## Generate openapi go
	docker exec -it api bash -c "ogen -package api -target api -clean ./openapi.yml"

migrate_up: ## Migrate up
	docker exec -it api bash -c "migrate --path database/migration --database '${DBMS}://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}' --verbose up"

migrate_down: ## Migrate down
	docker exec -it api bash -c "migrate --path database/migration --database '${DBMS}://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}' --verbose down 1"

migrate_create: ## Migrate create
	docker exec -it api bash -c "migrate create -ext sql -dir database/migration -seq ${name}"

generate_secret_key: ## Generate secret key
	docker exec -it api bash -c "openssl rand -hex 32"