package services

import "interview/BackendAPIHumanCapital/internal/models"

type EmployeeRepository interface {
	GetEmployeeByNIK(nik string) (*models.Employee, error)
	UpdateEmployeeByName(nik string, newName string) error
}

type Service struct {
	Repo EmployeeRepository
}

func NewService(repo EmployeeRepository) *Service {
	return &Service{
		Repo: repo,
	}
}
