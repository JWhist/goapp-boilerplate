.PHONY: run

run:
	docker build . -t go-api --build-arg GIT_CREDS=$(shell ./git-creds.sh);
	docker run -d -p 3000:3000 go-api