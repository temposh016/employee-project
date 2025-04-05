package repositories

import (
	"github.com/jinzhu/gorm"
	"hrSys/employee-database/db"
	"hrSys/employee-database/models"
)

type EmployeeRepository interface {
	Create(employee *models.Employee) (*models.Employee, error)
	FindAll() ([]models.Employee, error)
	FindByID(id string) (*models.Employee, error)
	Update(employee *models.Employee) (*models.Employee, error)
	Delete(id string) error
}

type employeeRepository struct{}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{}
}

func (r *employeeRepository) Create(employee *models.Employee) (*models.Employee, error) {
	if err := db.DB.Create(employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (r *employeeRepository) FindAll() ([]models.Employee, error) {
	var employees []models.Employee
	if err := db.DB.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *employeeRepository) FindByID(id string) (*models.Employee, error) {
	var employee models.Employee
	if err := db.DB.First(&employee, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) Update(employee *models.Employee) (*models.Employee, error) {
	if err := db.DB.Save(employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (r *employeeRepository) Delete(id string) error {
	if err := db.DB.Delete(&models.Employee{}, id).Error; err != nil {
		return err
	}
	return nil
}
