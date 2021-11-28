.PHONY: all

SHELL=/bin/bash -e

.DEFAULT_GOAL := help

-include .env

help: ## Справка
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

env: ## Создаёт .env
	@if [ ! -f ./.env ]; then \
		cp .env.example .env; \
	fi

up: ## Запуск проекта
	docker-compose up -d

down: ## Остановка всех контейнеров проекта
	docker-compose down

rb: down up ## Перезапуск
