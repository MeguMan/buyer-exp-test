.PHONY: build postgres adminer migrate

build:
	go build -v ./cmd/apiserver

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=12345 postgres

adminer:
	docker run --rm -ti --network host adminer

migrate:
	migrate -path migrations -database postgres://postgres:12345@192.168.65.0:5432/docker_test?sslmode=disable up

.DEFAULT_GOAL := build
