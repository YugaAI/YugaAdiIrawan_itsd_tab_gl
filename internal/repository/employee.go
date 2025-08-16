package repository

import "interview/BackendAPIHumanCapital/internal/models"

func (r *repository) GetEmployeeByNIK(nik string) (*models.Employee, error) {
	employee := models.Employee{}
	res := r.db.Where("nik = ?", nik).First(&employee)
	if res.Error != nil {
		return nil, res.Error
	}
	return &employee, nil
}
func (r *repository) UpdateEmployeeByName(nik string, newName string) error {
	if err := r.db.Model(&models.Employee{}).
		Where("nik = ?", nik).
		Update("nama", newName).Error; err != nil {
		return err
	}
	return nil
}
