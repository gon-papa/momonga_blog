.PHONY: help build-local up down logs openapi
.DEFAULT_GOAL := help

build-local: ## Build local environment
	docker compose build --no-cache

up: ## Start local environment
	docker compose up -d

down: ## Stop local environment
	docker compose down

logs: ## Show logs
	docker compose logs -f

openapi_lint: ## Generate openapi
	docker run --rm -v ${PWD}:/spec redocly/cli lint openapi/openapi.yml

openapi_ignore: ## Generate openapi ignore
	docker run --rm -v ${PWD}:/spec redocly/cli lint --generate-ignore-file openapi/openapi.yml

openapi: ## Generate openapi
	make openapi_lint
	docker run --rm -v ${PWD}:/spec redocly/cli bundle openapi/openapi.yml -o ./openapi.yml
	cp ./openapi.yml ./src/openapi.yml

openapi_g:
	docker exec -it api bash -c "ogen -package ogen -target ogen -clean ./openapi.yml"