run:
	@go run cmd/api/main.go

seed:
	@go run cmd/seed/*.go

build:
	@go build -o ./software2backend ./cmd/api/main.go

.PHONY: run seed