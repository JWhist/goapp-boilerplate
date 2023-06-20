.PHONY: build run dev


dev:
	make build
	make run

run:
	docker-compose up go-api

build:
	docker build . -t go-api --build-arg GIT_CREDS=$(shell ./git-creds.sh);