CONFIG_FILE ?= ./local.yaml
DSN ?= $(shell sed -n 's/^dsn:[[:space:]]\(.*\)/\1/p' $(CONFIG_FILE))
MIGRATE := migrate -database $(DSN) -path ./migrations

migrate:
	@echo "Running database migrations."
	@$(MIGRATE) up

local-test:
	docker-compose up -d
	sleep 2
	make migrate

mock:
	go install github.com/vektra/mockery/v2@latest
	cd database/; \
		mockery --name=DatabaseInterface --filename=DatabaseInterface.go;