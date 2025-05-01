package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"hrSys/employee-database/db"
	"hrSys/employee-database/handlers"
	"hrSys/employee-database/repositories"
	"hrSys/employee-database/routes"
	"hrSys/employee-database/services"
)

func main() {
	// 1) Инициализируем БД
	db.InitDB()

	// 2) Создаём репозиторий, сервис и хендлер
	empRepo := repositories.NewEmployeeRepository(db.DB)
	empService := services.NewEmployeeService(empRepo)
	empHandler := handlers.NewEmployeeHandler(empService)

	// 3) Настраиваем Gin и роуты, передавая готовый handler
	r := gin.Default()
	routes.SetupRouter(r, empHandler)

	// 4) Запускаем
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
