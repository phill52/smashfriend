package utils

import (
	"math"

	"gorm.io/gorm"
)

var DB *gorm.DB

type PaginationData struct {
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	NextPage   int
	PrevPage   int
	CurrPage   int
	TotalPages int
}

func GetPaginationData(page int, limit int, model interface{}) PaginationData {
	// Calculate total pages
	var totalRows int64
	DB.Model(model).Count(&totalRows)
	totalPages := math.Ceil(float64(totalRows / int64(limit)))

	offset := (page - 1) * limit

	return PaginationData{
		Offset:     offset,
		Limit:      limit,
		NextPage:   page + 1,
		PrevPage:   page - 1,
		CurrPage:   page,
		TotalPages: int(totalPages),
	}
}
