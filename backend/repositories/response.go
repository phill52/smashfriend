package repositories

import (
	"smashfriend/database"
	"smashfriend/models"
	"smashfriend/utils"
)

func GetResponse(model interface{}, page, limit *int, statusCode int, message string) *models.Response {
	meta := GetMetaData(message, statusCode)

	response := &models.Response{
		Data: model,
		Meta: *meta,
	}

	if page != nil && limit != nil {
		query := database.DB.Model(&model)
		paginationData, err := utils.GetPaginationData(query, *page, *limit)
		if err != nil {
			return nil
		}
		response.Pagination = *paginationData
	}

	return response
}

func GetMetaData(message string, statusCode int) *models.Meta {
	return &models.Meta{
		Message: message,
		Status:  statusCode,
	}
}
