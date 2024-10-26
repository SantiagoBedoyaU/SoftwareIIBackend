FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./software2backend ./cmd/api

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/software2backend .
EXPOSE 8000
ENTRYPOINT ["./software2backend"]
