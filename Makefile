include scripts/library_rest_service.mk

DOCKER_COMPOSE?=deployment/development/docker-compose.yml

.PHONY: up down

up:
	docker compose -f ${DOCKER_COMPOSE} down -v
	docker compose -f ${DOCKER_COMPOSE} up -d --build --force-recreate library_rest_service
	docker compose -f ${DOCKER_COMPOSE} up -d grafana loki

down:
	docker compose -f deployment/development/docker-compose.yml down -v

.DEFAULT_GOAL := up
