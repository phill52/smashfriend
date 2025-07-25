package utils

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

type PaginationData struct {
	Offset     int   `json:"offset"`
	Limit      int   `json:"limit"`
	PageNumber int   `json:"page"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int   `json:"totalPages"`
}

type PaginationError struct {
	message string
}

func (e PaginationError) Error() string {
	return fmt.Sprint(e.message)
}

func GetPaginationData(db *gorm.DB, page int, limit int) (*PaginationData, error) {
	offset := (page - 1) * limit

	if limit > 500 {
		return nil, &PaginationError{"limit cannot be greater than 500"}
	}

	if page < 1 {
		return nil, &PaginationError{"page cannot be less than 1"}
	}

	var totalItems int64
	db.Count(&totalItems)
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	return &PaginationData{
		Offset:     offset,
		Limit:      limit,
		PageNumber: page,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}, nil
}

func PaginateData(db *gorm.DB, page, limit int) (*gorm.DB, error) {
	pagination, err := GetPaginationData(db, page, limit)
	if err != nil {
		return nil, err
	}

	return db.Limit(pagination.Limit).Offset(pagination.Offset), nil
}
