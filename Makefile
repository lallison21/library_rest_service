include scripts/library_rest_service.mk

DOCKER_COMPOSE?=deployment/development/docker-compose.yml

.PHONY: up down create-migration migration-up migration-down

up:
	docker compose -f ${DOCKER_COMPOSE} down -v
	docker compose -f ${DOCKER_COMPOSE} up -d library-db
	docker compose -f ${DOCKER_COMPOSE} --profile tools run --rm migrate up
	docker compose -f ${DOCKER_COMPOSE} up -d --build --force-recreate library-rest-service
	docker compose -f ${DOCKER_COMPOSE} up -d grafana loki

down:
	docker compose -f deployment/development/docker-compose.yml down -v

create-migration:
	migrate create -ext sql -dir migrations -seq ${MIGRATION_NAME}

migration-up:
	docker compose -f ${DOCKER_COMPOSE} --profile tools run --rm migrate up

migration-down:
	docker compose -f ${DOCKER_COMPOSE} --profile tools run --rm migrate down

.DEFAULT_GOAL := up
