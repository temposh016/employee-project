package services

import (
	"hrSys/employee-database/models"
	"hrSys/employee-database/repositories"
)

type EmployeeService interface {
	CreateEmployee(employee *models.Employee) (*models.Employee, error)
	GetAllEmployees() ([]models.Employee, error)
	GetEmployeeByID(id string) (*models.Employee, error)
	UpdateEmployee(id string, employee *models.Employee) (*models.Employee, error)
	DeleteEmployee(id string) error
}

type employeeService struct {
	employeeRepo repositories.EmployeeRepository
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeService {
	return &employeeService{
		employeeRepo: repo,
	}
}

func (s *employeeService) CreateEmployee(employee *models.Employee) (*models.Employee, error) {
	return s.employeeRepo.Create(employee)
}

func (s *employeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.employeeRepo.FindAll()
}

func (s *employeeService) GetEmployeeByID(id string) (*models.Employee, error) {
	return s.employeeRepo.FindByID(id)
}

func (s *employeeService) UpdateEmployee(id string, employee *models.Employee) (*models.Employee, error) {
	return s.employeeRepo.Update(employee)
}

func (s *employeeService) DeleteEmployee(id string) error {
	return s.employeeRepo.Delete(id)
}
