run:
	@go run ./cmd/api

seed:
	@go run ./cmd/seed

build:
	@go build -o ./bin/software2backend ./cmd/api

swag:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@swag init -g cmd/api/main.go

test:
	@go test ./... -v

test-coverage:
	@go test ./... -coverprofile=coverage.out -covermode=atomic -v
	@go tool cover -func=coverage.out

.PHONY: run seed build swag test
