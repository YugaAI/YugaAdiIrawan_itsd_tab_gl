package services

import (
	"fmt"
	"interview/BackendAPIHumanCapital/internal/models"
)

func (s *Service) GetEmployeeByNIK(nik string) (*models.Employee, error) {
	emp, err := s.Repo.GetEmployeeByNIK(nik)
	if err != nil {
		return nil, err
	}
	return emp, nil
}

func (s *Service) UpdateEmployeeByName(nik string, newName string) (string, error) {
	emp, err := s.Repo.GetEmployeeByNIK(nik)
	if err != nil {
		return "", fmt.Errorf("employee not found: %w", err)
	}

	oldName := emp.Nama
	if err := s.Repo.UpdateEmployeeByName(nik, newName); err != nil {
		return "", fmt.Errorf("failed to update employee: %w", err)
	}
	return oldName, nil
}
