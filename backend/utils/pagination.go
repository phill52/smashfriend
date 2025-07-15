package utils

import (
	"errors"

	"gorm.io/gorm"
)

var DB *gorm.DB

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
