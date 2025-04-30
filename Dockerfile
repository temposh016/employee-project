# Этап сборки
FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
COPY vendor ./vendor
ENV GOFLAGS=-mod=vendor

# Копируем миграции ДО сборки
COPY employee-database/db/migrations ./employee-database/db/migrations

COPY . .

RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main /main
COPY .env /app/.env

COPY employee-database/db/migrations /app/employee-database/db/migrations

WORKDIR /app
EXPOSE 8080
CMD ["/main"]
