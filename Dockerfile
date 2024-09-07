FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./software2backend ./cmd/api/main.go
 
 
FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/software2backend .
EXPOSE 8080
ENTRYPOINT ["./software2backend"]