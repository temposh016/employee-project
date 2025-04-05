package routes

import (
	"github.com/gin-gonic/gin"
	"hrSys/employee-database/handlers"
	"hrSys/employee-database/repositories"
	"hrSys/employee-database/services"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Инициализируем репозиторий и сервис
	employeeRepo := repositories.NewEmployeeRepository()
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)

	// Устанавливаем маршруты
	r.POST("/employees", employeeHandler.CreateEmployee)
	r.GET("/employees", employeeHandler.GetAllEmployees)
	r.GET("/employees/:id", employeeHandler.GetEmployeeByID)
	r.PUT("/employees/:id", employeeHandler.UpdateEmployee)
	r.DELETE("/employees/:id", employeeHandler.DeleteEmployee)

	return r
}
