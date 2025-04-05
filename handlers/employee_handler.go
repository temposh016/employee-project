package handlers

import (
	"github.com/gin-gonic/gin"
	"hrSys/employee-database/models"
	"hrSys/employee-database/services"
	"net/http"
)

type EmployeeHandler struct {
	employeeService services.EmployeeService
}

func NewEmployeeHandler(service services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		employeeService: service,
	}
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdEmployee, err := h.employeeService.CreateEmployee(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return
	}

	c.JSON(http.StatusOK, createdEmployee)
}

func (h *EmployeeHandler) GetAllEmployees(c *gin.Context) {
	employees, err := h.employeeService.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return
	}

	c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	employee, err := h.employeeService.GetEmployeeByID(id)
	if err != nil || employee == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedEmployee, err := h.employeeService.UpdateEmployee(id, &employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
		return
	}

	c.JSON(http.StatusOK, updatedEmployee)
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	if err := h.employeeService.DeleteEmployee(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}
