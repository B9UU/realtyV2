include .env

## start_docker: starts the docker images is CONTAINER_ID (postgresql)
.PHONY: start_docker
start_docker:
	docker start ${CONTAINER_ID}

## migrations: execute migrate command with with args (i.e: up 1, down 2, force, version..)
.PHONY: migrations
migrations:
	@echo "Running migrate " ${A}
	@migrate -database=${POSTGRES_URL} -path=./migrations ${A}

## migrations/new name=$1: crate a new database migration
.PHONY: migrations/new
migrations/new:
	@echo "Creating migration files for ${name}..."
	migrate create -seq -ext=.sql -dir=./migrations ${name}


# run: runs the application
.PHONY: run
run:
	@echo "Starting the application"
	@go run ./cmd/api

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## seed: to seed the db
.PHONY: seed
seed:
	psql ${POSTGRES_URL} -f ./seed.sql
