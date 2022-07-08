package repository

import (
	models "dexa-training-crud-karyawan/pkg/models/entities"
)

type Repository interface {
	SaveKaryawan(m *models.Karyawan) error
	ShowKaryawanPagination() ([]*models.Karyawan, error)
	UpdateKaryawan(update *models.Karyawan) error
	DeleteKaryawan(delete *models.Karyawan) error
}
