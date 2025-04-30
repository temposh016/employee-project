package repositories

import (
	"gorm.io/gorm"
	"hrSys/employee-database/models"
)

type EmployeeRepository interface {
	Create(e *models.Employee) (*models.Employee, error)
	FindAll() ([]models.Employee, error)
	FindByID(id uint) (*models.Employee, error)
	Update(e *models.Employee) (*models.Employee, error)
	Delete(id uint) error
}

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepo{db: db}
}

func (r *employeeRepo) Create(e *models.Employee) (*models.Employee, error) {
	if err := r.db.Create(e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (r *employeeRepo) FindAll() ([]models.Employee, error) {
	var employees []models.Employee
	if err := r.db.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *employeeRepo) FindByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	if err := r.db.First(&employee, id).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepo) Update(e *models.Employee) (*models.Employee, error) {
	if err := r.db.Save(e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (r *employeeRepo) Delete(id uint) error {
	if err := r.db.Delete(&models.Employee{}, id).Error; err != nil {
		return err
	}
	return nil
}
