package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"hrSys/employee-database/auth"
	"hrSys/employee-database/db"
	"hrSys/employee-database/handlers"
	"hrSys/employee-database/middleware"
	"hrSys/employee-database/services"
)

func main() {
	db.InitDB()

	employeeService := services.NewEmployeeService(db.DB)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)

	router := gin.Default()
	router.POST("/login", auth.Login)

	protected := router.Group("/", middleware.JwtAuthMiddleware())
	{
		protected.POST("/employees", employeeHandler.CreateEmployee)
		protected.GET("/employees", employeeHandler.GetAllEmployees)
		protected.GET("/employees/:id", employeeHandler.GetEmployeeByID)
		protected.PUT("/employees/:id", employeeHandler.UpdateEmployee)
		protected.DELETE("/employees/:id", middleware.RoleMiddleware("admin"), employeeHandler.DeleteEmployee)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
