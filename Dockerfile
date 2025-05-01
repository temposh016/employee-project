FROM golang:1.23-alpine AS builder
WORKDIR /app

# копируем go.mod/go.sum
COPY go.mod go.sum ./

# копируем vendor из корня проекта
COPY vendor ./vendor

# включаем vendoring
ENV GOFLAGS=-mod=vendor

# копируем миграции
COPY employee-database/db/migrations ./employee-database/db/migrations

# копируем остальной код (включая main.go, auth, middleware и т.д.)
COPY . .

# сборка с модулем vendor
RUN go build -mod=vendor -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main /main
COPY .env /app/.env
COPY employee-database/db/migrations /app/employee-database/db/migrations

WORKDIR /app
EXPOSE 8080
CMD ["/main"]
