package utils

import (
	"fmt"

	"gorm.io/gorm"
)

type PaginationData struct {
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	PageNumber int `json:"page"`
}

type PaginationError struct {
	message string
}

func (e PaginationError) Error() string {
	return fmt.Sprintln(e.message)
}

func GetPaginationData(page int, limit int) (*PaginationData, error) {
	offset := (page - 1) * limit

	if limit > 500 {
		return nil, &PaginationError{"limit cannot be greater than 500"}
	}

	if page < 1 {
		return nil, &PaginationError{"page cannot be less than 1"}
	}

	return &PaginationData{
		Offset:     offset,
		Limit:      limit,
		PageNumber: page,
	}, nil
}

func GetPagination(db *gorm.DB, page, limit int) (*gorm.DB, error) {
	pagination, err := GetPaginationData(page, limit)

	if err != nil {
		return nil, err
	}

	return db.Limit(pagination.Limit).Offset(pagination.Offset), nil
}
