run:
	@go run ./cmd/api
seed:
	@go run ./cmd/seed
build:
	@go build -o ./bin/software2backend ./cmd/api

swag:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@swag init -g cmd/api/main.go

.PHONY: run seed build swag
