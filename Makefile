.PHONY: up down lint enter
up:
	docker-compose up --watch

down:
	docker-compose down

lint:
	gofmt -s -w .
	golangci-lint run ./...

enter:
	docker-compose exec db /bin/sh