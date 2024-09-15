.PHONY: help build-local up down logs
.DEFAULT_GOAL := help

build-local: ## Build local environment
	docker compose build --no-cache

up: ## Start local environment
	docker compose up -d

down: ## Stop local environment
	docker compose down

logs: ## Show logs
	docker compose logs -f
