package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// DB переменная для подключения к базе данных
var DB *gorm.DB

// Инициализация подключения к базе данных
func Init() error {
	// Загружаем переменные из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Получаем значения из переменных окружения
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// Подключаемся к базе данных
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return err // Возвращаем ошибку, если подключение не удалось
	}
	DB = db
	return nil // Возвращаем nil, если подключение успешно
}
