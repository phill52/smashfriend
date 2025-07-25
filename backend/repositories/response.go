package repositories

import (
	"smashfriend/database"
	"smashfriend/models"
	"smashfriend/utils"
)

func GetResponse(model interface{}, page int, limit int, statusCode int, message string) (*models.Response, error) {
	query := database.DB.Model(&model)

	paginationData, err := utils.GetPaginationData(query, page, limit)
	if err != nil {
		return nil, err
	}
	meta, err := GetMetaData(message, statusCode)
	if err != nil {
		return nil, err
	}
	return &models.Response{
		Data:       model,
		Pagination: *paginationData,
		Meta:       *meta,
	}, nil
}

func GetResponseWithoutPagination(model interface{}, statusCode int, message string) (*models.Response, error) {

	meta, err := GetMetaData(message, statusCode)
	if err != nil {
		return nil, err
	}
	return &models.Response{
		Data: model,
		Meta: *meta,
	}, nil
}

func GetMetaData(message string, statusCode int) (*models.Meta, error) {
	return &models.Meta{
		Message: message,
		Status:  statusCode,
	}, nil
}
