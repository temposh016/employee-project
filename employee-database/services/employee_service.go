package services

import (
	"hrSys/employee-database/models"
	"hrSys/employee-database/repositories"
)

type EmployeeService struct {
	repo repositories.EmployeeRepository
}

// NewEmployeeService теперь принимает репозиторий, а не *gorm.DB
func NewEmployeeService(repo repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(e *models.Employee) (*models.Employee, error) {
	return s.repo.Create(e)
}

func (s *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.repo.FindAll()
}

func (s *EmployeeService) GetEmployeeByID(id uint) (*models.Employee, error) {
	emp, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if emp == nil {
		// можно возвращать nil,nil, чтобы хэндлер сам делал 404
		return nil, nil
	}
	return emp, nil
}

func (s *EmployeeService) UpdateEmployee(id uint, data *models.Employee) (*models.Employee, error) {
	// сначала проверим, есть ли запись
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, nil
	}
	// обновляем поля
	existing.Name = data.Name
	existing.Position = data.Position
	existing.Department = data.Department
	existing.Salary = data.Salary

	return s.repo.Update(existing)
}

func (s *EmployeeService) DeleteEmployee(id uint) error {
	// можно тоже сначала проверить, есть ли запись, но не обязательно
	return s.repo.Delete(id)
}
