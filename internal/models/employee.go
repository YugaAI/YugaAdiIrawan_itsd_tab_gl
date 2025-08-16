package models

type Employee struct {
	Nik          string `gorm:"primaryKey;not null" json:"nik"`
	Nama         string `gorm:"not null" json:"nama"`
	TanggalLahir string `gorm:"not null" json:"tanggal_lahir"`
	JenisKelamin string `gorm:"not null" json:"jenis_kelamin"`
	Departemen   string `gorm:"not null" json:"departemen"`
}

