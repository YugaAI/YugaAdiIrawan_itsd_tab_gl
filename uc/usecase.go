package uc

import "main.go/intranal/models"

type EmployeeRepository interface {
	GetEmployeeByNIK(nik string) (*models.Employee, error)
	UpdateEmployeeByName(nik string, newName string) (string, error)
}

type EmployeeUseCase struct {
	Repo EmployeeRepository
}

func NewEmployeeUseCase(repo EmployeeRepository) *EmployeeUseCase {
	return &EmployeeUseCase{
		Repo: repo,
	}
}

func (r *EmployeeUseCase) GetEmployeeByNIK(nik string, newName string) (string, string, error) {
	oldName, err := r.Repo.UpdateEmployeeByName(nik, newName)
	if err != nil {
		return "", "", err
	}
	return oldName, newName, nil
}
