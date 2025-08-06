package utils

type Response struct {
	Data  interface{} `json:"data"`
	Meta  Meta        `json:"meta"`
	Error interface{} `json:"error,omitempty"`
}

type Meta struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Pagination interface{} `json:"pagination,omitempty"`
}

type Error struct {
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"code"`
}

func GetResponse(data, pagination interface{}, statusCode int, message string) *Response {
	meta := GetMetaData(pagination, statusCode, message)
	return &Response{
		Data: data,
		Meta: *meta,
	}
}

func GetErrorResponse(statusCode int, message string) *Response {
	meta := GetMetaData(nil, statusCode, message)
	responseError := GetError(message, statusCode)
	return &Response{
		Meta:  *meta,
		Error: *responseError,
	}

}

func GetMetaData(paginationData interface{}, statusCode int, message string) *Meta {
	response := Meta{
		Message: message,
		Status:  statusCode,
	}

	if paginationData != nil {
		response.Pagination = paginationData
	}
	return &response
}

func GetError(message string, statusCode int) *Error {
	return &Error{
		Message:    message,
		StatusCode: statusCode,
	}
}
