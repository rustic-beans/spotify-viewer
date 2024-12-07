generate:
	sqlc generate
	go run github.com/99designs/gqlgen generate

start:
	go run ./cmd/main.go

