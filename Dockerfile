# Этап сборки (builder)
FROM golang:1.23-alpine AS builder

# Установим рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod tidy

# Копируем весь код
COPY . ./

# Сборка приложения
RUN go build -o main .

# Этап с минимальным образом (stage-2)
FROM alpine:latest

# Устанавливаем сертификаты
RUN apk --no-cache add ca-certificates

# Копируем собранный файл из предыдущего этапа
COPY --from=builder /app/main /main
COPY .env .env

# Открываем порт
EXPOSE 8080

# Запуск приложения
CMD ["/main"]
