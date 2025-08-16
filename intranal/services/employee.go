package services

import "main.go/intranal/models"

func (s *Service) GetEmployeeByNIK(nik string) (*models.Employee, error) {
	emp, err := s.Repo.GetEmployeeByNIK(nik)
	if err != nil {
		return nil, err
	}
	return emp, nil
}

func (s *Service) UpdateEmployeeByName(nik string, newName string) (string, error) {
	oldName, err := s.Repo.UpdateEmployeeByName(nik, newName)
	if err != nil {
		return "", err
	}
	return oldName, nil
}
