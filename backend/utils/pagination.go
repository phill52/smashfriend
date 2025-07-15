package utils

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type PaginationData struct {
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	PageNumber int `json:"page"`
}

func GetPaginationData(page int, limit int) PaginationData {
	offset := (page - 1) * limit

	return PaginationData{
		Offset:     offset,
		Limit:      limit,
		PageNumber: page,
	}
}
