package utils

import (
	"smashfriend/database"
	"smashfriend/utils/response"
)

func GetResponse(model interface{}, page, limit *int, statusCode int, message string) *response.Response {
	meta := GetMetaData(message, statusCode)

	response := &response.Response{
		Data: model,
		Meta: *meta,
	}

	if page != nil && limit != nil {
		query := database.DB.Model(&model)
		paginationData, err := GetPaginationData(query, *page, *limit)
		if err != nil {
			responseError := GetError(message, statusCode)
			response.Error = *responseError
			return response
		}
		response.Pagination = *paginationData
	}

	return response
}

func GetMetaData(message string, statusCode int) *response.Meta {
	return &response.Meta{
		Message: message,
		Status:  statusCode,
	}
}

func GetError(message string, statusCode int) *response.Error {
	return &response.Error{
		Message:    message,
		StatusCode: statusCode,
	}
}
