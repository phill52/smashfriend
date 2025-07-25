package models

type Response struct {
	Data       interface{} `json:"data,omitempty"`
	Meta       Meta        `json:"meta"`
	Pagination interface{} `json:"pagination,omitempty"`
}
