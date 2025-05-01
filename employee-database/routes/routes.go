package routes

import (
	"github.com/gin-gonic/gin"
	"hrSys/employee-database/auth"
	"hrSys/employee-database/handlers"
	"hrSys/employee-database/middleware"
)

func SetupRouter(r *gin.Engine, employeeHandler *handlers.EmployeeHandler) {
	// публичные
	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)

	// все /employees — только для авторизованных
	emp := r.Group("/employees")
	emp.Use(middleware.AuthMiddleware())
	{
		emp.GET("", employeeHandler.GetAllEmployees)
		emp.GET("/:id", employeeHandler.GetEmployeeByID)

		// только для админа
		emp.POST("", middleware.RoleMiddleware("admin"), employeeHandler.CreateEmployee)
		emp.PUT("/:id", middleware.RoleMiddleware("admin"), employeeHandler.UpdateEmployee)
		emp.DELETE("/:id", middleware.RoleMiddleware("admin"), employeeHandler.DeleteEmployee)
	}
}
