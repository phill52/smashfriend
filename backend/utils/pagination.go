package utils

import (
	"errors"

	"gorm.io/gorm"
)

type PaginationData struct {
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	PageNumber int `json:"page"`
}

func GetPaginationData(page int, limit int) (*PaginationData, error) {
	offset := (page - 1) * limit

	if limit > 500 {
		return nil, errors.New("limit parameter is greater than max 500")
	}

	if page < 1 {
		return nil, errors.New("page parameter cannot be less than 1")
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
		return nil, errors.New("error")
	}

	return db.Limit(pagination.Limit).Offset(pagination.Offset), nil
}
