CONTAINER_CLI?=docker
TAG?=$(shell git rev-parse --short HEAD)
GOLANGCILINTER_BINARY=golangci-lint

all: lint vet test fmt

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test -race ./...

lint:
	$(GOLANGCILINTER_BINARY) run ./...

cover:
	go test -race -coverprofile cover.out -coverpkg=./... ./...
	go tool cover -html=cover.out