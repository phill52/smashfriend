package utils

import (
	"smashfriend/utils/response"
)

func GetResponse(data, pagination interface{}, statusCode int, message string) *response.Response {
	meta := GetMetaData(pagination, statusCode, message)
	if meta.Pagination == nil && statusCode != 200 {
		responseError := GetError(message, statusCode)
		return &response.Response{
			Meta:  *meta,
			Error: *responseError,
		}
	}
	return &response.Response{
		Data: data,
		Meta: *meta,
	}
}

func GetMetaData(paginationData interface{}, statusCode int, message string) *response.Meta {
	response := response.Meta{
		Message: message,
		Status:  statusCode,
	}

	if paginationData != nil {
		response.Pagination = paginationData
		return &response
	}
	return &response
}

func GetError(message string, statusCode int) *response.Error {
	return &response.Error{
		Message:    message,
		StatusCode: statusCode,
	}
}
