package routes

import (
	"github.com/gin-gonic/gin"
	"hrSys/employee-database/auth"
	"hrSys/employee-database/handlers"
	"hrSys/employee-database/middleware"
	"hrSys/employee-database/repositories"
	"hrSys/employee-database/services"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.POST("/login", auth.Login)

	employeeRepo := repositories.NewEmployeeRepository()
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)

	api := r.Group("/employees")
	api.Use(middleware.JwtAuthMiddleware())
	{
		api.POST("", employeeHandler.CreateEmployee)
		api.GET("", employeeHandler.GetAllEmployees)
		api.GET("/:id", employeeHandler.GetEmployeeByID)
		api.PUT("/:id", employeeHandler.UpdateEmployee)
		api.DELETE("/:id", employeeHandler.DeleteEmployee)
	}

	return r
}
