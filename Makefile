include .env
MIGRATIONS_PATH = ./pkg/cmd/migrate/migrations
DC = docker compose
STORAGES_FILE = docker_compose/storages.yaml
ENV = --env-file .env

.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(name)

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATOR_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATOR_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: all
all:
	${DC} -f ${STORAGES_FILE} ${ENV} up --build -d
