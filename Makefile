.PHONY: "migrate-up"
migrate-up:
	goose -dir="./migrations" postgres "host=localhost user=postgres password=postgres dbname=l0 sslmode=disable" up

.PHONY: "migrate-reset"
migrate-reset:
	goose -dir="./migrations" postgres "host=localhost user=postgres password=postgres dbname=l0 sslmode=disable" reset

build:
	docker-compose up

prod:
	go run prod.go

start:
	go run main.go

build-env:
	cp .env-example .env