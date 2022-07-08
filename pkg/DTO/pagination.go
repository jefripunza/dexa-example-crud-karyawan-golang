package models

import models "dexa-training-crud-karyawan/pkg/models/entities"

type Pagination struct {
	Data        []models.Karyawan `json:"data"        validate:"required"`
	TotalData   int               `json:"totalData"   validate:"required"`
	CurrentPage int               `json:"currentPage" validate:"required"`
	LastPage    int               `json:"lastPage"    validate:"required"`
	NextPage    int               `json:"nextPage"    validate:"required"`
	PrevPage    int               `json:"prevPage"    validate:"required"`
}
