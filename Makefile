include .env

.PHONY: api

install:
	brew install pre-commit
	brew install golangci-lint
	pre-commit install --hook-type pre-commit --config .githooks/pre-commit-config.yaml
	pre-commit install --hook-type pre-push --config .githooks/pre-push-config.yaml

api:
	oapi-codegen -generate gorilla -package api openapi.yaml > api/server.gen.go
	oapi-codegen -generate types -package api openapi.yaml > api/types.gen.go

run:
	go run main.go

lint:
	golangci-lint run ./...

format:
	golangci-lint run ./... --fix
