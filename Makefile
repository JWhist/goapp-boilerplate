.PHONY: build build-ui run dev test clean

dev:
	make build
	make run

run:
	docker compose up go-api

build: build-ui
	docker build . -t go-api --build-arg GIT_CREDS=$(shell ./git-creds.sh);

build-ui:
	@echo "Building React UI..."
	cd ui && npm install && npm run build
	@echo "Copying UI build to static directory..."
	rm -rf static/ui
	cp -r ui/build static/ui
	@echo "UI build complete!"

build-local: build-ui
	@echo "Building Go binary locally..."
	go build -o bin/goapp-boilerplate ./main.go
	@echo "Local build complete! Run with: ./bin/goapp-boilerplate serve"

run-local: build-local
	./bin/goapp-boilerplate serve

clean:
	rm -rf ui/build
	rm -rf ui/node_modules
	rm -rf static/ui
	rm -rf bin/

test:
	go test --race ./...