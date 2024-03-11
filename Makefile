include .env

migration_source ?= "file://db/migrations"
migration_destination ?= "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@localhost:3306/$(MYSQL_DATABASE)"

# Takes the first target as command
Command := $(firstword $(MAKECMDGOALS))
# Skips the first word
Arguments := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

migrate:
	migrate -source $(migration_source) -database $(migration_destination) up

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(Arguments)

docker-up:
	docker compose up -d

docker-down:
	docker compose down