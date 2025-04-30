package services

import (
	"errors"
	"gorm.io/gorm"
	"hrSys/employee-database/models"
)

type EmployeeService struct {
	DB *gorm.DB
}

func NewEmployeeService(db *gorm.DB) *EmployeeService {
	return &EmployeeService{DB: db}
}

func (s *EmployeeService) CreateEmployee(employee *models.Employee) (*models.Employee, error) {
	if err := s.DB.Create(employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	if err := s.DB.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (s *EmployeeService) GetEmployeeByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	if err := s.DB.First(&employee, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &employee, nil
}

func (s *EmployeeService) UpdateEmployee(id uint, employee *models.Employee) (*models.Employee, error) {
	var existing models.Employee
	if err := s.DB.First(&existing, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	existing.Name = employee.Name
	existing.Position = employee.Position
	existing.Department = employee.Department
	existing.Salary = employee.Salary

	if err := s.DB.Save(&existing).Error; err != nil {
		return nil, err
	}
	return &existing, nil
}

func (s *EmployeeService) DeleteEmployee(id uint) error {
	if err := s.DB.Delete(&models.Employee{}, id).Error; err != nil {
		return err
	}
	return nil
}
