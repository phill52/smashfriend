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

func GetPaginationData(page int, limit int, model interface{}) PaginationData {
	// Calculate total pages
	// var totalRows int64
	// DB.Model(model).Count(&totalRows)
	// fmt.Printf("total rows %d", totalRows)
	//totalPages := math.Ceil(float64(totalRows) / float64(limit))
	//fmt.Printf("total pages %g", totalPages)

	offset := (page - 1) * limit

	return PaginationData{
		Offset:     offset,
		Limit:      limit,
		PageNumber: page,
	}
}
