.PHONY: run
run:
	docker-compose -f docker-compose.yaml up --remove-orphans --build

rebuild: build
	docker-compose up --remove-orphans --build

.PHONY: stop
stop:
	docker-compose down --remove-orphans