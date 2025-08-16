package repository

import (
	"fmt"

	"main.go/intranal/models"
)

func (r *repository) GetEmployeeByNIK(nik string) (*models.Employee, error) {
	employee := models.Employee{}
	res := r.db.Where("nik = ?", nik).First(&employee)
	if res.Error != nil {
		return nil, res.Error
	}
	return &employee, nil
}
func (r *repository) UpdateEmployeeByName(nik string, newName string) (string, error) {
	var emp models.Employee
	if err := r.db.Where("nik = ?", nik).First(&emp).Error; err != nil {
		return "", fmt.Errorf("employee with nik %s not found: %w", nik, err)
	}

	oldName := emp.Nama
	emp.Nama = newName
	if err := r.db.Save(&emp).Error; err != nil {
		return "", fmt.Errorf("failed to update employee name: %w", err)
	}

	return oldName, nil
}
