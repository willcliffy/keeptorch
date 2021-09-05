TIMESTAMP ?= $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
SHA ?= $(shell git rev-parse --short HEAD)
DIST_DIR = ./dist

up:
	CGO_ENABLED=0 go build -ldflags "-s -w -X main.buildtime=$(TIMESTAMP) -X main.commitsha=$(SHA)" -o ${DIST_DIR}/api ./main.go

local:
	docker-compose up

local-rebuild:
	docker-compose up --build
