.PHONY: run
run-prod:
	docker-compose -f docker-compose.prod.yaml up --remove-orphans --build

run-local:
	docker-compose -f docker-compose.local.yaml up --remove-orphans --build

rebuild: build
	docker-compose up --remove-orphans --build

.PHONY: stop
stop:
	docker-compose down --remove-orphans