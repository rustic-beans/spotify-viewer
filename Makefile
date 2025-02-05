generate:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc generate
	go run github.com/99designs/gqlgen generate

start:
	go run ./cmd/main.go

setup:
	docker-compose up -d
	go mod tidy
	$(MAKE) generate
