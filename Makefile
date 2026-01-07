main_file := ./cmd/main.go

generate:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc generate
	go run github.com/99designs/gqlgen generate

start:
	go run $(main_file)

watch:
	gow run $(main_file)

setup:
	docker-compose up -d
	go mod tidy
	$(MAKE) generate

test:
	go test ./...
