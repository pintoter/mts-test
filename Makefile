.DEFAULT_GOAL := re

ifeq ($(version), prod)
	DOCKER_COMPOSE_FILE = -f docker-compose.prod.yaml
else
	DOCKER_COMPOSE_FILE = -f docker-compose.local.yaml
endif

.PHONY: run
run:
	docker-compose ${DOCKER_COMPOSE_FILE} up --remove-orphans --build

.PHONY: down
stop:
	docker-compose ${DOCKER_COMPOSE_FILE} down

.PHONY: re
re: stop run

.PHONY: lint
lint:
	golangci-lint run ./order-service/...
	golangci-lint run ./store-service/...