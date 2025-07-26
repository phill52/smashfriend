package repositories

import (
	"smashfriend/database"
	"smashfriend/models"
	"smashfriend/utils"
)

func GetResponse(model interface{}, page int, limit int, statusCode int, message string) *models.Response {
	query := database.DB.Model(&model)

	paginationData, err := utils.GetPaginationData(query, page, limit)
	if err != nil {
		return nil
	}
	meta, err := GetMetaData(message, statusCode)
	if err != nil {
		return nil
	}
	return &models.Response{
		Data:       model,
		Pagination: *paginationData,
		Meta:       *meta,
	}
}

func GetResponseWithoutPagination(model interface{}, statusCode int, message string) *models.Response {

	meta, err := GetMetaData(message, statusCode)
	if err != nil {
		return nil
	}
	return &models.Response{
		Data: model,
		Meta: *meta,
	}
}

func GetMetaData(message string, statusCode int) (*models.Meta, error) {
	return &models.Meta{
		Message: message,
		Status:  statusCode,
	}, nil
}
