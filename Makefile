.PHONY: build run dev test


dev:
	make build
	make run

run:
	docker-compose up go-api

build:
	docker build . -t go-api --build-arg GIT_CREDS=$(shell ./git-creds.sh);

test:
	go test --race ./...