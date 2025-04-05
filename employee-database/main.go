package main

import (
	"fmt"
	"hrSys/employee-database/db"
	"hrSys/employee-database/routes"
	"log"
	"time"
)

func main() {
	// Пытаемся подключиться к базе данных несколько раз с задержкой
	var err error
	for i := 0; i < 5; i++ {
		err = db.Init()
		if err == nil {
			fmt.Println("Успешное подключение к базе данных!")
			break
		}

		log.Println("Ошибка при подключении к базе данных, пробую снова...")
		time.Sleep(5 * time.Second) // Задержка между попытками
	}

	// Если ошибка подключения сохраняется, завершаем работу
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v\n", err)
	}

	// Настройка маршрутов
	r := routes.SetupRouter()

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	} else {
		fmt.Println("Server started at http://localhost:8080")
	}
}
