package response

type Meta struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Pagination interface{} `json:"pagination,omitempty"`
}
