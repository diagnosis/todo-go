FROM ubuntu:latest
# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-list-backend ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /todo-list-backend .
COPY .env .
COPY db/migrations ./db/migrations

EXPOSE 8080

CMD ["./todo-list-backend"]