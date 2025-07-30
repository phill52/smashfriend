package utils

import (
	"smashfriend/database"
	"smashfriend/utils/response"
)

func GetResponse(model interface{}, pagination interface{}, page, limit *int, statusCode int, message string) *response.Response {
	meta, err := GetMetaData(model, pagination, page, limit, message, statusCode)
	if err != nil {
		response := response.Response{}
		responseError := GetError(message, statusCode)
		response.Error = *responseError
		response.Meta = *meta
		return &response
	}
	response := &response.Response{
		Data: model,
		Meta: *meta,
	}
	return response
}

func GetMetaData(model interface{}, pagination interface{}, page, limit *int, message string, statusCode int) (*response.Meta, error) {
	response := response.Meta{
		Message: message,
		Status:  statusCode,
	}

	if pagination != nil {
		query := database.DB.Model(&model)
		paginationData, err := GetPaginationData(query, *page, *limit)
		if err != nil {
			return nil, err
		}
		response.Pagination = *paginationData
	}
	return &response, nil
}

func GetError(message string, statusCode int) *response.Error {
	return &response.Error{
		Message:    message,
		StatusCode: statusCode,
	}
}
