ifneq (,$(wildcard ./.env))
	include .env
	export
endif

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
.DEFAULT_GOAL := help

POSTGRES_VERSION=16
DB_IMAGE=bitnami/postgresql:$(POSTGRES_VERSION)
DB_CONTAINER_NAME=fold-sd-pg

help: ## Show this help.
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  ${YELLOW}%-24s${GREEN}%s${RESET}\n", $$1, $$2}' $(MAKEFILE_LIST)

net_up: ## create local db net
	docker network create -d bridge $(LOCAL_NET_NAME);

net_down: ## remove local db net
	docker network rm $(LOCAL_NET_NAME);

database_up: ## starts postgres database in a container and seeds it appropriately
	$(eval DB_CONTAINER_STATUS := $(shell docker ps -a --format '{{.State}}' --filter "name=^/$(DB_CONTAINER_NAME)$$"))
	if [[ "$(DB_CONTAINER_STATUS)" == "" ]]; then \
	docker run -d --name $(DB_CONTAINER_NAME) \
	--network $(LOCAL_NET_NAME) \
	-p $(POSTGRESQL_PORT):5432 \
	-e POSTGRESQL_POSTGRES_PASSWORD=$(POSTGRESQL_POSTGRES_PASSWORD) \
	-e POSTGRESQL_USERNAME=$(POSTGRESQL_USERNAME) \
	-e POSTGRESQL_PASSWORD=$(POSTGRESQL_PASSWORD) \
	-e POSTGRESQL_DATABASE=$(POSTGRESQL_DATABASE) \
	-e POSTGRESQL_SEED_FILENAME=/tmp/seed.sql \
	$(DB_IMAGE); \
	fi;
	if [[ "$(DB_CONTAINER_STATUS)" == "exited" ]]; then \
	docker start $(DB_CONTAINER_NAME); \
	fi

database_down: ## stops postgres database container
	docker stop $(DB_CONTAINER_NAME)

database_nuke: ## nukes postgres database container
	docker rm --force $(DB_CONTAINER_NAME)

database_backup: ## backups database locally
	docker exec -it \
	-e PGPASSWORD=$(POSTGRESQL_POSTGRES_PASSWORD) \
	$(DB_CONTAINER_NAME) \
	pg_dump \
	--clean \
	--if-exists \
	--no-privileges \
	--no-owner \
    --no-password \
    --verbose \
    --username postgres \
    --host localhost \
    --dbname $(POSTGRESQL_DATABASE) \
    --format c \
    --file /tmp/backup.sql;
	docker cp $(DB_CONTAINER_NAME):/tmp/backup.sql "`date +"%Y-%m-%d_%H-%M-00"`.sql";

database_logs: ## logs db server
	docker logs -f $(DB_CONTAINER_NAME)

psql: ## psql into database
	docker exec \
	-e PGPASSWORD=$(POSTGRESQL_PASSWORD) \
	-it $(DB_CONTAINER_NAME) psql \
	-U $(POSTGRESQL_USERNAME) \
	-d $(POSTGRESQL_DATABASE)

install: ## install dependencies
	go mod tidy

build: ## build all images
	@docker-compose build

up: ## stack up
	@docker-compose up
