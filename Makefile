run:
	@go run cmd/api/main.go

seed:
	@go run cmd/seed/*.go

.PHONY: run seed